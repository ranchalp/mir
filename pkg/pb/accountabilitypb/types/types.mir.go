// Code generated by Mir codegen. DO NOT EDIT.

package accountabilitypbtypes

import (
	mirreflect "github.com/filecoin-project/mir/codegen/mirreflect"
	types1 "github.com/filecoin-project/mir/codegen/model/types"
	accountabilitypb "github.com/filecoin-project/mir/pkg/pb/accountabilitypb"
	types2 "github.com/filecoin-project/mir/pkg/pb/trantorpb/types"
	types "github.com/filecoin-project/mir/pkg/types"
	reflectutil "github.com/filecoin-project/mir/pkg/util/reflectutil"
)

type Event struct {
	Type Event_Type
}

type Event_Type interface {
	mirreflect.GeneratedType
	isEvent_Type()
	Pb() accountabilitypb.Event_Type
}

type Event_TypeWrapper[T any] interface {
	Event_Type
	Unwrap() *T
}

func Event_TypeFromPb(pb accountabilitypb.Event_Type) Event_Type {
	if pb == nil {
		return nil
	}
	switch pb := pb.(type) {
	case *accountabilitypb.Event_Predecided:
		return &Event_Predecided{Predecided: PredecidedFromPb(pb.Predecided)}
	case *accountabilitypb.Event_Decided:
		return &Event_Decided{Decided: DecidedFromPb(pb.Decided)}
	case *accountabilitypb.Event_Poms:
		return &Event_Poms{Poms: PoMsFromPb(pb.Poms)}
	case *accountabilitypb.Event_InstanceParams:
		return &Event_InstanceParams{InstanceParams: InstanceParamsFromPb(pb.InstanceParams)}
	}
	return nil
}

type Event_Predecided struct {
	Predecided *Predecided
}

func (*Event_Predecided) isEvent_Type() {}

func (w *Event_Predecided) Unwrap() *Predecided {
	return w.Predecided
}

func (w *Event_Predecided) Pb() accountabilitypb.Event_Type {
	if w == nil {
		return nil
	}
	if w.Predecided == nil {
		return &accountabilitypb.Event_Predecided{}
	}
	return &accountabilitypb.Event_Predecided{Predecided: (w.Predecided).Pb()}
}

func (*Event_Predecided) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*accountabilitypb.Event_Predecided]()}
}

type Event_Decided struct {
	Decided *Decided
}

func (*Event_Decided) isEvent_Type() {}

func (w *Event_Decided) Unwrap() *Decided {
	return w.Decided
}

func (w *Event_Decided) Pb() accountabilitypb.Event_Type {
	if w == nil {
		return nil
	}
	if w.Decided == nil {
		return &accountabilitypb.Event_Decided{}
	}
	return &accountabilitypb.Event_Decided{Decided: (w.Decided).Pb()}
}

func (*Event_Decided) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*accountabilitypb.Event_Decided]()}
}

type Event_Poms struct {
	Poms *PoMs
}

func (*Event_Poms) isEvent_Type() {}

func (w *Event_Poms) Unwrap() *PoMs {
	return w.Poms
}

func (w *Event_Poms) Pb() accountabilitypb.Event_Type {
	if w == nil {
		return nil
	}
	if w.Poms == nil {
		return &accountabilitypb.Event_Poms{}
	}
	return &accountabilitypb.Event_Poms{Poms: (w.Poms).Pb()}
}

func (*Event_Poms) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*accountabilitypb.Event_Poms]()}
}

type Event_InstanceParams struct {
	InstanceParams *InstanceParams
}

func (*Event_InstanceParams) isEvent_Type() {}

func (w *Event_InstanceParams) Unwrap() *InstanceParams {
	return w.InstanceParams
}

func (w *Event_InstanceParams) Pb() accountabilitypb.Event_Type {
	if w == nil {
		return nil
	}
	if w.InstanceParams == nil {
		return &accountabilitypb.Event_InstanceParams{}
	}
	return &accountabilitypb.Event_InstanceParams{InstanceParams: (w.InstanceParams).Pb()}
}

func (*Event_InstanceParams) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*accountabilitypb.Event_InstanceParams]()}
}

func EventFromPb(pb *accountabilitypb.Event) *Event {
	if pb == nil {
		return nil
	}
	return &Event{
		Type: Event_TypeFromPb(pb.Type),
	}
}

func (m *Event) Pb() *accountabilitypb.Event {
	if m == nil {
		return nil
	}
	pbMessage := &accountabilitypb.Event{}
	{
		if m.Type != nil {
			pbMessage.Type = (m.Type).Pb()
		}
	}

	return pbMessage
}

func (*Event) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*accountabilitypb.Event]()}
}

type Predecided struct {
	Data []uint8
}

func PredecidedFromPb(pb *accountabilitypb.Predecided) *Predecided {
	if pb == nil {
		return nil
	}
	return &Predecided{
		Data: pb.Data,
	}
}

func (m *Predecided) Pb() *accountabilitypb.Predecided {
	if m == nil {
		return nil
	}
	pbMessage := &accountabilitypb.Predecided{}
	{
		pbMessage.Data = m.Data
	}

	return pbMessage
}

func (*Predecided) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*accountabilitypb.Predecided]()}
}

type Decided struct {
	Data []uint8
}

func DecidedFromPb(pb *accountabilitypb.Decided) *Decided {
	if pb == nil {
		return nil
	}
	return &Decided{
		Data: pb.Data,
	}
}

func (m *Decided) Pb() *accountabilitypb.Decided {
	if m == nil {
		return nil
	}
	pbMessage := &accountabilitypb.Decided{}
	{
		pbMessage.Data = m.Data
	}

	return pbMessage
}

func (*Decided) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*accountabilitypb.Decided]()}
}

type PoM struct {
	NodeId           types.NodeID
	ConflictingMsg_1 *SignedPredecision
	ConflictingMsg_2 *SignedPredecision
}

func PoMFromPb(pb *accountabilitypb.PoM) *PoM {
	if pb == nil {
		return nil
	}
	return &PoM{
		NodeId:           (types.NodeID)(pb.NodeId),
		ConflictingMsg_1: SignedPredecisionFromPb(pb.ConflictingMsg_1),
		ConflictingMsg_2: SignedPredecisionFromPb(pb.ConflictingMsg_2),
	}
}

func (m *PoM) Pb() *accountabilitypb.PoM {
	if m == nil {
		return nil
	}
	pbMessage := &accountabilitypb.PoM{}
	{
		pbMessage.NodeId = (string)(m.NodeId)
		if m.ConflictingMsg_1 != nil {
			pbMessage.ConflictingMsg_1 = (m.ConflictingMsg_1).Pb()
		}
		if m.ConflictingMsg_2 != nil {
			pbMessage.ConflictingMsg_2 = (m.ConflictingMsg_2).Pb()
		}
	}

	return pbMessage
}

func (*PoM) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*accountabilitypb.PoM]()}
}

type LightCertificate struct {
	Data []uint8
}

func LightCertificateFromPb(pb *accountabilitypb.LightCertificate) *LightCertificate {
	if pb == nil {
		return nil
	}
	return &LightCertificate{
		Data: pb.Data,
	}
}

func (m *LightCertificate) Pb() *accountabilitypb.LightCertificate {
	if m == nil {
		return nil
	}
	pbMessage := &accountabilitypb.LightCertificate{}
	{
		pbMessage.Data = m.Data
	}

	return pbMessage
}

func (*LightCertificate) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*accountabilitypb.LightCertificate]()}
}

type PoMs struct {
	Poms []*PoM
}

func PoMsFromPb(pb *accountabilitypb.PoMs) *PoMs {
	if pb == nil {
		return nil
	}
	return &PoMs{
		Poms: types1.ConvertSlice(pb.Poms, func(t *accountabilitypb.PoM) *PoM {
			return PoMFromPb(t)
		}),
	}
}

func (m *PoMs) Pb() *accountabilitypb.PoMs {
	if m == nil {
		return nil
	}
	pbMessage := &accountabilitypb.PoMs{}
	{
		pbMessage.Poms = types1.ConvertSlice(m.Poms, func(t *PoM) *accountabilitypb.PoM {
			return (t).Pb()
		})
	}

	return pbMessage
}

func (*PoMs) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*accountabilitypb.PoMs]()}
}

type Message struct {
	Type Message_Type
}

type Message_Type interface {
	mirreflect.GeneratedType
	isMessage_Type()
	Pb() accountabilitypb.Message_Type
}

type Message_TypeWrapper[T any] interface {
	Message_Type
	Unwrap() *T
}

func Message_TypeFromPb(pb accountabilitypb.Message_Type) Message_Type {
	if pb == nil {
		return nil
	}
	switch pb := pb.(type) {
	case *accountabilitypb.Message_SignedPredecision:
		return &Message_SignedPredecision{SignedPredecision: SignedPredecisionFromPb(pb.SignedPredecision)}
	case *accountabilitypb.Message_Certificate:
		return &Message_Certificate{Certificate: FullCertificateFromPb(pb.Certificate)}
	case *accountabilitypb.Message_Poms:
		return &Message_Poms{Poms: PoMsFromPb(pb.Poms)}
	case *accountabilitypb.Message_LightCertificate:
		return &Message_LightCertificate{LightCertificate: LightCertificateFromPb(pb.LightCertificate)}
	}
	return nil
}

type Message_SignedPredecision struct {
	SignedPredecision *SignedPredecision
}

func (*Message_SignedPredecision) isMessage_Type() {}

func (w *Message_SignedPredecision) Unwrap() *SignedPredecision {
	return w.SignedPredecision
}

func (w *Message_SignedPredecision) Pb() accountabilitypb.Message_Type {
	if w == nil {
		return nil
	}
	if w.SignedPredecision == nil {
		return &accountabilitypb.Message_SignedPredecision{}
	}
	return &accountabilitypb.Message_SignedPredecision{SignedPredecision: (w.SignedPredecision).Pb()}
}

func (*Message_SignedPredecision) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*accountabilitypb.Message_SignedPredecision]()}
}

type Message_Certificate struct {
	Certificate *FullCertificate
}

func (*Message_Certificate) isMessage_Type() {}

func (w *Message_Certificate) Unwrap() *FullCertificate {
	return w.Certificate
}

func (w *Message_Certificate) Pb() accountabilitypb.Message_Type {
	if w == nil {
		return nil
	}
	if w.Certificate == nil {
		return &accountabilitypb.Message_Certificate{}
	}
	return &accountabilitypb.Message_Certificate{Certificate: (w.Certificate).Pb()}
}

func (*Message_Certificate) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*accountabilitypb.Message_Certificate]()}
}

type Message_Poms struct {
	Poms *PoMs
}

func (*Message_Poms) isMessage_Type() {}

func (w *Message_Poms) Unwrap() *PoMs {
	return w.Poms
}

func (w *Message_Poms) Pb() accountabilitypb.Message_Type {
	if w == nil {
		return nil
	}
	if w.Poms == nil {
		return &accountabilitypb.Message_Poms{}
	}
	return &accountabilitypb.Message_Poms{Poms: (w.Poms).Pb()}
}

func (*Message_Poms) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*accountabilitypb.Message_Poms]()}
}

type Message_LightCertificate struct {
	LightCertificate *LightCertificate
}

func (*Message_LightCertificate) isMessage_Type() {}

func (w *Message_LightCertificate) Unwrap() *LightCertificate {
	return w.LightCertificate
}

func (w *Message_LightCertificate) Pb() accountabilitypb.Message_Type {
	if w == nil {
		return nil
	}
	if w.LightCertificate == nil {
		return &accountabilitypb.Message_LightCertificate{}
	}
	return &accountabilitypb.Message_LightCertificate{LightCertificate: (w.LightCertificate).Pb()}
}

func (*Message_LightCertificate) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*accountabilitypb.Message_LightCertificate]()}
}

func MessageFromPb(pb *accountabilitypb.Message) *Message {
	if pb == nil {
		return nil
	}
	return &Message{
		Type: Message_TypeFromPb(pb.Type),
	}
}

func (m *Message) Pb() *accountabilitypb.Message {
	if m == nil {
		return nil
	}
	pbMessage := &accountabilitypb.Message{}
	{
		if m.Type != nil {
			pbMessage.Type = (m.Type).Pb()
		}
	}

	return pbMessage
}

func (*Message) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*accountabilitypb.Message]()}
}

type SignedPredecision struct {
	Predecision []uint8
	Signature   []uint8
}

func SignedPredecisionFromPb(pb *accountabilitypb.SignedPredecision) *SignedPredecision {
	if pb == nil {
		return nil
	}
	return &SignedPredecision{
		Predecision: pb.Predecision,
		Signature:   pb.Signature,
	}
}

func (m *SignedPredecision) Pb() *accountabilitypb.SignedPredecision {
	if m == nil {
		return nil
	}
	pbMessage := &accountabilitypb.SignedPredecision{}
	{
		pbMessage.Predecision = m.Predecision
		pbMessage.Signature = m.Signature
	}

	return pbMessage
}

func (*SignedPredecision) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*accountabilitypb.SignedPredecision]()}
}

type FullCertificate struct {
	Certificate map[types.NodeID]*SignedPredecision
}

func FullCertificateFromPb(pb *accountabilitypb.FullCertificate) *FullCertificate {
	if pb == nil {
		return nil
	}
	return &FullCertificate{
		Certificate: types1.ConvertMap(pb.Certificate, func(k string, v *accountabilitypb.SignedPredecision) (types.NodeID, *SignedPredecision) {
			return (types.NodeID)(k), SignedPredecisionFromPb(v)
		}),
	}
}

func (m *FullCertificate) Pb() *accountabilitypb.FullCertificate {
	if m == nil {
		return nil
	}
	pbMessage := &accountabilitypb.FullCertificate{}
	{
		pbMessage.Certificate = types1.ConvertMap(m.Certificate, func(k types.NodeID, v *SignedPredecision) (string, *accountabilitypb.SignedPredecision) {
			return (string)(k), (v).Pb()
		})
	}

	return pbMessage
}

func (*FullCertificate) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*accountabilitypb.FullCertificate]()}
}

type InstanceParams struct {
	Membership *types2.Membership
}

func InstanceParamsFromPb(pb *accountabilitypb.InstanceParams) *InstanceParams {
	if pb == nil {
		return nil
	}
	return &InstanceParams{
		Membership: types2.MembershipFromPb(pb.Membership),
	}
}

func (m *InstanceParams) Pb() *accountabilitypb.InstanceParams {
	if m == nil {
		return nil
	}
	pbMessage := &accountabilitypb.InstanceParams{}
	{
		if m.Membership != nil {
			pbMessage.Membership = (m.Membership).Pb()
		}
	}

	return pbMessage
}

func (*InstanceParams) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*accountabilitypb.InstanceParams]()}
}
