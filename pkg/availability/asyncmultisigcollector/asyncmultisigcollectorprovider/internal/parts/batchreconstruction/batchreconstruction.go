package batchreconstruction

import (
	"encoding/hex"
	"fmt"
	"github.com/filecoin-project/mir/pkg/availability/asyncmultisigcollector/asyncmultisigcollectorprovider/internal/common"
	"github.com/filecoin-project/mir/pkg/dsl"
	mempooldsl "github.com/filecoin-project/mir/pkg/mempool/dsl"
	batchdbpbdsl "github.com/filecoin-project/mir/pkg/pb/availabilitypb/batchdbpb/dsl"
	apbdsl "github.com/filecoin-project/mir/pkg/pb/availabilitypb/dsl"
	mscpbdsl "github.com/filecoin-project/mir/pkg/pb/availabilitypb/mscpb/dsl"
	mscpbtypes "github.com/filecoin-project/mir/pkg/pb/availabilitypb/mscpb/types"
	apbtypes "github.com/filecoin-project/mir/pkg/pb/availabilitypb/types"
	requestpbtypes "github.com/filecoin-project/mir/pkg/pb/requestpb/types"
	t "github.com/filecoin-project/mir/pkg/types"
	"github.com/filecoin-project/mir/pkg/util/maputil"
	"reflect"
)

// State represents the state related to this part of the module.
type State struct {
	NextReqID    t.RequestID
	RequestState map[t.RequestID]*RequestState
}

// RequestState represents the state related to a request on the source node of the request.
// The node disposes of this state as soon as the request is completed.
type RequestState struct {
	BatchID   t.BatchID
	ReqOrigin *apbtypes.RequestTransactionsOrigin
}

// IncludeBatchReconstruction registers event handlers for processing availabilitypb.RequestTransactions events.
func IncludeBatchReconstruction(
	m dsl.Module,
	mc *common.ModuleConfig,
	params *common.ModuleParams,
	nodeID t.NodeID,
) {

	apbdsl.UponRequestTransactions(m, func(_cert *apbtypes.Cert, origin *apbtypes.RequestTransactionsOrigin) error {
		// Identify module, ask for the batch to the module and recursively ask for previous batches if needed.
		var (
			ok      bool
			certMsc *apbtypes.Cert_Msc
			batchLoc common.BatchLoc
		)

		if params.LastReqBatch != nil {
			return fmt.Errorf("last request batch is not nil")
		}

		if certMsc, ok = _cert.Type.(*apbtypes.Cert_Msc); !ok {
			return fmt.Errorf(" invalid cert.Type type: %T, val %v", _cert.Type, _cert.Type)
		}
		if batchLoc, ok = params.CertLocation[hex.EncodeToString(certMsc.Msc.BatchId)]; !ok {
			// I do not locally have the batch
			// Invalid cert or garbage collected, or not yet delivered
			// return invalid cert event
			TODO ASK FOR THE BATCH(ES?) TO OTHER NODES
			CONTINUE HERE
		}

		lastDecision := params.LastDecision[batchLoc.NodeID]
		certLocation := params.CertLocation[hex.EncodeToString(certMsc.Msc.BatchId)]
		workerIDKeys := maputil.GetSortedKeys(params.Certs[certLocation.NodeID])
		if lastDecision != nil { //not first time
			// Advance to request only new delivered certs not already provided
			for i, workerID := range workerIDKeys {
				if reflect.DeepEqual(params.Certs[certLocation.NodeID][workerID].BatchId, lastDecision.Cert.BatchId) {
					workerIDKeys = workerIDKeys[i+1:]
					break
				}
			}
		}

		params.LastReqBatch = make(map[t.ModuleID][]*requestpbtypes.Request)//TODO ATT store with ReqNo to fix race condition with consensus
		// Request all batches from these workers until reaching the requested batch
		for _, workerID := range workerIDKeys {
			apbdsl.RequestTransactions(m, workerID, _cert, &txRequestContext{
				workerID: workerID,
				origin: origin,
			})
			params.LastReqBatch[workerID] = nil // add entry to map
			if reflect.DeepEqual(params.Certs[certLocation.NodeID][workerID].BatchId, certMsc.Msc.BatchId) { // reached requested batch
				// Update to the new last decision
				params.LastDecision[certLocation.NodeID] = &common.Decision{Cert: certMsc.Msc, WorkerID: workerID}
				break
			}
		}

		return nil
	})

	apbdsl.UponProvideTransactions(m, func(_txs []*requestpbtypes.Request, c *txRequestContext) error {
		//If last batch that was requested by the provider itself, then provide all transactions at once to original requester
		//Otherwise, store for response and do nothing
		val, ok := params.LastReqBatch[c.workerID]
		// If the key exists
		if !ok {
			return fmt.Errorf("workerID %v not found in LastReqBatch", c.workerID)
		} else if val != nil {
			return fmt.Errorf("workerID %v already has a value in LastReqBatch", c.workerID)
		}

		params.LastReqBatch[c.workerID] = _txs
		for _, val := range params.LastReqBatch {
			if val == nil {
				return nil // not all workers have responded yet
			}
		}// All workers have responded

		//append all transactions from all workers
		txs := make([]*requestpbtypes.Request, 0, len(params.LastReqBatch))
		for _, _txs := range params.LastReqBatch {
			txs = append(txs, _txs...) // TODO perhaps too big a request at some point. In that case, improve here
		}
		apbdsl.ProvideTransactions(m, c.origin.Module, txs, c.origin)

		// restore to nil for next request
		params.LastReqBatch = nil
		return nil
	})

	mscpbdsl.UponRequestCertRangeMessageReceived(m, func(from t.NodeID, batchIDFrom []byte, batchIDTo []byte) error {
		// TODO IF locally available batch, then provide a range of batches in a message that goes from
		// the provided last decision to the actual batch
		var(
			certLocTo,
			certLocFrom common.BatchLoc
			ok          bool
		)

		if certLocFrom, ok = params.CertLocation[hex.EncodeToString(batchIDFrom)]; !ok {
			//this is impossible, wrong request or safety broken
			return nil
		}
		if certLocTo, ok = params.CertLocation[hex.EncodeToString(batchIDTo]; !ok {
			return nil // I do not locally have the batch, so ignore (TODO Response with some ack?)
		}

		workerIDKeys := maputil.GetSortedKeys(params.Certs[certLocFrom.NodeID])
		//iterate and aggregate batches from to and send them
		// Advance to provide requested batches
		for i, workerID := range workerIDKeys {
			if workerID == certLocFrom.WorkerID {
				workerIDKeys = workerIDKeys[i:]// send back the from cert so that the receiver can identify where certs go
				break
			}
		}
		linkedCert := make([]*mscpbtypes.Cert, 0)
		for _, workerID := range workerIDKeys {
			linkedCert = append(linkedCert,
					params.Certs[certLocFrom.NodeID][workerID])
			if workerID == certLocTo.WorkerID {
				break; // done with requested cert range
			}

		}

		return nil
	})
	mscpbdsl.UponProvideCertRangeMessageReceived(m, func(from t.NodeID, certs []*mscpbtypes.Cert) error {
		// verify the batches from first (the one that prevBatch poinst to the local Last decision)
		// recursively update local state artificially creating and deciding workers
		// call again uponRequestTransaction so that this time there is a return (or compute it here directly)

		return nil
	})


	// If the batch is present in the local storage, return it. Otherwise, ask the nodes that signed the certificate.
	batchdbpbdsl.UponLookupBatchResponse(m, func(found bool, txs []*requestpbtypes.Request, metadata []byte, context *lookupBatchLocallyContext) error {
		Verify this//This should not happen at the level of the provider, but workers instead
	})

	// When receive a request for batch from another node, lookup the batch in the local storage.
	mscpbdsl.UponRequestBatchMessageReceived(m, func(from t.NodeID, batchID t.BatchID, reqID t.RequestID) error {
		//This should not happen at the level of the provider, but workers instead
	})

	// If the batch is found in the local storage, send it to the requesting node.
	batchdbpbdsl.UponLookupBatchResponse(m, func(found bool, txs []*requestpbtypes.Request, metadata []byte, context *lookupBatchOnRemoteRequestContext) error {
		//This should not happen at the level of the provider, but workers instead
	})

	// When receive a requested batch, compute the ids of the received transactions.
	mscpbdsl.UponProvideBatchMessageReceived(m, func(from t.NodeID, txs []*requestpbtypes.Request, reqID t.RequestID) error {
		//This should not happen at the level of the provider, but workers instead
	})

	// When transaction ids are computed, compute the id of the batch.
	mempooldsl.UponTransactionIDsResponse(m, func(txIDs []t.TxID, context *requestTxIDsContext) error {
		//This should not happen at the level of the provider, but workers instead
	})

	// When the id of the batch is computed, check if it is correct, store the transactions, and complete the request.
	mempooldsl.UponBatchIDResponse(m, func(batchID t.BatchID, context *requestBatchIDContext) error {
		//This should not happen at the level of the provider, but workers instead
	})

	batchdbpbdsl.UponBatchStored(m, func(_ *storeBatchContext) error {
		//This should not happen at the level of the provider, but workers instead
		return nil
	})
}

// Context data structures

type lookupBatchLocallyContext struct {
	cert   *mscpbtypes.Cert
	origin *apbtypes.RequestTransactionsOrigin
}

type lookupBatchOnRemoteRequestContext struct {
	requester t.NodeID
	reqID     t.RequestID
}

type requestTxIDsContext struct {
	reqID t.RequestID
	txs   []*requestpbtypes.Request
}

type requestBatchIDContext struct {
	reqID t.RequestID
	txs   []*requestpbtypes.Request
	txIDs []t.TxID
}

type storeBatchContext struct{}

// txRequestContext saves the context of requesting transactions from the worker.
type txRequestContext struct {
	workerID t.ModuleID
	origin *apbtypes.RequestTransactionsOrigin
}
