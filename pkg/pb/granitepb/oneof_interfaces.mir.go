// Code generated by Mir codegen. DO NOT EDIT.

package granitepb

type Message_Type = isMessage_Type

type Message_TypeWrapper[T any] interface {
	Message_Type
	Unwrap() *T
}

func (w *Message_Decision) Unwrap() *Decision {
	return w.Decision
}

func (w *Message_ConsensusMsg) Unwrap() *ConsensusMsg {
	return w.ConsensusMsg
}

type Event_Type = isEvent_Type

type Event_TypeWrapper[T any] interface {
	Event_Type
	Unwrap() *T
}

func (w *Event_ConvergeTimeout) Unwrap() *ConvergeTimeout {
	return w.ConvergeTimeout
}
