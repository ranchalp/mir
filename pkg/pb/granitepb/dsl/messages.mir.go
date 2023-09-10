// Code generated by Mir codegen. DO NOT EDIT.

package granitepbdsl

import (
	dsl "github.com/filecoin-project/mir/pkg/dsl"
	types3 "github.com/filecoin-project/mir/pkg/granite/types"
	types "github.com/filecoin-project/mir/pkg/pb/granitepb/types"
	dsl1 "github.com/filecoin-project/mir/pkg/pb/messagepb/dsl"
	types2 "github.com/filecoin-project/mir/pkg/pb/messagepb/types"
	types1 "github.com/filecoin-project/mir/pkg/types"
)

// Module-specific dsl functions for processing net messages.

func UponMessageReceived[W types.Message_TypeWrapper[M], M any](m dsl.Module, handler func(from types1.NodeID, msg *M) error) {
	dsl1.UponMessageReceived[*types2.Message_Granite](m, func(from types1.NodeID, msg *types.Message) error {
		w, ok := msg.Type.(W)
		if !ok {
			return nil
		}

		return handler(from, w.Unwrap())
	})
}

func UponDecisionReceived(m dsl.Module, handler func(from types1.NodeID, round types3.RoundNr, data []uint8, signature []uint8) error) {
	UponMessageReceived[*types.Message_Decision](m, func(from types1.NodeID, msg *types.Decision) error {
		return handler(from, msg.Round, msg.Data, msg.Signature)
	})
}

func UponConsensusMsgReceived(m dsl.Module, handler func(from types1.NodeID, msgType types3.MsgType, round types3.RoundNr, data []uint8, ticket *types.Ticket, signature []uint8) error) {
	UponMessageReceived[*types.Message_ConsensusMsg](m, func(from types1.NodeID, msg *types.ConsensusMsg) error {
		return handler(from, msg.MsgType, msg.Round, msg.Data, msg.Ticket, msg.Signature)
	})
}