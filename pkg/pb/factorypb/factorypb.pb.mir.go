// Code generated by Mir codegen. DO NOT EDIT.

package factorypb

import (
	reflect "reflect"
)

func (*Event) ReflectTypeOptions() []reflect.Type {
	return []reflect.Type{
		reflect.TypeOf((*Event_NewModule)(nil)),
		reflect.TypeOf((*Event_GarbageCollect)(nil)),
	}
}

func (*GeneratorParams) ReflectTypeOptions() []reflect.Type {
	return []reflect.Type{
		reflect.TypeOf((*GeneratorParams_MultisigCollector)(nil)),
		reflect.TypeOf((*GeneratorParams_Checkpoint)(nil)),
		reflect.TypeOf((*GeneratorParams_EchoTestModule)(nil)),
		reflect.TypeOf((*GeneratorParams_PbftModule)(nil)),
		reflect.TypeOf((*GeneratorParams_PpvModule)(nil)),
		reflect.TypeOf((*GeneratorParams_AccModule)(nil)),
	}
}
