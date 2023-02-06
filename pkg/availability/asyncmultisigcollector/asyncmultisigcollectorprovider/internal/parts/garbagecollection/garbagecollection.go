package garbagecollection

import (
	"github.com/filecoin-project/mir/pkg/availability/asyncmultisigcollector/asyncmultisigcollectorprovider/internal/common"
	"github.com/filecoin-project/mir/pkg/dsl"
	t "github.com/filecoin-project/mir/pkg/types"
)

// IncludeGarbageCollect garbage collects checkpointed epochs
func IncludeGarbageCollect(
	m dsl.Module,
	mc *common.ModuleConfig,
	params *common.ModuleParams,
	nodeID t.NodeID,
) {

	//UponGarbageCollect(decidedEpoch){
	// Go through all the the worker modules and remove the ones that are decided at decidedEpoch or before
	//}

}
