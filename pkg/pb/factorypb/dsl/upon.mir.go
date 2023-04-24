package factorypbdsl

import (
	dsl "github.com/filecoin-project/mir/pkg/dsl"
	types1 "github.com/filecoin-project/mir/pkg/pb/eventpb/types"
	types "github.com/filecoin-project/mir/pkg/pb/factorypb/types"
	types3 "github.com/filecoin-project/mir/pkg/trantor/types"
	types2 "github.com/filecoin-project/mir/pkg/types"
)

// Module-specific dsl functions for processing events.

func UponEvent[W types.Event_TypeWrapper[Ev], Ev any](m dsl.Module, handler func(ev *Ev) error) {
	dsl.UponMirEvent[*types1.Event_Factory](m, func(ev *types.Event) error {
		w, ok := ev.Type.(W)
		if !ok {
			return nil
		}

		return handler(w.Unwrap())
	})
}

func UponNewModule(m dsl.Module, handler func(moduleId types2.ModuleID, retentionIndex types3.RetentionIndex, params *types.GeneratorParams) error) {
	UponEvent[*types.Event_NewModule](m, func(ev *types.NewModule) error {
		return handler(ev.ModuleId, ev.RetentionIndex, ev.Params)
	})
}

func UponGarbageCollect(m dsl.Module, handler func(retentionIndex types3.RetentionIndex) error) {
	UponEvent[*types.Event_GarbageCollect](m, func(ev *types.GarbageCollect) error {
		return handler(ev.RetentionIndex)
	})
}