// Code generated by Mir codegen. DO NOT EDIT.

package granitepbtypes

import (
	mirreflect "github.com/filecoin-project/mir/codegen/mirreflect"
	types "github.com/filecoin-project/mir/pkg/granite/types"
	granitepb "github.com/filecoin-project/mir/pkg/pb/granitepb"
	types1 "github.com/filecoin-project/mir/pkg/pb/trantorpb/types"
	reflectutil "github.com/filecoin-project/mir/pkg/util/reflectutil"
)

type Message struct {
	Type Message_Type
}

type Message_Type interface {
	mirreflect.GeneratedType
	isMessage_Type()
	Pb() granitepb.Message_Type
}

type Message_TypeWrapper[T any] interface {
	Message_Type
	Unwrap() *T
}

func Message_TypeFromPb(pb granitepb.Message_Type) Message_Type {
	if pb == nil {
		return nil
	}
	switch pb := pb.(type) {
	case *granitepb.Message_Decision:
		return &Message_Decision{Decision: DecisionFromPb(pb.Decision)}
	case *granitepb.Message_ConsensusMsg:
		return &Message_ConsensusMsg{ConsensusMsg: ConsensusMsgFromPb(pb.ConsensusMsg)}
	}
	return nil
}

type Message_Decision struct {
	Decision *Decision
}

func (*Message_Decision) isMessage_Type() {}

func (w *Message_Decision) Unwrap() *Decision {
	return w.Decision
}

func (w *Message_Decision) Pb() granitepb.Message_Type {
	if w == nil {
		return nil
	}
	if w.Decision == nil {
		return &granitepb.Message_Decision{}
	}
	return &granitepb.Message_Decision{Decision: (w.Decision).Pb()}
}

func (*Message_Decision) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*granitepb.Message_Decision]()}
}

type Message_ConsensusMsg struct {
	ConsensusMsg *ConsensusMsg
}

func (*Message_ConsensusMsg) isMessage_Type() {}

func (w *Message_ConsensusMsg) Unwrap() *ConsensusMsg {
	return w.ConsensusMsg
}

func (w *Message_ConsensusMsg) Pb() granitepb.Message_Type {
	if w == nil {
		return nil
	}
	if w.ConsensusMsg == nil {
		return &granitepb.Message_ConsensusMsg{}
	}
	return &granitepb.Message_ConsensusMsg{ConsensusMsg: (w.ConsensusMsg).Pb()}
}

func (*Message_ConsensusMsg) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*granitepb.Message_ConsensusMsg]()}
}

func MessageFromPb(pb *granitepb.Message) *Message {
	if pb == nil {
		return nil
	}
	return &Message{
		Type: Message_TypeFromPb(pb.Type),
	}
}

func (m *Message) Pb() *granitepb.Message {
	if m == nil {
		return nil
	}
	pbMessage := &granitepb.Message{}
	{
		if m.Type != nil {
			pbMessage.Type = (m.Type).Pb()
		}
	}

	return pbMessage
}

func (*Message) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*granitepb.Message]()}
}

type Event struct {
	Type Event_Type
}

type Event_Type interface {
	mirreflect.GeneratedType
	isEvent_Type()
	Pb() granitepb.Event_Type
}

type Event_TypeWrapper[T any] interface {
	Event_Type
	Unwrap() *T
}

func Event_TypeFromPb(pb granitepb.Event_Type) Event_Type {
	if pb == nil {
		return nil
	}
	switch pb := pb.(type) {
	case *granitepb.Event_ConvergeTimeout:
		return &Event_ConvergeTimeout{ConvergeTimeout: ConvergeTimeoutFromPb(pb.ConvergeTimeout)}
	}
	return nil
}

type Event_ConvergeTimeout struct {
	ConvergeTimeout *ConvergeTimeout
}

func (*Event_ConvergeTimeout) isEvent_Type() {}

func (w *Event_ConvergeTimeout) Unwrap() *ConvergeTimeout {
	return w.ConvergeTimeout
}

func (w *Event_ConvergeTimeout) Pb() granitepb.Event_Type {
	if w == nil {
		return nil
	}
	if w.ConvergeTimeout == nil {
		return &granitepb.Event_ConvergeTimeout{}
	}
	return &granitepb.Event_ConvergeTimeout{ConvergeTimeout: (w.ConvergeTimeout).Pb()}
}

func (*Event_ConvergeTimeout) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*granitepb.Event_ConvergeTimeout]()}
}

func EventFromPb(pb *granitepb.Event) *Event {
	if pb == nil {
		return nil
	}
	return &Event{
		Type: Event_TypeFromPb(pb.Type),
	}
}

func (m *Event) Pb() *granitepb.Event {
	if m == nil {
		return nil
	}
	pbMessage := &granitepb.Event{}
	{
		if m.Type != nil {
			pbMessage.Type = (m.Type).Pb()
		}
	}

	return pbMessage
}

func (*Event) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*granitepb.Event]()}
}

type ConsensusMsg struct {
	MsgType   types.MsgType
	Round     types.RoundNr
	Data      []uint8
	Ticket    *Ticket
	Signature []uint8
}

func ConsensusMsgFromPb(pb *granitepb.ConsensusMsg) *ConsensusMsg {
	if pb == nil {
		return nil
	}
	return &ConsensusMsg{
		MsgType:   (types.MsgType)(pb.MsgType),
		Round:     (types.RoundNr)(pb.Round),
		Data:      pb.Data,
		Ticket:    TicketFromPb(pb.Ticket),
		Signature: pb.Signature,
	}
}

func (m *ConsensusMsg) Pb() *granitepb.ConsensusMsg {
	if m == nil {
		return nil
	}
	pbMessage := &granitepb.ConsensusMsg{}
	{
		pbMessage.MsgType = (int32)(m.MsgType)
		pbMessage.Round = (uint64)(m.Round)
		pbMessage.Data = m.Data
		if m.Ticket != nil {
			pbMessage.Ticket = (m.Ticket).Pb()
		}
		pbMessage.Signature = m.Signature
	}

	return pbMessage
}

func (*ConsensusMsg) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*granitepb.ConsensusMsg]()}
}

type Ticket struct {
	Data []uint8
}

func TicketFromPb(pb *granitepb.Ticket) *Ticket {
	if pb == nil {
		return nil
	}
	return &Ticket{
		Data: pb.Data,
	}
}

func (m *Ticket) Pb() *granitepb.Ticket {
	if m == nil {
		return nil
	}
	pbMessage := &granitepb.Ticket{}
	{
		pbMessage.Data = m.Data
	}

	return pbMessage
}

func (*Ticket) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*granitepb.Ticket]()}
}

type ConvergeTimeout struct{}

func ConvergeTimeoutFromPb(pb *granitepb.ConvergeTimeout) *ConvergeTimeout {
	if pb == nil {
		return nil
	}
	return &ConvergeTimeout{}
}

func (m *ConvergeTimeout) Pb() *granitepb.ConvergeTimeout {
	if m == nil {
		return nil
	}
	pbMessage := &granitepb.ConvergeTimeout{}
	{
	}

	return pbMessage
}

func (*ConvergeTimeout) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*granitepb.ConvergeTimeout]()}
}

type Decision struct {
	Round     types.RoundNr
	Data      []uint8
	Signature []uint8
}

func DecisionFromPb(pb *granitepb.Decision) *Decision {
	if pb == nil {
		return nil
	}
	return &Decision{
		Round:     (types.RoundNr)(pb.Round),
		Data:      pb.Data,
		Signature: pb.Signature,
	}
}

func (m *Decision) Pb() *granitepb.Decision {
	if m == nil {
		return nil
	}
	pbMessage := &granitepb.Decision{}
	{
		pbMessage.Round = (uint64)(m.Round)
		pbMessage.Data = m.Data
		pbMessage.Signature = m.Signature
	}

	return pbMessage
}

func (*Decision) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*granitepb.Decision]()}
}

type InstanceParams struct {
	InstanceUid uint64
	Membership  *types1.Membership
}

func InstanceParamsFromPb(pb *granitepb.InstanceParams) *InstanceParams {
	if pb == nil {
		return nil
	}
	return &InstanceParams{
		InstanceUid: pb.InstanceUid,
		Membership:  types1.MembershipFromPb(pb.Membership),
	}
}

func (m *InstanceParams) Pb() *granitepb.InstanceParams {
	if m == nil {
		return nil
	}
	pbMessage := &granitepb.InstanceParams{}
	{
		pbMessage.InstanceUid = m.InstanceUid
		if m.Membership != nil {
			pbMessage.Membership = (m.Membership).Pb()
		}
	}

	return pbMessage
}

func (*InstanceParams) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*granitepb.InstanceParams]()}
}
