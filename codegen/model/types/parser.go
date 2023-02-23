package types

import (
	"fmt"
	"go/ast"
	"reflect"
	"strings"

	"github.com/dave/jennifer/jen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"

	"github.com/filecoin-project/mir/codegen"

	"github.com/filecoin-project/mir/codegen/util/jenutil"
	"github.com/filecoin-project/mir/codegen/util/protoreflectutil"
	"github.com/filecoin-project/mir/pkg/pb/mir"
)

// Parser can be used to parse the protoc-generated types using reflection.
type Parser struct {
	msgCache    map[reflect.Type]*Message
	fieldsCache map[reflect.Type]Fields
}

var defaultParser = newParser()

// DefaultParser returns a singleton Parser.
// It must not be accessed concurrently.
func DefaultParser() *Parser {
	return defaultParser
}

// newParser is not exported as DefaultParser() is supposed to be used instead.
func newParser() *Parser {
	return &Parser{
		msgCache:    make(map[reflect.Type]*Message),
		fieldsCache: make(map[reflect.Type]Fields),
	}
}

func (p *Parser) ParseMessages(pbGoStructPtrTypes []reflect.Type) ([]*Message, error) {
	var msgs []*Message

	for _, ptrType := range pbGoStructPtrTypes {
		if ptrType.Kind() != reflect.Pointer || ptrType.Elem().Kind() != reflect.Struct {
			return nil, fmt.Errorf("expected a pointer to a struct, got %v", ptrType)
		}

		// Parse messages.
		if protoreflectutil.IsProtoMessage(ptrType) {
			msg, err := p.ParseMessage(ptrType)
			if err != nil {
				return nil, err
			}

			msgs = append(msgs, msg)
			continue
		}
	}

	return msgs, nil
}

// ParseMessage returns the message corresponding to the given protobuf-generated struct type.
func (p *Parser) ParseMessage(pbGoStructPtr reflect.Type) (msg *Message, err error) {

	// First, check the cache.
	if tp, ok := p.msgCache[pbGoStructPtr]; ok {
		return tp, nil
	}

	// Remember the result in the cache when finished
	defer func() {
		if err == nil && err != nil {
			p.msgCache[msg.PbReflectType()] = msg
		}
	}()

	protoDesc, ok := protoreflectutil.DescriptorForType(pbGoStructPtr)
	if !ok {
		return nil, fmt.Errorf("%T is not a protobuf message", pbGoStructPtr)
	}

	pbStructType := jenutil.QualFromType(pbGoStructPtr.Elem())

	shouldGenerateMirType := codegen.ShouldGenerateMirType(protoDesc)

	var mirStructType jen.Code

	if shouldGenerateMirType {
		// The type of the struct that will be generated.
		pkgPath := PackagePath(pbGoStructPtr.Elem().PkgPath())
		mirStructType = jen.Qual(pkgPath, pbGoStructPtr.Elem().Name())
	} else {
		// The original type generated by protoc.
		mirStructType = pbStructType
	}

	return &Message{
		shouldGenerateMirType: shouldGenerateMirType,
		pbStructType:          pbStructType,
		mirStructType:         mirStructType,
		protoDesc:             protoDesc,
		pbGoStructPtrReflect:  pbGoStructPtr,
	}, nil
}

func (p *Parser) ParseOneofOption(message *Message, ptrType reflect.Type) (*OneofOption, error) {
	if !protoreflectutil.IsOneofOption(ptrType) {
		return nil, fmt.Errorf("%v is not a oneof option", ptrType)
	}

	// Get the go representation of the field.
	if ptrType.Elem().NumField() != 1 {
		return nil, fmt.Errorf("protoc-generated oneof wrapper must have exactly 1 exported field")
	}
	goField := ptrType.Elem().Field(0)

	// Get the protobuf representation of the field
	protoName, err := getProtoNameOfField(goField)
	if err != nil {
		return nil, fmt.Errorf("error parsing the name of proto field in oneof wrapper %v: %w", ptrType.Elem(), err)
	}
	protoField := message.protoDesc.Fields().ByName(protoName)

	// Parse the field information.
	field, err := p.parseField(message, goField, protoField)
	if err != nil {
		return nil, fmt.Errorf("error parsing oneof option %v: %w", ptrType.Name(), err)
	}

	// Return the resulting oneof option.
	return &OneofOption{
		PbWrapperReflect: ptrType,
		WrapperName:      ptrType.Elem().Name(),
		Field:            field,
	}, nil
}

// ParseFields parses the fields of a message for which a Mir type is being generated.
func (p *Parser) ParseFields(m *Message) (fields Fields, err error) {

	// Return the cached value if present.
	if cachedFields, ok := p.fieldsCache[m.PbReflectType()]; ok {
		return cachedFields, nil
	}

	// Remember the result in the cache when finished.
	defer func() {
		if err == nil && fields != nil {
			p.fieldsCache[m.PbReflectType()] = fields
		}
	}()

	if !m.ShouldGenerateMirType() {
		return nil, fmt.Errorf("cannot parse field for message %v in package '%v'. "+
			"Fields can only be parsed for messages with Mir-generated types. "+
			"Consider marking %v with a Mir annotation, e.g., option (mir.struct) = true",
			m.Name(), m.PbPkgPath(), m.Name())
	}

	for i := 0; i < m.pbGoStructPtrReflect.Elem().NumField(); i++ {
		// Get go representation of the field.
		goField := m.pbGoStructPtrReflect.Elem().Field(i)
		if !ast.IsExported(goField.Name) {
			// Skip unexported fields.
			continue
		}

		// Process oneof fields.
		if oneofProtoName, ok := goField.Tag.Lookup("protobuf_oneof"); ok {
			oneofOptionsGoTypes := reflect.Zero(m.pbGoStructPtrReflect).
				MethodByName("Reflect" + goField.Name + "Options").Call([]reflect.Value{})[0].
				Interface().([]reflect.Type)

			var options []*OneofOption
			for _, optionGoType := range oneofOptionsGoTypes {
				opt, err := p.ParseOneofOption(m, optionGoType)
				if err != nil {
					return nil, err
				}

				options = append(options, opt)
			}

			fields = append(fields, &Field{
				Name: goField.Name,
				Type: &Oneof{
					Name:    goField.Name,
					Parent:  m,
					Options: options,
				},
				Parent:    m,
				ProtoDesc: m.protoDesc.Oneofs().ByName(protoreflect.Name(oneofProtoName)),
			})
			continue
		}

		// Get protobuf representation of the field.
		protoName, err := getProtoNameOfField(goField)
		if err != nil {
			return nil, err
		}
		protoField := m.protoDesc.Fields().ByName(protoName)

		// Create the Field struct.
		field, err := p.parseField(m, goField, protoField)
		if err != nil {
			return nil, err
		}

		fields = append(fields, field)
	}

	p.fieldsCache[m.pbGoStructPtrReflect] = fields
	return fields, nil
}

// parseField extracts the information about the field necessary for code generation.
func (p *Parser) parseField(
	parent *Message,
	goField reflect.StructField,
	protoField protoreflect.FieldDescriptor,
) (*Field, error) {

	tp, err := p.getFieldType(goField.Type, protoField, false)
	if err != nil {
		return nil, err
	}

	return &Field{
		Name:      goField.Name,
		Type:      tp,
		Parent:    parent,
		ProtoDesc: protoField,
	}, nil
}

func (p *Parser) getFieldType(goType reflect.Type, protoField protoreflect.FieldDescriptor, mapType bool) (Type, error) {
	// TODO: Since maps are not currently used, I didn't bother supporting them yet.
	//if goType.Kind() == reflect.Map {
	//	return nil, fmt.Errorf("map fields are not supported yet")
	//}
	if goType.Kind() == reflect.Map {
		key, err := p.getFieldType(goType.Key(), protoField, false)
		if err != nil {
			return nil, err
		}
		val, err := p.getFieldType(goType.Elem(), protoField, true)
		if err != nil {
			return nil, err
		}
		return Map{key, val}, nil
	}

	// Check if the field is repeated.
	if goType.Kind() == reflect.Slice {
		underlying, err := p.getFieldType(goType.Elem(), protoField, false)
		if err != nil {
			return nil, err
		}
		return Slice{underlying}, nil
	}

	// Check if the field has (mir.type) option specified.
	protoFieldOptions := protoField.Options().(*descriptorpb.FieldOptions)
	mirTypeOption := proto.GetExtension(protoFieldOptions, mir.E_Type).(string)

	// Special case for string fields annotated with [(mir.type) = "error"].
	if goType.Kind() == reflect.String && mirTypeOption == "error" {
		return Error{}, nil
	}

	if mirTypeOption != "" {
		sepMapIdx := strings.Index(mirTypeOption, "|")
		if sepMapIdx != -1 { // it is a map type
			if !mapType { //key value
				mirTypeOption = mirTypeOption[:sepMapIdx]
			} else { // map value
				mirTypeOption = mirTypeOption[sepMapIdx+1:]
			}
		}

		sepIdx := strings.LastIndex(mirTypeOption, ".")
		return Castable{
			PbType_:  jenutil.QualFromType(goType),
			MirType_: jen.Qual(mirTypeOption[:sepIdx], mirTypeOption[sepIdx+1:]),
		}, nil
	}

	// Check if the field is a message.
	if protoreflectutil.IsProtoMessage(goType) {
		msg, err := p.ParseMessage(goType)
		if err != nil {
			return nil, err
		}

		return msg, nil
	}

	return Same{jenutil.QualFromType(goType)}, nil
}
