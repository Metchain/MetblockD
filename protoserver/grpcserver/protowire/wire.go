package protowire

import (
	"github.com/Metchain/Metblock/appmessage"
	"github.com/pkg/errors"
)

type converter interface {
	toAppMessage() (appmessage.Message, error)
}

func FromAppMessage(message appmessage.Message) (*MetchainMessage, error) {
	payload, err := toPayload(message)
	if err != nil {
		return nil, err
	}
	return &MetchainMessage{
		Payload: payload,
	}, nil
}

func (x *MetchainMessage) ToAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "MetchainD Message is nil")
	}
	converter, ok := x.Payload.(converter)
	if !ok {
		return nil, errors.Errorf("received invalid message")
	}
	appMessage, err := converter.toAppMessage()
	if err != nil {
		return nil, err
	}
	return appMessage, nil
}

func toPayload(message appmessage.Message) (isMetchainMessage_Payload, error) {
	p2pPayload, err := toP2PPayload(message)
	if err != nil {
		return nil, err
	}
	if p2pPayload != nil {
		return p2pPayload, nil
	}

	rpcPayload, err := toRPCPayload(message)
	if err != nil {
		return nil, err
	}
	if rpcPayload != nil {
		return rpcPayload, nil
	}

	return nil, errors.Errorf("unknown message type %T", message)
}

func toP2PPayload(message appmessage.Message) (isMetchainMessage_Payload, error) {

	return nil, nil

}

func toRPCPayload(message appmessage.Message) (isMetchainMessage_Payload, error) {
	switch message := message.(type) {

	case *appmessage.SubmitBlockRequestMessage:
		payload := new(MetchainMessage_SubmitBlockRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.SubmitBlockResponseMessage:
		payload := new(MetchainMessage_SubmitBlockResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlockTemplateRequestMessage:
		payload := new(MetchainMessage_GetBlockTemplateRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlockTemplateResponseMessage:
		payload := new(MetchainMessage_GetBlockTemplateResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil

	case *appmessage.GetBlockDAGInfoRequestMessage:
		payload := new(MetchainMessage_GetBlockDagInfoRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlockDAGInfoResponseMessage:
		payload := new(MetchainMessage_GetBlockDagInfoResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil

	case *appmessage.GetBalanceByAddressRequestMessage:
		payload := new(MetchainMessage_GetBalanceByAddressRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBalanceByAddressResponseMessage:
		payload := new(MetchainMessage_GetBalanceByAddressResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil

	case *appmessage.EstimateNetworkHashesPerSecondRequestMessage:
		payload := new(MetchainMessage_EstimateNetworkHashesPerSecondRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.EstimateNetworkHashesPerSecondResponseMessage:
		payload := new(MetchainMessage_EstimateNetworkHashesPerSecondResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil

	case *appmessage.NotifyNewBlockTemplateRequestMessage:
		payload := new(MetchainMessage_NotifyNewBlockTemplateRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyNewBlockTemplateResponseMessage:
		payload := new(MetchainMessage_NotifyNewBlockTemplateResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NewBlockTemplateNotificationMessage:
		payload := new(MetchainMessage_NewBlockTemplateNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil

	default:
		return nil, nil
	}
}
