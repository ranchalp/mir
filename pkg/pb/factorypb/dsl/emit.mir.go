package factorypbdsl

import (
	dsl "github.com/filecoin-project/mir/pkg/dsl"
	events "github.com/filecoin-project/mir/pkg/pb/factorypb/events"
	types2 "github.com/filecoin-project/mir/pkg/pb/factorypb/types"
	types1 "github.com/filecoin-project/mir/pkg/trantor/types"
	types "github.com/filecoin-project/mir/pkg/types"
)

// Module-specific dsl functions for emitting events.

func NewModule(m dsl.Module, destModule types.ModuleID, moduleId types.ModuleID, retentionIndex types1.RetentionIndex, params *types2.GeneratorParams) {
	dsl.EmitMirEvent(m, events.NewModule(destModule, moduleId, retentionIndex, params))
}

func GarbageCollect(m dsl.Module, destModule types.ModuleID, retentionIndex types1.RetentionIndex) {
	dsl.EmitMirEvent(m, events.GarbageCollect(destModule, retentionIndex))
}
