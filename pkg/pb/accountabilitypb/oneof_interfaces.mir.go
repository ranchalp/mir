// Code generated by Mir codegen. DO NOT EDIT.

package accountabilitypb

type Message_Type = isMessage_Type

type Message_TypeWrapper[T any] interface {
	Message_Type
	Unwrap() *T
}

func (w *Message_SignedPredecision) Unwrap() *SignedPredecision {
	return w.SignedPredecision
}

func (w *Message_Certificate) Unwrap() *FullCertificate {
	return w.Certificate
}

func (w *Message_Poms) Unwrap() *PoMs {
	return w.Poms
}

func (w *Message_LightCertificate) Unwrap() *LightCertificate {
	return w.LightCertificate
}

func (w *Message_RequestSbMessage) Unwrap() *RequestSBMessage {
	return w.RequestSbMessage
}

func (w *Message_ProvideSbMessage) Unwrap() *ProvideSBMessage {
	return w.ProvideSbMessage
}
