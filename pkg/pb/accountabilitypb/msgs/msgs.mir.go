// Code generated by Mir codegen. DO NOT EDIT.

package accountabilitypbmsgs

import (
	types2 "github.com/filecoin-project/mir/pkg/pb/accountabilitypb/types"
	types3 "github.com/filecoin-project/mir/pkg/pb/isspb/types"
	types1 "github.com/filecoin-project/mir/pkg/pb/messagepb/types"
	types "github.com/filecoin-project/mir/pkg/types"
)

func SignedPredecision(destModule types.ModuleID, predecision []uint8, signature []uint8) *types1.Message {
	return &types1.Message{
		DestModule: destModule,
		Type: &types1.Message_Accountability{
			Accountability: &types2.Message{
				Type: &types2.Message_SignedPredecision{
					SignedPredecision: &types2.SignedPredecision{
						Predecision: predecision,
						Signature:   signature,
					},
				},
			},
		},
	}
}

func FullCertificate(destModule types.ModuleID, decision []uint8, signatures map[types.NodeID][]uint8) *types1.Message {
	return &types1.Message{
		DestModule: destModule,
		Type: &types1.Message_Accountability{
			Accountability: &types2.Message{
				Type: &types2.Message_Certificate{
					Certificate: &types2.FullCertificate{
						Decision:   decision,
						Signatures: signatures,
					},
				},
			},
		},
	}
}

func PoMs(destModule types.ModuleID, poms []*types2.PoM) *types1.Message {
	return &types1.Message{
		DestModule: destModule,
		Type: &types1.Message_Accountability{
			Accountability: &types2.Message{
				Type: &types2.Message_Poms{
					Poms: &types2.PoMs{
						Poms: poms,
					},
				},
			},
		},
	}
}

func LightCertificate(destModule types.ModuleID, data []uint8) *types1.Message {
	return &types1.Message{
		DestModule: destModule,
		Type: &types1.Message_Accountability{
			Accountability: &types2.Message{
				Type: &types2.Message_LightCertificate{
					LightCertificate: &types2.LightCertificate{
						Data: data,
					},
				},
			},
		},
	}
}

func RequestSBMessage(destModule types.ModuleID, predecision []uint8) *types1.Message {
	return &types1.Message{
		DestModule: destModule,
		Type: &types1.Message_Accountability{
			Accountability: &types2.Message{
				Type: &types2.Message_RequestSbMessage{
					RequestSbMessage: &types2.RequestSBMessage{
						Predecision: predecision,
					},
				},
			},
		},
	}
}

func ProvideSBMessage(destModule types.ModuleID, sbDeliver *types3.SBDeliver) *types1.Message {
	return &types1.Message{
		DestModule: destModule,
		Type: &types1.Message_Accountability{
			Accountability: &types2.Message{
				Type: &types2.Message_ProvideSbMessage{
					ProvideSbMessage: &types2.ProvideSBMessage{
						SbDeliver: sbDeliver,
					},
				},
			},
		},
	}
}