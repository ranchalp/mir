package reconfiguration

import (
	"github.com/filecoin-project/mir/pkg/availability/asyncmultisigcollector/asyncmultisigcollectorprovider/internal/common"
	"github.com/filecoin-project/mir/pkg/dsl"
	t "github.com/filecoin-project/mir/pkg/types"
)

// IncludeReconfiguration handles reconfiguration for the async availability module
func IncludeReconfiguration(
	m dsl.Module,
	mc *common.ModuleConfig,
	params *common.ModuleParams,
	nodeID t.NodeID,
) {

	//UponReconfigure(newmembership){
	// if there is a membershipchange, start (v0) by simply trashing all modules and batches and creating a new one with new membership??
	//}

}
