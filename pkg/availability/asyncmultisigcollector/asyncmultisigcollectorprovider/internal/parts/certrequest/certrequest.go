package certrequest

import (
	"encoding/hex"
	"fmt"
	"github.com/filecoin-project/mir/pkg/availability/asyncmultisigcollector/asyncmultisigcollectorprovider/internal/common"
	"github.com/filecoin-project/mir/pkg/dsl"
	factoryevents "github.com/filecoin-project/mir/pkg/factorymodule/events"
	apbdsl "github.com/filecoin-project/mir/pkg/pb/availabilitypb/dsl"
	"github.com/filecoin-project/mir/pkg/pb/availabilitypb/mscpb"
	apbtypes "github.com/filecoin-project/mir/pkg/pb/availabilitypb/types"
	"github.com/filecoin-project/mir/pkg/pb/factorymodulepb"
	t "github.com/filecoin-project/mir/pkg/types"
	"github.com/filecoin-project/mir/pkg/util/maputil"
	"reflect"
)

// State represents the state related to this part of the module.
type State struct {
	NextReqID    RequestID
	RequestState map[RequestID]*RequestState
}

// RequestID is used to uniquely identify an outgoing request.
type RequestID = uint64

// RequestState represents the state related to a request on the source node of the request.
// The node disposes of this state as soon as the request is completed.
type RequestState struct {
	ReqOrigin *apbtypes.RequestCertOrigin
	BatchID   t.BatchID

	receivedSig map[t.NodeID]bool
	sigs        map[t.NodeID][]byte
}

// IncludeRequestingCertificates constantly asks the worker to produce new certificates.
func IncludeRequestingCertificates(
	m dsl.Module,
	mc *common.ModuleConfig,
	params *common.ModuleParams,
	nodeID t.NodeID,
) {

	//
	// Events for workers whose nodeID is the source of the worker
	//
	apbdsl.UponRequestCert(m, func(origin *apbtypes.RequestCertOrigin) error {
		latestModule := true
		for _, key := range maputil.GetSortedKeysFunc(params.Certs[nodeID], func(i, j t.ModuleID) bool {
			return i > j // latest first
		}) {
			if cert, ok := params.Certs[nodeID][key]; ok && cert != nil {
				alreadyDecided := reflect.DeepEqual(cert.BatchId, params.LastDecision[nodeID].Cert.BatchId)
				if cert != nil && !alreadyDecided {
					//cert is ready at request
					apbdsl.NewCert(m, origin.Module, &apbtypes.Cert{Type: &apbtypes.Cert_Msc{Msc: cert}}, origin)
					//update latest decided batch
					params.LastDecision[nodeID] = &common.Decision{
						Cert:     cert,
						WorkerID: key,
					}
					if latestModule {
						//create new module and request a cert
						newWorkerID := createNewModule(m, mc, params, nodeID)
						apbdsl.RequestCert(m, newWorkerID, &RequestCertContext{
							workerID:    newWorkerID,
							workerCount: params.WorkerCount[nodeID],
							nodeID:      nodeID,
						})
					}
					return nil
				} else if alreadyDecided {
					// all previous are already decided too
					break
				} else {
					latestModule = false
				}
			}
		}

		// if we get here, that means that we need to wait to get a new cert (not yet ready but already requested) or that this is the first time

		params.LastReqOrigin = origin // store request to respond on response

		// If first time -> no workers -> create all workers and request cert (should happen once at the start)
		if _, ok := params.WorkerCount[nodeID]; !ok {
			for _, _nodeID := range maputil.GetSortedKeys(params.CurrentNodes) {
				params.WorkerCount[_nodeID] = 0
				newWorkerID := createNewModule(m, mc, params, _nodeID)
				if _nodeID == nodeID {
					apbdsl.RequestCert(m, newWorkerID, &RequestCertContext{
						workerID:    newWorkerID,
						workerCount: params.WorkerCount[_nodeID],
						nodeID:      nodeID,
					})
				}
			}
		}

		return nil
		// Otherwise, just mark this request and return once a batch is ready from one of the workers
	})

	apbdsl.UponNewCert(m,
		func(_cert *apbtypes.Cert, c *RequestCertContext) error {
			var (
				cert *apbtypes.Cert_Msc
				ok   bool
			)
			if cert, ok = _cert.Type.(*apbtypes.Cert_Msc); !ok {
				return fmt.Errorf(" invalid cert.Type type: %T, val %v", _cert.Type, _cert.Type)
			}
			params.Certs[nodeID][c.workerID] = cert.Msc

			params.CertLocation[hex.EncodeToString(cert.Msc.BatchId)] = common.BatchLoc{
				WorkerID: c.workerID,
				NodeID:   nodeID,
			}
			if params.LastReqOrigin != nil {
				//ISS waiting a certificate
				apbdsl.NewCert(m, params.LastReqOrigin.Module, _cert, params.LastReqOrigin)
				params.LastReqOrigin = nil // TODO handle more than one request at a time
			}

			if c.workerCount == params.WorkerCount[nodeID] {
				//This is the last worker created
				//create new module and request a cert
				newWorkerID := createNewModule(m, mc, params, nodeID)
				apbdsl.RequestCert(m, newWorkerID, &RequestCertContext{
					workerID:    newWorkerID,
					workerCount: params.WorkerCount[nodeID],
					nodeID:      nodeID,
				})
			}
			return nil
		},
	)

	//
	// Events for workers whose nodeID is NOT the source of the worker
	//

	// This event handles the notification of a certificate received for workers whose nodeID is not the worker
	apbdsl.UponCertReceived(m, func(_cert *apbtypes.Cert, source t.NodeID, workerID t.ModuleID, workerCount int64) error {
		if source == nodeID {
			return nil // do nothing, this is treated by UponNewCert
		}
		var (
			cert *apbtypes.Cert_Msc
			ok   bool
		)
		if cert, ok = _cert.Type.(*apbtypes.Cert_Msc); !ok {
			return fmt.Errorf(" invalid cert.Type type: %T, val %v", _cert.Type, _cert.Type)
		}
		params.Certs[source][workerID] = cert.Msc
		params.CertLocation[hex.EncodeToString(cert.Msc.BatchId)] = common.BatchLoc{
			WorkerID: workerID,
			NodeID:   source,
		}
		params.LastDecision[source] = &common.Decision{
			Cert:     cert.Msc,
			WorkerID: workerID,
		}
		if workerCount == params.WorkerCount[source] {
			//This is the last worker created
			//create new module and request a cert
			//but do not request cert (not source)
			createNewModule(m, mc, params, nodeID)
		}
		return nil
	})

}

func createNewModule(m dsl.Module, mc *common.ModuleConfig, params *common.ModuleParams, nodeID t.NodeID) t.ModuleID {
	defer func() {
		params.WorkerCount[nodeID] += 1
	}()

	newWorkerID := common.GetWorkerID(mc, params.WorkerCount[nodeID], nodeID)
	var prevBatchID string
	if params.WorkerCount[nodeID] > 0 { // not first worker, keep track of prev batch
		prevBatchID = string(params.Certs[nodeID][common.GetWorkerID(mc, params.WorkerCount[nodeID], nodeID)].BatchId)
	}
	dsl.EmitEvent(m, factoryevents.NewModule(
		mc.AvailabilityWorker,
		newWorkerID,
		t.RetentionIndex(params.WorkerCount[nodeID]),
		&factorymodulepb.GeneratorParams{Type: &factorymodulepb.GeneratorParams_MultisigCollector{
			MultisigCollector: &mscpb.InstanceParams{
				Membership:  t.MembershipPb(params.CurrentNodes),
				Source:      string(nodeID),
				WorkerCount: params.WorkerCount[nodeID],
				PrevBatchId: prevBatchID,
			},
		}},
	))

	params.Certs[nodeID][newWorkerID] = nil
	return newWorkerID
}

type RequestCertContext struct {
	workerID    t.ModuleID
	workerCount int64
	nodeID      t.NodeID
}
