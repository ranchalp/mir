// Code generated by Mir codegen. DO NOT EDIT.

package checkpointvaliditycheckerpbevents

import (
	types4 "github.com/filecoin-project/mir/pkg/pb/checkpointpb/checkpointvaliditycheckerpb/types"
	types1 "github.com/filecoin-project/mir/pkg/pb/checkpointpb/types"
	types5 "github.com/filecoin-project/mir/pkg/pb/eventpb/types"
	types3 "github.com/filecoin-project/mir/pkg/pb/trantorpb/types"
	types2 "github.com/filecoin-project/mir/pkg/trantor/types"
	types "github.com/filecoin-project/mir/pkg/types"
)

func ValidateCheckpoint(destModule types.ModuleID, checkpoint *types1.StableCheckpoint, epochNr types2.EpochNr, memberships []*types3.Membership, origin *types4.ValidateChkpOrigin) *types5.Event {
	return &types5.Event{
		DestModule: destModule,
		Type: &types5.Event_Cvc{
			Cvc: &types4.Event{
				Type: &types4.Event_ValidateCheckpoint{
					ValidateCheckpoint: &types4.ValidateCheckpoint{
						Checkpoint:  checkpoint,
						EpochNr:     epochNr,
						Memberships: memberships,
						Origin:      origin,
					},
				},
			},
		},
	}
}

func CheckpointValidated(destModule types.ModuleID, error error, origin *types4.ValidateChkpOrigin) *types5.Event {
	return &types5.Event{
		DestModule: destModule,
		Type: &types5.Event_Cvc{
			Cvc: &types4.Event{
				Type: &types4.Event_CheckpointValidated{
					CheckpointValidated: &types4.CheckpointValidated{
						Error:  error,
						Origin: origin,
					},
				},
			},
		},
	}
}
