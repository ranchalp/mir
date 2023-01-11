/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package issutil

import (
	"fmt"
	"sort"

	t "github.com/filecoin-project/mir/pkg/types"
)

type LeaderPolicyType uint64

const (
	Simple LeaderPolicyType = iota
	Blacklist
)

// A LeaderSelectionPolicy implements the algorithm for selecting a set of leaders in each ISS epoch.
// In a nutshell, it gathers information about suspected leaders in the past epochs
// and uses it to calculate the set of leaders for future epochs.
// Its state can be updated using Suspect() and the leader set for an epoch is queried using Leaders().
// A leader set policy must be deterministic, i.e., calling Leaders() after the same sequence of Suspect() invocations
// always returns the same set of leaders at every Node.
type LeaderSelectionPolicy interface {

	// Leaders returns the (ordered) list of leaders based on the given epoch e and on the state of this policy object.
	Leaders() []t.NodeID

	// Suspect updates the state of the policy object by announcing it that node `node` has been suspected in epoch `e`.
	Suspect(e t.EpochNr, node t.NodeID)

	// Reconfigure returns a new LeaderSelectionPolicy based on the state of the current one,
	// but using a new configuration.
	// TODO: Use the whole configuration, not just the node IDs.
	Reconfigure(nodeIDs []t.NodeID) LeaderSelectionPolicy

	//// TODO: Define, implement, and use state serialization and restoration for leader selection policies.
	//Pb() commonpb.LeaderSelectionPolicy
	Bytes() []byte
}

func LeaderPolicyFromBytes(bytes []byte) (LeaderSelectionPolicy, error) {
	leaderPolicyType := t.Uint64FromBytes(bytes[0:8])

	switch LeaderPolicyType(leaderPolicyType) {
	case Simple:
		return SimpleLeaderPolicyFromBytes(bytes[8:]), nil
	case Blacklist:
		return BlacklistLeaderPolicyFromBytes(bytes[8:])
	default:
		return nil, fmt.Errorf("invalid LeaderSelectionPolicy type: %v", leaderPolicyType)
	}

}

// The SimpleLeaderPolicy is a trivial leader selection policy.
// It must be initialized with a set of node IDs and always returns that full set as leaders,
// regardless of which nodes have been suspected. In other words, each node is leader each epoch with this policy.
type SimpleLeaderPolicy struct {
	Membership []t.NodeID
}

// Leaders always returns the whole membership for the SimpleLeaderPolicy. All nodes are always leaders.
func (simple SimpleLeaderPolicy) Leaders() []t.NodeID {
	// All nodes are always leaders.
	return simple.Membership
}

// Suspect does nothing for the SimpleLeaderPolicy.
func (simple SimpleLeaderPolicy) Suspect(e t.EpochNr, node t.NodeID) {
	// Do nothing.
}

// Reconfigure informs the leader selection policy about a change in the membership.
func (simple SimpleLeaderPolicy) Reconfigure(nodeIDs []t.NodeID) LeaderSelectionPolicy {
	newPolicy := SimpleLeaderPolicy{Membership: make([]t.NodeID, len(nodeIDs))}
	copy(newPolicy.Membership, nodeIDs)
	return &newPolicy
}

func (simple *SimpleLeaderPolicy) Bytes() []byte {
	membershipBytes := t.NodeIDSliceBytes(simple.Membership)
	out := make([]byte, 0, len(membershipBytes)+8)
	out = append(out, t.Uint64ToBytes(uint64(Simple))...)
	return append(out, membershipBytes...)
}

func SimpleLeaderPolicyFromBytes(bytes []byte) *SimpleLeaderPolicy {
	return &SimpleLeaderPolicy{
		t.NodeIDSliceFromBytes(bytes),
	}
}

type BlacklistLeaderPolicy struct {
	Membership     map[t.NodeID]struct{}
	Suspected      map[t.NodeID]t.EpochNr
	minLeaders     int
	currentLeaders []t.NodeID
	updated        bool
}

func NewBlackListLeaderPolicy(members []t.NodeID, minLeaders int) *BlacklistLeaderPolicy {
	membership := make(map[t.NodeID]struct{}, len(members))
	for _, node := range members {
		membership[node] = struct{}{}
	}
	return &BlacklistLeaderPolicy{
		membership,
		make(map[t.NodeID]t.EpochNr, len(members)),
		minLeaders,
		members,
		true,
	}
}

// Leaders always returns the whole membership for the SimpleLeaderPolicy. All nodes are always leaders.
func (l *BlacklistLeaderPolicy) Leaders() []t.NodeID {
	// All nodes are always leaders.
	if l.updated { //no need to recalculate, no changes since last time
		return l.currentLeaders
	}
	// defer marking l.currentLeaders as updated
	defer func() {
		l.updated = true
	}()

	curLeaders := make([]t.NodeID, 0, len(l.Membership))
	for nodeID := range l.Membership {
		curLeaders = append(curLeaders, nodeID)
	}

	// Sort the values in ascending order of latest epoch where they each were suspected
	sort.Slice(curLeaders, func(i, j int) bool {
		_, oki := l.Suspected[curLeaders[i]]
		_, okj := l.Suspected[curLeaders[j]]
		// If both nodeIDs have never been suspected, then lexicographical order (both will be leaders)
		if !oki && !okj {
			return string(curLeaders[i]) < string(curLeaders[j])
		}

		// If one node is suspected and the other is not, the suspected one comes last
		if !oki && okj {
			return true
		}
		if oki && !okj {
			return false
		}

		// If both keys are in the suspected map, sort by value in the suspected map
		return l.Suspected[curLeaders[i]] < l.Suspected[curLeaders[j]]
	})

	// Calculate number of leaders
	leadersSize := len(l.Membership) - len(l.Suspected)
	if leadersSize < l.minLeaders {
		leadersSize = l.minLeaders // must at least return l.minLeaders
	}

	// return the leadersSize least recently suspected nodes as currentLeaders
	return curLeaders[:leadersSize]
}

// Suspect adds a new suspect to the list of suspects, or updates its epoch where it was suspected to the given epoch
// if this one is more recent than the one it already has
func (l *BlacklistLeaderPolicy) Suspect(e t.EpochNr, node t.NodeID) {
	if _, ok := l.Membership[node]; !ok { //node is a not a member
		//TODO error but cannot be passed through
		return
	}

	if epochNr, ok := l.Suspected[node]; !ok || epochNr < e {
		l.updated = false     // mark that next call to Leaders() needs to recalculate the leaders
		l.Suspected[node] = e // update with latest epoch that node was suspected
	}
}

// Reconfigure informs the leader selection policy about a change in the membership.
func (l *BlacklistLeaderPolicy) Reconfigure(nodeIDs []t.NodeID) LeaderSelectionPolicy {
	membership := make(map[t.NodeID]struct{}, len(nodeIDs))
	suspected := make(map[t.NodeID]t.EpochNr, len(nodeIDs))
	for _, nodeID := range nodeIDs {
		if sus, ok := l.Membership[nodeID]; ok {
			membership[nodeID] = sus // keep former suspect as suspected
			if epoch, ok := l.Suspected[nodeID]; ok {
				suspected[nodeID] = epoch
			}
		} else {
			membership[nodeID] = struct{}{}
		}
	}

	newPolicy := BlacklistLeaderPolicy{
		membership,
		suspected,
		l.minLeaders,
		nil,
		false,
	}

	return &newPolicy
}

func (l *BlacklistLeaderPolicy) Bytes() []byte {
	var membersBytes, _ = l.membersBytes()
	var suspectedBytes, _ = l.suspectedBytes()
	out := make([]byte, 0, len(membersBytes)+8+8)
	out = append(out, t.Uint64ToBytes(uint64(Blacklist))...)
	out = append(out, membersBytes...)
	out = append(out, suspectedBytes...)
	return append(out, t.Uint64ToBytes(uint64(l.minLeaders))...)
}

func (l *BlacklistLeaderPolicy) membersBytes() ([]byte, error) {
	return t.OrderedMapToBytes(l.Membership, t.NodeID(t.NodeIDSeparator).Bytes()[0], func(k t.NodeID, v struct{}) []byte {
		return (k + t.NodeIDSeparator).Bytes()
	})
}

func (l *BlacklistLeaderPolicy) suspectedBytes() ([]byte, error) {
	return t.OrderedMapToBytes(l.Suspected, t.NodeID(t.NodeIDSeparator).Bytes()[0], func(k t.NodeID, v t.EpochNr) []byte {
		out := (k + t.NodeIDSeparator).Bytes()
		out = append(out, v.Bytes()...)
		return out
	})
}

func BlacklistLeaderPolicyFromBytes(data []byte) (*BlacklistLeaderPolicy, error) {
	members, err, data := blacklistMembersFromBytes(data)
	if err != nil {
		return nil, err
	}
	suspects, err, data := blacklistSuspectsFromBytes(data)
	if err != nil {
		return nil, err
	}
	minLeaders := t.Uint64FromBytes(data)
	return &BlacklistLeaderPolicy{
		members,
		suspects,
		int(minLeaders),
		nil,
		false,
	}, nil
}

func blacklistMembersFromBytes(b []byte) (map[t.NodeID]struct{}, error, []byte) {
	separator := t.NodeID(t.NodeIDSeparator).Bytes()[0] // len(separator) = 1

	return t.BytesToMap[t.NodeID, struct{}](b, separator, func(b []byte, m *map[t.NodeID]struct{}) ([]byte, error, bool) {
		var i int
		curByte := b[0]

		for i = 1; curByte != separator && i < len(b)-1; i++ {
			curByte = b[i]
		}

		var end bool
		if i <= 1 && curByte == separator { // last separator
			if len(b) >= 1 {
				return b[i:], nil, true
			} else {
				return nil, nil, true
			}
		}

		if curByte == separator { // otherwise last separator
			(*m)[t.NodeID(string(b[0:i-1]))] = struct{}{}
		} else if i >= len(b)-1 { //b[i] != separator
			return nil, fmt.Errorf("unexpected end of []byte slice"), end
		}

		if len(b[i-1:]) > 1 {
			return b[i:], nil, end
		} else { //correct behaviour
			return nil, nil, end
		}
	})
}

func blacklistSuspectsFromBytes(b []byte) (map[t.NodeID]t.EpochNr, error, []byte) {
	separator := t.NodeID(t.NodeIDSeparator).Bytes()[0] // len(separator) = 1
	return t.BytesToMap[t.NodeID, t.EpochNr](b, separator, func(b []byte, m *map[t.NodeID]t.EpochNr) ([]byte, error, bool) {
		var i int
		if len(b) == 0 {
			return b, fmt.Errorf("0-byte slice"), false
		}
		curByte := b[0]

		for i = 1; curByte != separator && i < len(b)-8; i++ {
			curByte = b[i]
		}

		var end bool
		if i <= 1 && curByte == separator { // last separator
			if len(b) >= 1 {
				return b[i:], nil, true
			} else {
				return nil, nil, true
			}
		}

		if curByte == separator { // otherwise last separator
			(*m)[t.NodeID(string(b[0:i-1]))] = t.EpochNr(t.Uint64FromBytes(b[i : i+8]))
		} else { //b[i] != separator
			return nil, fmt.Errorf("unexpected end of []byte slice"), end
		}

		if len(b[i:]) > 8 {
			return b[i+8:], nil, end
		} else { //correct behaviour
			return nil, nil, end
		}
	})
}
