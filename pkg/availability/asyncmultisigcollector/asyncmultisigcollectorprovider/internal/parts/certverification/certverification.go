package certverification

import (
	"encoding/hex"
	"fmt"
	"github.com/filecoin-project/mir/pkg/availability/asyncmultisigcollector/asyncmultisigcollectorprovider/internal/common"
	"github.com/filecoin-project/mir/pkg/dsl"
	apbdsl "github.com/filecoin-project/mir/pkg/pb/availabilitypb/dsl"
	apbtypes "github.com/filecoin-project/mir/pkg/pb/availabilitypb/types"
	t "github.com/filecoin-project/mir/pkg/types"
)

// IncludeVerificationOfCertificates registers event handlers for processing availabilitypb.VerifyCert events.
func IncludeVerificationOfCertificates(
	m dsl.Module,
	mc *common.ModuleConfig,
	params *common.ModuleParams,
	nodeID t.NodeID,
) {
	// When receive a request to verify a certificate, check that it is structurally correct and verify the signatures.
	apbdsl.UponVerifyCert(m, func(_cert *apbtypes.Cert, origin *apbtypes.VerifyCertOrigin) error {
		//Simply Relay to specific worker (verifying the last certificate should suffice)
		var (
			ok       bool
			certMsc  *apbtypes.Cert_Msc
			batchLoc common.BatchLoc
		)
		if certMsc, ok = _cert.Type.(*apbtypes.Cert_Msc); !ok {
			return fmt.Errorf(" invalid cert.Type type: %T, val %v", _cert.Type, _cert.Type)
		}

		if batchLoc, ok = params.CertLocation[hex.EncodeToString(certMsc.Msc.BatchId)]; !ok {
			// I do not locally have the batch
			// Invalid cert or garbage collected, or not yet delivered
			// return invalid cert event
			// TODO verify with Matej and previous code that this is what one should do in this case
			apbdsl.CertVerified(m, origin.Module, false, "certificate not found in any worker", origin)
		}

		// just relay to particular worker and tell worker to respond directly to ISS
		apbdsl.VerifyCert(m, batchLoc.WorkerID, _cert, origin)

		return nil
	})

}
