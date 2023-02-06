package asyncmultisigcollectorprovider

import (
	"github.com/filecoin-project/mir/pkg/availability/asyncmultisigcollector/asyncmultisigcollectorprovider/internal/common"
	"github.com/filecoin-project/mir/pkg/availability/asyncmultisigcollector/asyncmultisigcollectorprovider/internal/parts/batchreconstruction"
	"github.com/filecoin-project/mir/pkg/availability/asyncmultisigcollector/asyncmultisigcollectorprovider/internal/parts/certrequest"
	"github.com/filecoin-project/mir/pkg/availability/asyncmultisigcollector/asyncmultisigcollectorprovider/internal/parts/certverification"
	"github.com/filecoin-project/mir/pkg/availability/asyncmultisigcollector/asyncmultisigcollectorprovider/internal/parts/garbagecollection"
	"github.com/filecoin-project/mir/pkg/availability/asyncmultisigcollector/asyncmultisigcollectorprovider/internal/parts/reconfiguration"
	"github.com/filecoin-project/mir/pkg/dsl"
	"github.com/filecoin-project/mir/pkg/modules"
	t "github.com/filecoin-project/mir/pkg/types"
)

// ModuleConfig sets the module ids. All replicas are expected to use identical module configurations.
type ModuleConfig = common.ModuleConfig

// DefaultModuleConfig returns a valid module config with default names for all modules.
func DefaultModuleConfig() *ModuleConfig {
	return &ModuleConfig{
		Self:               "availabilityprovider",
		Mempool:            "mempool",
		BatchDB:            "batchdb",
		Net:                "net",
		Crypto:             "crypto",
		AvailabilityWorker: "availabilityworker",
	}
}

// ModuleParams sets the values for the parameters of an instance of the protocol.
// All replicas are expected to use identical module parameters.
type ModuleParams = common.ModuleParams

// NewModule creates a new instance of the multisig collector module.
// Multisig collector is the simplest implementation of the availability layer.
// Whenever an availability certificate is requested, it pulls a batch from the mempool module,
// sends it to all replicas and collects params.F+1 signatures confirming that
// other nodes have persistently stored the batch.
func NewModule(mc *ModuleConfig, params *ModuleParams, nodeID t.NodeID) (modules.PassiveModule, error) {
	m := dsl.NewModule(mc.Self)

	certrequest.IncludeRequestingCertificates(m, mc, params, nodeID)
	batchreconstruction.IncludeBatchReconstruction(m, mc, params, nodeID)
	certverification.IncludeVerificationOfCertificates(m, mc, params, nodeID)
	reconfiguration.IncludeReconfiguration(m, mc, params, nodeID) // if there is a membershipchange, start (v0) by simply trashing all modules and batches and creating a new one with new membership
	garbagecollection.IncludeGarbageCollect(m, mc, params, nodeID)
	return m, nil
}

//func NewReconfigurableModule(mc *ModuleConfig, nodeID t.NodeID, logger logging.Logger) modules.PassiveModule {
//	if logger == nil {
//		logger = logging.ConsoleErrorLogger
//	}
//	return factorymodule.New(
//		mc.Self,
//		factorymodule.DefaultParams(
//
//			// This function will be called whenever the factory module
//			// is asked to create a new instance of the multisig collector.
//			func(mscID t.ModuleID, params *factorymodulepb.GeneratorParams) (modules.PassiveModule, error) {
//
//				// Extract the IDs of the nodes in the membership associated with this instance
//				m := params.Type.(*factorymodulepb.GeneratorParams_MultisigCollector).MultisigCollector.Membership
//				mscNodeIDs := maputil.GetSortedKeys(t.Membership(m))
//
//				// Create a copy of basic module config with an adapted ID for the submodule.
//				submc := *mc
//				submc.Self = mscID
//
//				// Create a new instance of the multisig collector.
//				asyncmultisigcollectorworker, err := NewModule(
//					&submc,
//					&ModuleParams{
//						// TODO: Use InstanceUIDs properly.
//						//       (E.g., concatenate this with the instantiating protocol's InstanceUID when introduced.)
//						InstanceUID: []byte(mscID),
//						AllNodes:    mscNodeIDs,
//						// TODO: Consider lowering this threshold or make it configurable
//						//       for the case where fault assumptions are stricter because of other modules.
//						F: (len(mscNodeIDs) - 1) / 2,
//					},
//					nodeID,
//				)
//				if err != nil {
//					return nil, err
//				}
//				return asyncmultisigcollectorworker, nil
//			},
//		),
//		logger,
//	)
//}
