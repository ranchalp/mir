// Code generated by Mir codegen. DO NOT EDIT.

package bcbpbtypes

import (
	mirreflect "github.com/filecoin-project/mir/codegen/mirreflect"
	types1 "github.com/filecoin-project/mir/codegen/model/types"
	bcbpb "github.com/filecoin-project/mir/pkg/pb/bcbpb"
	types "github.com/filecoin-project/mir/pkg/types"
	reflectutil "github.com/filecoin-project/mir/pkg/util/reflectutil"
)

type Event struct {
	Type Event_Type
}

type Event_Type interface {
	mirreflect.GeneratedType
	isEvent_Type()
	Pb() bcbpb.Event_Type
}

type Event_TypeWrapper[T any] interface {
	Event_Type
	Unwrap() *T
}

func Event_TypeFromPb(pb bcbpb.Event_Type) Event_Type {
	if pb == nil {
		return nil
	}
	switch pb := pb.(type) {
	case *bcbpb.Event_Request:
		return &Event_Request{Request: BroadcastRequestFromPb(pb.Request)}
	case *bcbpb.Event_Deliver:
		return &Event_Deliver{Deliver: DeliverFromPb(pb.Deliver)}
	}
	return nil
}

type Event_Request struct {
	Request *BroadcastRequest
}

func (*Event_Request) isEvent_Type() {}

func (w *Event_Request) Unwrap() *BroadcastRequest {
	return w.Request
}

func (w *Event_Request) Pb() bcbpb.Event_Type {
	if w == nil {
		return nil
	}
	if w.Request == nil {
		return &bcbpb.Event_Request{}
	}
	return &bcbpb.Event_Request{Request: (w.Request).Pb()}
}

func (*Event_Request) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*bcbpb.Event_Request]()}
}

type Event_Deliver struct {
	Deliver *Deliver
}

func (*Event_Deliver) isEvent_Type() {}

func (w *Event_Deliver) Unwrap() *Deliver {
	return w.Deliver
}

func (w *Event_Deliver) Pb() bcbpb.Event_Type {
	if w == nil {
		return nil
	}
	if w.Deliver == nil {
		return &bcbpb.Event_Deliver{}
	}
	return &bcbpb.Event_Deliver{Deliver: (w.Deliver).Pb()}
}

func (*Event_Deliver) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*bcbpb.Event_Deliver]()}
}

func EventFromPb(pb *bcbpb.Event) *Event {
	if pb == nil {
		return nil
	}
	return &Event{
		Type: Event_TypeFromPb(pb.Type),
	}
}

func (m *Event) Pb() *bcbpb.Event {
	if m == nil {
		return nil
	}
	pbMessage := &bcbpb.Event{}
	{
		if m.Type != nil {
			pbMessage.Type = (m.Type).Pb()
		}
	}

	return pbMessage
}

func (*Event) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*bcbpb.Event]()}
}

type BroadcastRequest struct {
	Data []uint8
}

func BroadcastRequestFromPb(pb *bcbpb.BroadcastRequest) *BroadcastRequest {
	if pb == nil {
		return nil
	}
	return &BroadcastRequest{
		Data: pb.Data,
	}
}

func (m *BroadcastRequest) Pb() *bcbpb.BroadcastRequest {
	if m == nil {
		return nil
	}
	pbMessage := &bcbpb.BroadcastRequest{}
	{
		pbMessage.Data = m.Data
	}

	return pbMessage
}

func (*BroadcastRequest) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*bcbpb.BroadcastRequest]()}
}

type Deliver struct {
	Data []uint8
}

func DeliverFromPb(pb *bcbpb.Deliver) *Deliver {
	if pb == nil {
		return nil
	}
	return &Deliver{
		Data: pb.Data,
	}
}

func (m *Deliver) Pb() *bcbpb.Deliver {
	if m == nil {
		return nil
	}
	pbMessage := &bcbpb.Deliver{}
	{
		pbMessage.Data = m.Data
	}

	return pbMessage
}

func (*Deliver) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*bcbpb.Deliver]()}
}

type Message struct {
	Type Message_Type
}

type Message_Type interface {
	mirreflect.GeneratedType
	isMessage_Type()
	Pb() bcbpb.Message_Type
}

type Message_TypeWrapper[T any] interface {
	Message_Type
	Unwrap() *T
}

func Message_TypeFromPb(pb bcbpb.Message_Type) Message_Type {
	if pb == nil {
		return nil
	}
	switch pb := pb.(type) {
	case *bcbpb.Message_StartMessage:
		return &Message_StartMessage{StartMessage: StartMessageFromPb(pb.StartMessage)}
	case *bcbpb.Message_EchoMessage:
		return &Message_EchoMessage{EchoMessage: EchoMessageFromPb(pb.EchoMessage)}
	case *bcbpb.Message_FinalMessage:
		return &Message_FinalMessage{FinalMessage: FinalMessageFromPb(pb.FinalMessage)}
	}
	return nil
}

type Message_StartMessage struct {
	StartMessage *StartMessage
}

func (*Message_StartMessage) isMessage_Type() {}

func (w *Message_StartMessage) Unwrap() *StartMessage {
	return w.StartMessage
}

func (w *Message_StartMessage) Pb() bcbpb.Message_Type {
	if w == nil {
		return nil
	}
	if w.StartMessage == nil {
		return &bcbpb.Message_StartMessage{}
	}
	return &bcbpb.Message_StartMessage{StartMessage: (w.StartMessage).Pb()}
}

func (*Message_StartMessage) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*bcbpb.Message_StartMessage]()}
}

type Message_EchoMessage struct {
	EchoMessage *EchoMessage
}

func (*Message_EchoMessage) isMessage_Type() {}

func (w *Message_EchoMessage) Unwrap() *EchoMessage {
	return w.EchoMessage
}

func (w *Message_EchoMessage) Pb() bcbpb.Message_Type {
	if w == nil {
		return nil
	}
	if w.EchoMessage == nil {
		return &bcbpb.Message_EchoMessage{}
	}
	return &bcbpb.Message_EchoMessage{EchoMessage: (w.EchoMessage).Pb()}
}

func (*Message_EchoMessage) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*bcbpb.Message_EchoMessage]()}
}

type Message_FinalMessage struct {
	FinalMessage *FinalMessage
}

func (*Message_FinalMessage) isMessage_Type() {}

func (w *Message_FinalMessage) Unwrap() *FinalMessage {
	return w.FinalMessage
}

func (w *Message_FinalMessage) Pb() bcbpb.Message_Type {
	if w == nil {
		return nil
	}
	if w.FinalMessage == nil {
		return &bcbpb.Message_FinalMessage{}
	}
	return &bcbpb.Message_FinalMessage{FinalMessage: (w.FinalMessage).Pb()}
}

func (*Message_FinalMessage) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*bcbpb.Message_FinalMessage]()}
}

func MessageFromPb(pb *bcbpb.Message) *Message {
	if pb == nil {
		return nil
	}
	return &Message{
		Type: Message_TypeFromPb(pb.Type),
	}
}

func (m *Message) Pb() *bcbpb.Message {
	if m == nil {
		return nil
	}
	pbMessage := &bcbpb.Message{}
	{
		if m.Type != nil {
			pbMessage.Type = (m.Type).Pb()
		}
	}

	return pbMessage
}

func (*Message) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*bcbpb.Message]()}
}

type StartMessage struct {
	Data []uint8
}

func StartMessageFromPb(pb *bcbpb.StartMessage) *StartMessage {
	if pb == nil {
		return nil
	}
	return &StartMessage{
		Data: pb.Data,
	}
}

func (m *StartMessage) Pb() *bcbpb.StartMessage {
	if m == nil {
		return nil
	}
	pbMessage := &bcbpb.StartMessage{}
	{
		pbMessage.Data = m.Data
	}

	return pbMessage
}

func (*StartMessage) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*bcbpb.StartMessage]()}
}

type EchoMessage struct {
	Signature []uint8
}

func EchoMessageFromPb(pb *bcbpb.EchoMessage) *EchoMessage {
	if pb == nil {
		return nil
	}
	return &EchoMessage{
		Signature: pb.Signature,
	}
}

func (m *EchoMessage) Pb() *bcbpb.EchoMessage {
	if m == nil {
		return nil
	}
	pbMessage := &bcbpb.EchoMessage{}
	{
		pbMessage.Signature = m.Signature
	}

	return pbMessage
}

func (*EchoMessage) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*bcbpb.EchoMessage]()}
}

type FinalMessage struct {
	Data       []uint8
	Signers    []types.NodeID
	Signatures [][]uint8
}

func FinalMessageFromPb(pb *bcbpb.FinalMessage) *FinalMessage {
	if pb == nil {
		return nil
	}
	return &FinalMessage{
		Data: pb.Data,
		Signers: types1.ConvertSlice(pb.Signers, func(t string) types.NodeID {
			return (types.NodeID)(t)
		}),
		Signatures: pb.Signatures,
	}
}

func (m *FinalMessage) Pb() *bcbpb.FinalMessage {
	if m == nil {
		return nil
	}
	pbMessage := &bcbpb.FinalMessage{}
	{
		pbMessage.Data = m.Data
		pbMessage.Signers = types1.ConvertSlice(m.Signers, func(t types.NodeID) string {
			return (string)(t)
		})
		pbMessage.Signatures = m.Signatures
	}

	return pbMessage
}

func (*FinalMessage) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*bcbpb.FinalMessage]()}
}
