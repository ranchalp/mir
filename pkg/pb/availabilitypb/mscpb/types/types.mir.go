package mscpbtypes

import (
	mirreflect "github.com/filecoin-project/mir/codegen/mirreflect"
	types2 "github.com/filecoin-project/mir/codegen/model/types"
	mscpb "github.com/filecoin-project/mir/pkg/pb/availabilitypb/mscpb"
	commonpb "github.com/filecoin-project/mir/pkg/pb/commonpb"
	requestpb "github.com/filecoin-project/mir/pkg/pb/requestpb"
	types "github.com/filecoin-project/mir/pkg/pb/requestpb/types"
	types1 "github.com/filecoin-project/mir/pkg/types"
	reflectutil "github.com/filecoin-project/mir/pkg/util/reflectutil"
)

type Message struct {
	Type Message_Type
}

type Message_Type interface {
	mirreflect.GeneratedType
	isMessage_Type()
	Pb() mscpb.Message_Type
}

type Message_TypeWrapper[T any] interface {
	Message_Type
	Unwrap() *T
}

func Message_TypeFromPb(pb mscpb.Message_Type) Message_Type {
	switch pb := pb.(type) {
	case *mscpb.Message_RequestSig:
		return &Message_RequestSig{RequestSig: RequestSigMessageFromPb(pb.RequestSig)}
	case *mscpb.Message_Sig:
		return &Message_Sig{Sig: SigMessageFromPb(pb.Sig)}
	case *mscpb.Message_RequestBatch:
		return &Message_RequestBatch{RequestBatch: RequestBatchMessageFromPb(pb.RequestBatch)}
	case *mscpb.Message_ProvideBatch:
		return &Message_ProvideBatch{ProvideBatch: ProvideBatchMessageFromPb(pb.ProvideBatch)}
	case *mscpb.Message_Cert:
		return &Message_Cert{Cert: CertMessageFromPb(pb.Cert)}
	case *mscpb.Message_RequestCertRange:
		return &Message_RequestCertRange{RequestCertRange: RequestCertRangeMessageFromPb(pb.RequestCertRange)}
	case *mscpb.Message_ProvideCertRange:
		return &Message_ProvideCertRange{ProvideCertRange: ProvideCertRangeMessageFromPb(pb.ProvideCertRange)}
	}
	return nil
}

type Message_RequestSig struct {
	RequestSig *RequestSigMessage
}

func (*Message_RequestSig) isMessage_Type() {}

func (w *Message_RequestSig) Unwrap() *RequestSigMessage {
	return w.RequestSig
}

func (w *Message_RequestSig) Pb() mscpb.Message_Type {
	return &mscpb.Message_RequestSig{RequestSig: (w.RequestSig).Pb()}
}

func (*Message_RequestSig) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mscpb.Message_RequestSig]()}
}

type Message_Sig struct {
	Sig *SigMessage
}

func (*Message_Sig) isMessage_Type() {}

func (w *Message_Sig) Unwrap() *SigMessage {
	return w.Sig
}

func (w *Message_Sig) Pb() mscpb.Message_Type {
	return &mscpb.Message_Sig{Sig: (w.Sig).Pb()}
}

func (*Message_Sig) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mscpb.Message_Sig]()}
}

type Message_RequestBatch struct {
	RequestBatch *RequestBatchMessage
}

func (*Message_RequestBatch) isMessage_Type() {}

func (w *Message_RequestBatch) Unwrap() *RequestBatchMessage {
	return w.RequestBatch
}

func (w *Message_RequestBatch) Pb() mscpb.Message_Type {
	return &mscpb.Message_RequestBatch{RequestBatch: (w.RequestBatch).Pb()}
}

func (*Message_RequestBatch) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mscpb.Message_RequestBatch]()}
}

type Message_ProvideBatch struct {
	ProvideBatch *ProvideBatchMessage
}

func (*Message_ProvideBatch) isMessage_Type() {}

func (w *Message_ProvideBatch) Unwrap() *ProvideBatchMessage {
	return w.ProvideBatch
}

func (w *Message_ProvideBatch) Pb() mscpb.Message_Type {
	return &mscpb.Message_ProvideBatch{ProvideBatch: (w.ProvideBatch).Pb()}
}

func (*Message_ProvideBatch) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mscpb.Message_ProvideBatch]()}
}

type Message_Cert struct {
	Cert *CertMessage
}

func (*Message_Cert) isMessage_Type() {}

func (w *Message_Cert) Unwrap() *CertMessage {
	return w.Cert
}

func (w *Message_Cert) Pb() mscpb.Message_Type {
	return &mscpb.Message_Cert{Cert: (w.Cert).Pb()}
}

func (*Message_Cert) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mscpb.Message_Cert]()}
}

type Message_RequestCertRange struct {
	RequestCertRange *RequestCertRangeMessage
}

func (*Message_RequestCertRange) isMessage_Type() {}

func (w *Message_RequestCertRange) Unwrap() *RequestCertRangeMessage {
	return w.RequestCertRange
}

func (w *Message_RequestCertRange) Pb() mscpb.Message_Type {
	return &mscpb.Message_RequestCertRange{RequestCertRange: (w.RequestCertRange).Pb()}
}

func (*Message_RequestCertRange) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mscpb.Message_RequestCertRange]()}
}

type Message_ProvideCertRange struct {
	ProvideCertRange *ProvideCertRangeMessage
}

func (*Message_ProvideCertRange) isMessage_Type() {}

func (w *Message_ProvideCertRange) Unwrap() *ProvideCertRangeMessage {
	return w.ProvideCertRange
}

func (w *Message_ProvideCertRange) Pb() mscpb.Message_Type {
	return &mscpb.Message_ProvideCertRange{ProvideCertRange: (w.ProvideCertRange).Pb()}
}

func (*Message_ProvideCertRange) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mscpb.Message_ProvideCertRange]()}
}

func MessageFromPb(pb *mscpb.Message) *Message {
	return &Message{
		Type: Message_TypeFromPb(pb.Type),
	}
}

func (m *Message) Pb() *mscpb.Message {
	return &mscpb.Message{
		Type: (m.Type).Pb(),
	}
}

func (*Message) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mscpb.Message]()}
}

type RequestSigMessage struct {
	Txs       []*types.Request
	PrevBatch types1.BatchID
	ReqId     uint64
}

func RequestSigMessageFromPb(pb *mscpb.RequestSigMessage) *RequestSigMessage {
	return &RequestSigMessage{
		Txs: types2.ConvertSlice(pb.Txs, func(t *requestpb.Request) *types.Request {
			return types.RequestFromPb(t)
		}),
		PrevBatch: (types1.BatchID)(pb.PrevBatch),
		ReqId:     pb.ReqId,
	}
}

func (m *RequestSigMessage) Pb() *mscpb.RequestSigMessage {
	return &mscpb.RequestSigMessage{
		Txs: types2.ConvertSlice(m.Txs, func(t *types.Request) *requestpb.Request {
			return (t).Pb()
		}),
		PrevBatch: (string)(m.PrevBatch),
		ReqId:     m.ReqId,
	}
}

func (*RequestSigMessage) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mscpb.RequestSigMessage]()}
}

type SigMessage struct {
	Signature []uint8
	ReqId     uint64
}

func SigMessageFromPb(pb *mscpb.SigMessage) *SigMessage {
	return &SigMessage{
		Signature: pb.Signature,
		ReqId:     pb.ReqId,
	}
}

func (m *SigMessage) Pb() *mscpb.SigMessage {
	return &mscpb.SigMessage{
		Signature: m.Signature,
		ReqId:     m.ReqId,
	}
}

func (*SigMessage) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mscpb.SigMessage]()}
}

type RequestBatchMessage struct {
	BatchId []uint8
	ReqId   uint64
}

func RequestBatchMessageFromPb(pb *mscpb.RequestBatchMessage) *RequestBatchMessage {
	return &RequestBatchMessage{
		BatchId: pb.BatchId,
		ReqId:   pb.ReqId,
	}
}

func (m *RequestBatchMessage) Pb() *mscpb.RequestBatchMessage {
	return &mscpb.RequestBatchMessage{
		BatchId: m.BatchId,
		ReqId:   m.ReqId,
	}
}

func (*RequestBatchMessage) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mscpb.RequestBatchMessage]()}
}

type ProvideBatchMessage struct {
	Txs   []*types.Request
	ReqId uint64
}

func ProvideBatchMessageFromPb(pb *mscpb.ProvideBatchMessage) *ProvideBatchMessage {
	return &ProvideBatchMessage{
		Txs: types2.ConvertSlice(pb.Txs, func(t *requestpb.Request) *types.Request {
			return types.RequestFromPb(t)
		}),
		ReqId: pb.ReqId,
	}
}

func (m *ProvideBatchMessage) Pb() *mscpb.ProvideBatchMessage {
	return &mscpb.ProvideBatchMessage{
		Txs: types2.ConvertSlice(m.Txs, func(t *types.Request) *requestpb.Request {
			return (t).Pb()
		}),
		ReqId: m.ReqId,
	}
}

func (*ProvideBatchMessage) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mscpb.ProvideBatchMessage]()}
}

type RequestCertRangeMessage struct {
	BatchIdFrom []uint8
	BatchIdTo   []uint8
}

func RequestCertRangeMessageFromPb(pb *mscpb.RequestCertRangeMessage) *RequestCertRangeMessage {
	return &RequestCertRangeMessage{
		BatchIdFrom: pb.BatchIdFrom,
		BatchIdTo:   pb.BatchIdTo,
	}
}

func (m *RequestCertRangeMessage) Pb() *mscpb.RequestCertRangeMessage {
	return &mscpb.RequestCertRangeMessage{
		BatchIdFrom: m.BatchIdFrom,
		BatchIdTo:   m.BatchIdTo,
	}
}

func (*RequestCertRangeMessage) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mscpb.RequestCertRangeMessage]()}
}

type ProvideCertRangeMessage struct {
	Certs []*Cert
}

func ProvideCertRangeMessageFromPb(pb *mscpb.ProvideCertRangeMessage) *ProvideCertRangeMessage {
	return &ProvideCertRangeMessage{
		Certs: types2.ConvertSlice(pb.Certs, func(t *mscpb.Cert) *Cert {
			return CertFromPb(t)
		}),
	}
}

func (m *ProvideCertRangeMessage) Pb() *mscpb.ProvideCertRangeMessage {
	return &mscpb.ProvideCertRangeMessage{
		Certs: types2.ConvertSlice(m.Certs, func(t *Cert) *mscpb.Cert {
			return (t).Pb()
		}),
	}
}

func (*ProvideCertRangeMessage) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mscpb.ProvideCertRangeMessage]()}
}

type Cert struct {
	BatchId    []uint8
	Signers    []types1.NodeID
	Signatures [][]uint8
}

func CertFromPb(pb *mscpb.Cert) *Cert {
	return &Cert{
		BatchId: pb.BatchId,
		Signers: types2.ConvertSlice(pb.Signers, func(t string) types1.NodeID {
			return (types1.NodeID)(t)
		}),
		Signatures: pb.Signatures,
	}
}

func (m *Cert) Pb() *mscpb.Cert {
	return &mscpb.Cert{
		BatchId: m.BatchId,
		Signers: types2.ConvertSlice(m.Signers, func(t types1.NodeID) string {
			return (string)(t)
		}),
		Signatures: m.Signatures,
	}
}

func (*Cert) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mscpb.Cert]()}
}

type CertMessage struct {
	Msc *Cert
}

func CertMessageFromPb(pb *mscpb.CertMessage) *CertMessage {
	return &CertMessage{
		Msc: CertFromPb(pb.Msc),
	}
}

func (m *CertMessage) Pb() *mscpb.CertMessage {
	return &mscpb.CertMessage{
		Msc: (m.Msc).Pb(),
	}
}

func (*CertMessage) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mscpb.CertMessage]()}
}

type InstanceParams struct {
	Membership  *commonpb.Membership
	Source      types1.NodeID
	WorkerCount int64
	PrevBatchId types1.BatchID
}

func InstanceParamsFromPb(pb *mscpb.InstanceParams) *InstanceParams {
	return &InstanceParams{
		Membership:  pb.Membership,
		Source:      (types1.NodeID)(pb.Source),
		WorkerCount: pb.WorkerCount,
		PrevBatchId: (types1.BatchID)(pb.PrevBatchId),
	}
}

func (m *InstanceParams) Pb() *mscpb.InstanceParams {
	return &mscpb.InstanceParams{
		Membership:  m.Membership,
		Source:      (string)(m.Source),
		WorkerCount: m.WorkerCount,
		PrevBatchId: (string)(m.PrevBatchId),
	}
}

func (*InstanceParams) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*mscpb.InstanceParams]()}
}
