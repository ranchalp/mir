package poms

import (
	"reflect"

	"github.com/filecoin-project/mir/pkg/accountability/simpleacc/common"
	incommon "github.com/filecoin-project/mir/pkg/accountability/simpleacc/internal/common"
	"github.com/filecoin-project/mir/pkg/dsl"
	"github.com/filecoin-project/mir/pkg/logging"
	accpbdsl "github.com/filecoin-project/mir/pkg/pb/accountabilitypb/dsl"
	accpbtypes "github.com/filecoin-project/mir/pkg/pb/accountabilitypb/types"
	cryptopbdsl "github.com/filecoin-project/mir/pkg/pb/cryptopb/dsl"
	cryptopbtypes "github.com/filecoin-project/mir/pkg/pb/cryptopb/types"
	t "github.com/filecoin-project/mir/pkg/types"
)

// IncludePoMs verifies receives PoMs and sends found PoMs to other members.
func IncludePoMs(
	m dsl.Module,
	mc *common.ModuleConfig,
	params *common.ModuleParams,
	state *incommon.State,
	logger logging.Logger,
) {
	accpbdsl.UponPoMsReceived(m, func(from t.NodeID, poms []*accpbtypes.PoM) error {
		nodeIds := make([]t.NodeID, 0, 2*len(poms))
		data := make([]*cryptopbtypes.SignedData, 0, 2*len(poms))
		signatures := make([][]byte, 0, 2*len(poms))

		for _, pom := range poms {
			if reflect.DeepEqual(pom.ConflictingMsg_1.Predecision, pom.ConflictingMsg_2.Predecision) ||
				reflect.DeepEqual(pom.ConflictingMsg_1.Signature, pom.ConflictingMsg_2.Signature) { // no PoM possible here
				continue
			}

			if _, ok := params.Membership.Nodes[pom.NodeId]; !ok {
				continue
			}

			if _, ok := state.HandledPoMs[pom.NodeId]; ok {
				continue
			}

			nodeIds = append(nodeIds, pom.NodeId, pom.NodeId)

			data = append(data,
				&cryptopbtypes.SignedData{Data: [][]byte{pom.ConflictingMsg_1.Predecision, []byte(mc.Self)}},
				&cryptopbtypes.SignedData{Data: [][]byte{pom.ConflictingMsg_2.Predecision, []byte(mc.Self)}})

			signatures = append(signatures, pom.ConflictingMsg_1.Signature, pom.ConflictingMsg_2.Signature)
		}

		if len(data) == 0 {
			logger.Log(logging.LevelDebug, "Received empty PoM")
			return nil
		}

		cryptopbdsl.VerifySigs(
			m,
			mc.Crypto,
			data,
			signatures,
			nodeIds,
			&verifyPoMs{poms},
		)
		return nil
	})

	cryptopbdsl.UponSigsVerified(m, func(nodeIds []t.NodeID, errs []error, allOk bool, vpoms *verifyPoMs) error {
		for i := 0; i < len(nodeIds); i += 2 {
			if errs[i] == nil && errs[i+1] == nil {
				state.UnhandledPoMs = append(state.UnhandledPoMs, vpoms.poms[i/2])
			}
		}

		HandlePoMs(m, mc, params, state, logger)

		return nil
	})
}

// HandlePoMs sends all PoMs in State.UnhandledPoMs to all nodes and to the application module (from the POV of this module, i.e. mc.App).
func HandlePoMs(
	m dsl.Module,
	mc *common.ModuleConfig,
	params *common.ModuleParams,
	state *incommon.State,
	logger logging.Logger,
) {
	if len(state.UnhandledPoMs) == 0 {
		return
	}
	logger.Log(logging.LevelWarn, "Found valid PoMs! sending...")

	// Handle PoMs according to the application's logic defined when creating the accountability factory
	params.PomsHandler(m, mc, params, state, state.UnhandledPoMs, logger)

	for _, pom := range state.UnhandledPoMs {
		state.HandledPoMs[pom.NodeId] = pom
	}

	state.UnhandledPoMs = make([]*accpbtypes.PoM, 0)
}

type verifyPoMs struct {
	poms []*accpbtypes.PoM
}