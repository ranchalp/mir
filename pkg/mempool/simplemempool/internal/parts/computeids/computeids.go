package computeids

import (
	"github.com/filecoin-project/mir/pkg/dsl"
	mpdsl "github.com/filecoin-project/mir/pkg/mempool/dsl"
	"github.com/filecoin-project/mir/pkg/mempool/simplemempool/internal/common"
	mppb "github.com/filecoin-project/mir/pkg/pb/mempoolpb"
	mppbdsl "github.com/filecoin-project/mir/pkg/pb/mempoolpb/dsl"
	mppbtypes "github.com/filecoin-project/mir/pkg/pb/mempoolpb/types"
	"github.com/filecoin-project/mir/pkg/pb/requestpb"
	"github.com/filecoin-project/mir/pkg/serializing"
	t "github.com/filecoin-project/mir/pkg/types"
)

// IncludeComputationOfTransactionAndBatchIDs registers event handler for processing RequestTransactionIDs and
// RequestBatchID events.
func IncludeComputationOfTransactionAndBatchIDs(
	m dsl.Module,
	mc *common.ModuleConfig,
	params *common.ModuleParams,
	commonState *common.State,
) {
	mpdsl.UponRequestTransactionIDs(m, func(txs []*requestpb.Request, origin *mppb.RequestTransactionIDsOrigin) error {
		txMsgs := make([][][]byte, len(txs))
		for i, tx := range txs {
			txMsgs[i] = serializing.RequestForHash(tx)
		}

		dsl.HashRequest(m, mc.Hasher, txMsgs, &computeHashForTransactionIDsContext{origin})
		return nil
	})

	dsl.UponHashResult(m, func(hashes [][]byte, context *computeHashForTransactionIDsContext) error {
		txIDs := make([]t.TxID, len(hashes))
		copy(txIDs, hashes)

		mpdsl.TransactionIDsResponse(m, t.ModuleID(context.origin.Module), txIDs, context.origin)
		return nil
	})

	mppbdsl.UponRequestBatchID(m, func(txIDs []t.TxID, prevBatch t.BatchID, origin *mppbtypes.RequestBatchIDOrigin) error {
		data := make([][]byte, len(txIDs))
		copy(data, txIDs)
		if prevBatch != nil || len(prevBatch) == 0 { // no prev batch
			data = append(data, prevBatch) //append prev batch to batchID
		}
		dsl.HashOneMessage(m, mc.Hasher, data, &computeHashForBatchIDContext{origin.Pb()})
		return nil
	})

	dsl.UponOneHashResult(m, func(hash []byte, context *computeHashForBatchIDContext) error {
		mpdsl.BatchIDResponse(m, t.ModuleID(context.origin.Module), hash, context.origin)
		return nil
	})
}

// Context data structures

type computeHashForTransactionIDsContext struct {
	origin *mppb.RequestTransactionIDsOrigin
}

type computeHashForBatchIDContext struct {
	origin *mppb.RequestBatchIDOrigin
}
