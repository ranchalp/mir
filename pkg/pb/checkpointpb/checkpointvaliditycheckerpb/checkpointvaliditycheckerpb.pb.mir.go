// Code generated by Mir codegen. DO NOT EDIT.

package checkpointvaliditycheckerpb

import (
	reflect "reflect"
)

func (*Event) ReflectTypeOptions() []reflect.Type {
	return []reflect.Type{
		reflect.TypeOf((*Event_ValidateCheckpoint)(nil)),
		reflect.TypeOf((*Event_CheckpointValidated)(nil)),
	}
}

func (*ValidateChkpOrigin) ReflectTypeOptions() []reflect.Type {
	return []reflect.Type{
		reflect.TypeOf((*ValidateChkpOrigin_ContextStore)(nil)),
		reflect.TypeOf((*ValidateChkpOrigin_Dsl)(nil)),
	}
}
