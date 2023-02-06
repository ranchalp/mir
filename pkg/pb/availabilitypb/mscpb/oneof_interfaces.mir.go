package mscpb

type Message_Type = isMessage_Type

type Message_TypeWrapper[T any] interface {
	Message_Type
	Unwrap() *T
}

func (w *Message_RequestSig) Unwrap() *RequestSigMessage {
	return w.RequestSig
}

func (w *Message_Sig) Unwrap() *SigMessage {
	return w.Sig
}

func (w *Message_RequestBatch) Unwrap() *RequestBatchMessage {
	return w.RequestBatch
}

func (w *Message_ProvideBatch) Unwrap() *ProvideBatchMessage {
	return w.ProvideBatch
}

func (w *Message_Cert) Unwrap() *CertMessage {
	return w.Cert
}

func (w *Message_RequestCertRange) Unwrap() *RequestCertRangeMessage {
	return w.RequestCertRange
}

func (w *Message_ProvideCertRange) Unwrap() *ProvideCertRangeMessage {
	return w.ProvideCertRange
}
