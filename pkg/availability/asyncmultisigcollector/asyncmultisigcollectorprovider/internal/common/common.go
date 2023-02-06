package common

import (
	"fmt"
	mscpbtypes "github.com/filecoin-project/mir/pkg/pb/availabilitypb/mscpb/types"
	apbtypes "github.com/filecoin-project/mir/pkg/pb/availabilitypb/types"
	requestpbtypes "github.com/filecoin-project/mir/pkg/pb/requestpb/types"
	t "github.com/filecoin-project/mir/pkg/types"
	"reflect"
)

// InstanceUID is used to uniquely identify an instance of multisig collector.
// It is used to prevent cross-instance signature replay attack and should be unique across all executions.
type InstanceUID []byte

// Bytes returns the binary representation of the InstanceUID.
func (uid InstanceUID) Bytes() []byte {
	return uid
}

// ModuleConfig sets the module ids. All replicas are expected to use identical module configurations.
type ModuleConfig struct {
	Self               t.ModuleID // id of this module
	Mempool            t.ModuleID
	BatchDB            t.ModuleID
	Net                t.ModuleID
	Crypto             t.ModuleID
	AvailabilityWorker t.ModuleID
}

type Decision struct {
	Cert     *mscpbtypes.Cert
	WorkerID t.ModuleID
}

type BatchLoc struct {
	WorkerID t.ModuleID
	NodeID   t.NodeID
}

// ModuleParams sets the values for the parameters of an instance of the protocol.
// All replicas are expected to use identical module parameters.
type ModuleParams struct {
	Certs         map[t.NodeID]map[t.ModuleID]*mscpbtypes.Cert // the current set of certificates sorted by a key of nodeID (source) first, and workerID, if any
	CertLocation  map[string]BatchLoc                          // string means hex.EncodeString(t.BatchID) type. Map that gives the inverse of Certs
	LastDecision  map[t.NodeID]*Decision                       // the last batch that was decided by ISS (per source)
	WorkerCount   map[t.NodeID]int64                           // counter for unique worker module IDs
	CurrentNodes  map[t.NodeID]t.NodeAddress                   // the list of participating nodes
	LastReqOrigin *apbtypes.RequestCertOrigin                  // the module that initiated the request
	LastReqBatch  map[t.ModuleID][]*requestpbtypes.Request     // the batch that was requested
}

func GetWorkerID(mc *ModuleConfig, workerCount int64, nodeID t.NodeID) t.ModuleID {
	return mc.AvailabilityWorker.Then(t.ModuleID(fmt.Sprintf("%v", nodeID))).Then(t.ModuleID(fmt.Sprintf("%v", workerCount)))
}

// SigData is the binary data that should be signed for forming a certificate.
func SigData(instanceUID InstanceUID, batchID t.BatchID) [][]byte {
	return [][]byte{instanceUID.Bytes(), []byte("BATCH_STORED"), batchID}
}

// findWorker returns the worker module that is responsible for the given cert.
func FindWorker(workers map[t.ModuleID]*mscpbtypes.Cert, cert *mscpbtypes.Cert) t.ModuleID {
	for workerID, deliveredCert := range workers {
		if reflect.DeepEqual(cert.BatchId, deliveredCert.BatchId) {
			return workerID
		}
	}
	return t.ModuleID("")
}
