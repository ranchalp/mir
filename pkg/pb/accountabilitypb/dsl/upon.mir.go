// Code generated by Mir codegen. DO NOT EDIT.

package accountabilitypbdsl

import (
	dsl "github.com/filecoin-project/mir/pkg/dsl"
	types "github.com/filecoin-project/mir/pkg/pb/accountabilitypb/types"
	types1 "github.com/filecoin-project/mir/pkg/pb/eventpb/types"
	types2 "github.com/filecoin-project/mir/pkg/pb/trantorpb/types"
)

// Module-specific dsl functions for processing events.

func UponEvent[W types.Event_TypeWrapper[Ev], Ev any](m dsl.Module, handler func(ev *Ev) error) {
	dsl.UponMirEvent[*types1.Event_Accountability](m, func(ev *types.Event) error {
		w, ok := ev.Type.(W)
		if !ok {
			return nil
		}

		return handler(w.Unwrap())
	})
}

func UponPredecided(m dsl.Module, handler func(data []uint8) error) {
	UponEvent[*types.Event_Predecided](m, func(ev *types.Predecided) error {
		return handler(ev.Data)
	})
}

func UponDecided(m dsl.Module, handler func(data []uint8) error) {
	UponEvent[*types.Event_Decided](m, func(ev *types.Decided) error {
		return handler(ev.Data)
	})
}

func UponPoMs(m dsl.Module, handler func(poms []*types.PoM) error) {
	UponEvent[*types.Event_Poms](m, func(ev *types.PoMs) error {
		return handler(ev.Poms)
	})
}

func UponInstanceParams(m dsl.Module, handler func(membership *types2.Membership) error) {
	UponEvent[*types.Event_InstanceParams](m, func(ev *types.InstanceParams) error {
		return handler(ev.Membership)
	})
}
