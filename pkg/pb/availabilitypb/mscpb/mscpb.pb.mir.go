package mscpb

import (
	reflect "reflect"
)

func (*Message) ReflectTypeOptions() []reflect.Type {
	return []reflect.Type{
		reflect.TypeOf((*Message_RequestSig)(nil)),
		reflect.TypeOf((*Message_Sig)(nil)),
		reflect.TypeOf((*Message_RequestBatch)(nil)),
		reflect.TypeOf((*Message_ProvideBatch)(nil)),
		reflect.TypeOf((*Message_Cert)(nil)),
		reflect.TypeOf((*Message_RequestCertRange)(nil)),
		reflect.TypeOf((*Message_ProvideCertRange)(nil)),
	}
}
