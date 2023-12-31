// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.3
// source: message.proto

package protowire

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MetchainMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Payload:
	//
	//	*MetchainMessage_NotifyNewBlockTemplateRequest
	//	*MetchainMessage_NotifyNewBlockTemplateResponse
	//	*MetchainMessage_NewBlockTemplateNotification
	//	*MetchainMessage_GetBlockTemplateRequest
	//	*MetchainMessage_GetBlockTemplateResponse
	//	*MetchainMessage_GetBlockDagInfoRequest
	//	*MetchainMessage_GetBlockDagInfoResponse
	//	*MetchainMessage_EstimateNetworkHashesPerSecondRequest
	//	*MetchainMessage_EstimateNetworkHashesPerSecondResponse
	//	*MetchainMessage_GetBalanceByAddressRequest
	//	*MetchainMessage_GetBalanceByAddressResponse
	//	*MetchainMessage_SubmitBlockRequest
	//	*MetchainMessage_SubmitBlockResponse
	//	*MetchainMessage_P2PLatestBlockHeightRequest
	//	*MetchainMessage_P2PLatestBlockHeightResponse
	//	*MetchainMessage_P2PBlockWithTrustedDataRequestMessage
	//	*MetchainMessage_P2PBlockWithTrustedDataResponseMessage
	//	*MetchainMessage_SubmitSignedTXRequestMessage
	//	*MetchainMessage_SubmitSignedTXResponseMessage
	Payload isMetchainMessage_Payload `protobuf_oneof:"payload"`
}

func (x *MetchainMessage) Reset() {
	*x = MetchainMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetchainMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetchainMessage) ProtoMessage() {}

func (x *MetchainMessage) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetchainMessage.ProtoReflect.Descriptor instead.
func (*MetchainMessage) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (m *MetchainMessage) GetPayload() isMetchainMessage_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *MetchainMessage) GetNotifyNewBlockTemplateRequest() *NotifyNewBlockTemplateRequestMessage {
	if x, ok := x.GetPayload().(*MetchainMessage_NotifyNewBlockTemplateRequest); ok {
		return x.NotifyNewBlockTemplateRequest
	}
	return nil
}

func (x *MetchainMessage) GetNotifyNewBlockTemplateResponse() *NotifyNewBlockTemplateResponseMessage {
	if x, ok := x.GetPayload().(*MetchainMessage_NotifyNewBlockTemplateResponse); ok {
		return x.NotifyNewBlockTemplateResponse
	}
	return nil
}

func (x *MetchainMessage) GetNewBlockTemplateNotification() *NewBlockTemplateNotificationMessage {
	if x, ok := x.GetPayload().(*MetchainMessage_NewBlockTemplateNotification); ok {
		return x.NewBlockTemplateNotification
	}
	return nil
}

func (x *MetchainMessage) GetGetBlockTemplateRequest() *GetBlockTemplateRequestMessage {
	if x, ok := x.GetPayload().(*MetchainMessage_GetBlockTemplateRequest); ok {
		return x.GetBlockTemplateRequest
	}
	return nil
}

func (x *MetchainMessage) GetGetBlockTemplateResponse() *GetBlockTemplateResponseMessage {
	if x, ok := x.GetPayload().(*MetchainMessage_GetBlockTemplateResponse); ok {
		return x.GetBlockTemplateResponse
	}
	return nil
}

func (x *MetchainMessage) GetGetBlockDagInfoRequest() *GetBlockDagInfoRequestMessage {
	if x, ok := x.GetPayload().(*MetchainMessage_GetBlockDagInfoRequest); ok {
		return x.GetBlockDagInfoRequest
	}
	return nil
}

func (x *MetchainMessage) GetGetBlockDagInfoResponse() *GetBlockDagInfoResponseMessage {
	if x, ok := x.GetPayload().(*MetchainMessage_GetBlockDagInfoResponse); ok {
		return x.GetBlockDagInfoResponse
	}
	return nil
}

func (x *MetchainMessage) GetEstimateNetworkHashesPerSecondRequest() *EstimateNetworkHashesPerSecondRequestMessage {
	if x, ok := x.GetPayload().(*MetchainMessage_EstimateNetworkHashesPerSecondRequest); ok {
		return x.EstimateNetworkHashesPerSecondRequest
	}
	return nil
}

func (x *MetchainMessage) GetEstimateNetworkHashesPerSecondResponse() *EstimateNetworkHashesPerSecondResponseMessage {
	if x, ok := x.GetPayload().(*MetchainMessage_EstimateNetworkHashesPerSecondResponse); ok {
		return x.EstimateNetworkHashesPerSecondResponse
	}
	return nil
}

func (x *MetchainMessage) GetGetBalanceByAddressRequest() *GetBalanceByAddressRequestMessage {
	if x, ok := x.GetPayload().(*MetchainMessage_GetBalanceByAddressRequest); ok {
		return x.GetBalanceByAddressRequest
	}
	return nil
}

func (x *MetchainMessage) GetGetBalanceByAddressResponse() *GetBalanceByAddressResponseMessage {
	if x, ok := x.GetPayload().(*MetchainMessage_GetBalanceByAddressResponse); ok {
		return x.GetBalanceByAddressResponse
	}
	return nil
}

func (x *MetchainMessage) GetSubmitBlockRequest() *SubmitBlockRequestMessage {
	if x, ok := x.GetPayload().(*MetchainMessage_SubmitBlockRequest); ok {
		return x.SubmitBlockRequest
	}
	return nil
}

func (x *MetchainMessage) GetSubmitBlockResponse() *SubmitBlockResponseMessage {
	if x, ok := x.GetPayload().(*MetchainMessage_SubmitBlockResponse); ok {
		return x.SubmitBlockResponse
	}
	return nil
}

func (x *MetchainMessage) GetP2PLatestBlockHeightRequest() *P2PLatestBlockHeightRequestMessge {
	if x, ok := x.GetPayload().(*MetchainMessage_P2PLatestBlockHeightRequest); ok {
		return x.P2PLatestBlockHeightRequest
	}
	return nil
}

func (x *MetchainMessage) GetP2PLatestBlockHeightResponse() *P2PLatestBlockHeightResponseMessage {
	if x, ok := x.GetPayload().(*MetchainMessage_P2PLatestBlockHeightResponse); ok {
		return x.P2PLatestBlockHeightResponse
	}
	return nil
}

func (x *MetchainMessage) GetP2PBlockWithTrustedDataRequestMessage() *P2PBlockWithTrustedDataRequestMessage {
	if x, ok := x.GetPayload().(*MetchainMessage_P2PBlockWithTrustedDataRequestMessage); ok {
		return x.P2PBlockWithTrustedDataRequestMessage
	}
	return nil
}

func (x *MetchainMessage) GetP2PBlockWithTrustedDataResponseMessage() *P2PBlockWithTrustedDataResponseMessage {
	if x, ok := x.GetPayload().(*MetchainMessage_P2PBlockWithTrustedDataResponseMessage); ok {
		return x.P2PBlockWithTrustedDataResponseMessage
	}
	return nil
}

func (x *MetchainMessage) GetSubmitSignedTXRequestMessage() *SubmitSignedTXRequestMessage {
	if x, ok := x.GetPayload().(*MetchainMessage_SubmitSignedTXRequestMessage); ok {
		return x.SubmitSignedTXRequestMessage
	}
	return nil
}

func (x *MetchainMessage) GetSubmitSignedTXResponseMessage() *SubmitSignedTXResponseMessage {
	if x, ok := x.GetPayload().(*MetchainMessage_SubmitSignedTXResponseMessage); ok {
		return x.SubmitSignedTXResponseMessage
	}
	return nil
}

type isMetchainMessage_Payload interface {
	isMetchainMessage_Payload()
}

type MetchainMessage_NotifyNewBlockTemplateRequest struct {
	NotifyNewBlockTemplateRequest *NotifyNewBlockTemplateRequestMessage `protobuf:"bytes,1,opt,name=notifyNewBlockTemplateRequest,proto3,oneof"`
}

type MetchainMessage_NotifyNewBlockTemplateResponse struct {
	NotifyNewBlockTemplateResponse *NotifyNewBlockTemplateResponseMessage `protobuf:"bytes,2,opt,name=notifyNewBlockTemplateResponse,proto3,oneof"`
}

type MetchainMessage_NewBlockTemplateNotification struct {
	NewBlockTemplateNotification *NewBlockTemplateNotificationMessage `protobuf:"bytes,3,opt,name=newBlockTemplateNotification,proto3,oneof"`
}

type MetchainMessage_GetBlockTemplateRequest struct {
	GetBlockTemplateRequest *GetBlockTemplateRequestMessage `protobuf:"bytes,4,opt,name=getBlockTemplateRequest,proto3,oneof"`
}

type MetchainMessage_GetBlockTemplateResponse struct {
	GetBlockTemplateResponse *GetBlockTemplateResponseMessage `protobuf:"bytes,5,opt,name=getBlockTemplateResponse,proto3,oneof"`
}

type MetchainMessage_GetBlockDagInfoRequest struct {
	GetBlockDagInfoRequest *GetBlockDagInfoRequestMessage `protobuf:"bytes,6,opt,name=getBlockDagInfoRequest,proto3,oneof"`
}

type MetchainMessage_GetBlockDagInfoResponse struct {
	GetBlockDagInfoResponse *GetBlockDagInfoResponseMessage `protobuf:"bytes,7,opt,name=getBlockDagInfoResponse,proto3,oneof"`
}

type MetchainMessage_EstimateNetworkHashesPerSecondRequest struct {
	EstimateNetworkHashesPerSecondRequest *EstimateNetworkHashesPerSecondRequestMessage `protobuf:"bytes,8,opt,name=estimateNetworkHashesPerSecondRequest,proto3,oneof"`
}

type MetchainMessage_EstimateNetworkHashesPerSecondResponse struct {
	EstimateNetworkHashesPerSecondResponse *EstimateNetworkHashesPerSecondResponseMessage `protobuf:"bytes,9,opt,name=estimateNetworkHashesPerSecondResponse,proto3,oneof"`
}

type MetchainMessage_GetBalanceByAddressRequest struct {
	GetBalanceByAddressRequest *GetBalanceByAddressRequestMessage `protobuf:"bytes,10,opt,name=getBalanceByAddressRequest,proto3,oneof"`
}

type MetchainMessage_GetBalanceByAddressResponse struct {
	GetBalanceByAddressResponse *GetBalanceByAddressResponseMessage `protobuf:"bytes,11,opt,name=getBalanceByAddressResponse,proto3,oneof"`
}

type MetchainMessage_SubmitBlockRequest struct {
	SubmitBlockRequest *SubmitBlockRequestMessage `protobuf:"bytes,12,opt,name=submitBlockRequest,proto3,oneof"`
}

type MetchainMessage_SubmitBlockResponse struct {
	SubmitBlockResponse *SubmitBlockResponseMessage `protobuf:"bytes,13,opt,name=submitBlockResponse,proto3,oneof"`
}

type MetchainMessage_P2PLatestBlockHeightRequest struct {
	P2PLatestBlockHeightRequest *P2PLatestBlockHeightRequestMessge `protobuf:"bytes,14,opt,name=P2PLatestBlockHeightRequest,proto3,oneof"`
}

type MetchainMessage_P2PLatestBlockHeightResponse struct {
	P2PLatestBlockHeightResponse *P2PLatestBlockHeightResponseMessage `protobuf:"bytes,15,opt,name=P2PLatestBlockHeightResponse,proto3,oneof"`
}

type MetchainMessage_P2PBlockWithTrustedDataRequestMessage struct {
	P2PBlockWithTrustedDataRequestMessage *P2PBlockWithTrustedDataRequestMessage `protobuf:"bytes,16,opt,name=P2PBlockWithTrustedDataRequestMessage,proto3,oneof"`
}

type MetchainMessage_P2PBlockWithTrustedDataResponseMessage struct {
	P2PBlockWithTrustedDataResponseMessage *P2PBlockWithTrustedDataResponseMessage `protobuf:"bytes,17,opt,name=P2PBlockWithTrustedDataResponseMessage,proto3,oneof"`
}

type MetchainMessage_SubmitSignedTXRequestMessage struct {
	SubmitSignedTXRequestMessage *SubmitSignedTXRequestMessage `protobuf:"bytes,18,opt,name=submitSignedTXRequestMessage,proto3,oneof"`
}

type MetchainMessage_SubmitSignedTXResponseMessage struct {
	SubmitSignedTXResponseMessage *SubmitSignedTXResponseMessage `protobuf:"bytes,19,opt,name=submitSignedTXResponseMessage,proto3,oneof"`
}

func (*MetchainMessage_NotifyNewBlockTemplateRequest) isMetchainMessage_Payload() {}

func (*MetchainMessage_NotifyNewBlockTemplateResponse) isMetchainMessage_Payload() {}

func (*MetchainMessage_NewBlockTemplateNotification) isMetchainMessage_Payload() {}

func (*MetchainMessage_GetBlockTemplateRequest) isMetchainMessage_Payload() {}

func (*MetchainMessage_GetBlockTemplateResponse) isMetchainMessage_Payload() {}

func (*MetchainMessage_GetBlockDagInfoRequest) isMetchainMessage_Payload() {}

func (*MetchainMessage_GetBlockDagInfoResponse) isMetchainMessage_Payload() {}

func (*MetchainMessage_EstimateNetworkHashesPerSecondRequest) isMetchainMessage_Payload() {}

func (*MetchainMessage_EstimateNetworkHashesPerSecondResponse) isMetchainMessage_Payload() {}

func (*MetchainMessage_GetBalanceByAddressRequest) isMetchainMessage_Payload() {}

func (*MetchainMessage_GetBalanceByAddressResponse) isMetchainMessage_Payload() {}

func (*MetchainMessage_SubmitBlockRequest) isMetchainMessage_Payload() {}

func (*MetchainMessage_SubmitBlockResponse) isMetchainMessage_Payload() {}

func (*MetchainMessage_P2PLatestBlockHeightRequest) isMetchainMessage_Payload() {}

func (*MetchainMessage_P2PLatestBlockHeightResponse) isMetchainMessage_Payload() {}

func (*MetchainMessage_P2PBlockWithTrustedDataRequestMessage) isMetchainMessage_Payload() {}

func (*MetchainMessage_P2PBlockWithTrustedDataResponseMessage) isMetchainMessage_Payload() {}

func (*MetchainMessage_SubmitSignedTXRequestMessage) isMetchainMessage_Payload() {}

func (*MetchainMessage_SubmitSignedTXResponseMessage) isMetchainMessage_Payload() {}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x77, 0x69, 0x72, 0x65, 0x1a, 0x09, 0x72, 0x70, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x70, 0x32, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xc2, 0x11, 0x0a, 0x0f, 0x4d, 0x65, 0x74, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x77, 0x0a, 0x1d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4e, 0x65,
	0x77, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x77, 0x69, 0x72, 0x65, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4e, 0x65,
	0x77, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x1d,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4e, 0x65, 0x77, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x7a, 0x0a,
	0x1e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4e, 0x65, 0x77, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x77, 0x69, 0x72,
	0x65, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4e, 0x65, 0x77, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x1e, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x4e, 0x65, 0x77, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x74, 0x0a, 0x1c, 0x6e, 0x65, 0x77,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x77, 0x69, 0x72, 0x65, 0x2e, 0x4e, 0x65, 0x77, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48,
	0x00, 0x52, 0x1c, 0x6e, 0x65, 0x77, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x65, 0x0a, 0x17, 0x67, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x77, 0x69, 0x72, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x17, 0x67,
	0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x68, 0x0a, 0x18, 0x67, 0x65, 0x74, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x77, 0x69, 0x72, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x18, 0x67, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x62, 0x0a, 0x16, 0x67, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x61, 0x67, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x77, 0x69, 0x72, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x61, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x16, 0x67, 0x65,
	0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x61, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x65, 0x0a, 0x17, 0x67, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x44, 0x61, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x77, 0x69, 0x72,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x61, 0x67, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x48, 0x00, 0x52, 0x17, 0x67, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x61, 0x67, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x8f, 0x01, 0x0a, 0x25,
	0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x48,
	0x61, 0x73, 0x68, 0x65, 0x73, 0x50, 0x65, 0x72, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x77, 0x69, 0x72, 0x65, 0x2e, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x50, 0x65, 0x72,
	0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x25, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x50, 0x65, 0x72,
	0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x92, 0x01,
	0x0a, 0x26, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x50, 0x65, 0x72, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x77, 0x69, 0x72, 0x65, 0x2e, 0x45, 0x73, 0x74, 0x69, 0x6d,
	0x61, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73,
	0x50, 0x65, 0x72, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x26, 0x65, 0x73, 0x74, 0x69,
	0x6d, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x65,
	0x73, 0x50, 0x65, 0x72, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x6e, 0x0a, 0x1a, 0x67, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x77, 0x69,
	0x72, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x79, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x1a, 0x67, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x71, 0x0a, 0x1b, 0x67, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x77,
	0x69, 0x72, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x79,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x1b, 0x67, 0x65, 0x74, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x12, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x77, 0x69, 0x72, 0x65, 0x2e, 0x53, 0x75,
	0x62, 0x6d, 0x69, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x12, 0x73, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x59, 0x0a,
	0x13, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x77, 0x69, 0x72, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x48, 0x00, 0x52, 0x13, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x70, 0x0a, 0x1b, 0x50, 0x32, 0x50, 0x4c,
	0x61, 0x74, 0x65, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x77, 0x69, 0x72, 0x65, 0x2e, 0x50, 0x32, 0x50, 0x4c, 0x61, 0x74,
	0x65, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x67, 0x65, 0x48, 0x00, 0x52, 0x1b, 0x50,
	0x32, 0x50, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x74, 0x0a, 0x1c, 0x50, 0x32,
	0x50, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x77, 0x69, 0x72, 0x65, 0x2e, 0x50, 0x32, 0x50,
	0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x48, 0x00, 0x52, 0x1c, 0x50, 0x32, 0x50, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x88, 0x01, 0x0a, 0x25, 0x50, 0x32, 0x50, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x57, 0x69, 0x74,
	0x68, 0x54, 0x72, 0x75, 0x73, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x30, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x77, 0x69, 0x72, 0x65, 0x2e, 0x50, 0x32, 0x50,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x57, 0x69, 0x74, 0x68, 0x54, 0x72, 0x75, 0x73, 0x74, 0x65, 0x64,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x48, 0x00, 0x52, 0x25, 0x50, 0x32, 0x50, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x57, 0x69,
	0x74, 0x68, 0x54, 0x72, 0x75, 0x73, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x8b, 0x01, 0x0a, 0x26,
	0x50, 0x32, 0x50, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x57, 0x69, 0x74, 0x68, 0x54, 0x72, 0x75, 0x73,
	0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x77, 0x69, 0x72, 0x65, 0x2e, 0x50, 0x32, 0x50, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x57, 0x69, 0x74, 0x68, 0x54, 0x72, 0x75, 0x73, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48,
	0x00, 0x52, 0x26, 0x50, 0x32, 0x50, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x57, 0x69, 0x74, 0x68, 0x54,
	0x72, 0x75, 0x73, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x6d, 0x0a, 0x1c, 0x73, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x58, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x77, 0x69, 0x72, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x58, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x1c, 0x73, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x58, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x70, 0x0a, 0x1d, 0x73, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x58, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x77, 0x69, 0x72, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x58, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x1d, 0x73, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x58, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x32, 0x54, 0x0a, 0x03, 0x52, 0x50, 0x43, 0x12, 0x4d, 0x0a, 0x0d,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x77, 0x69, 0x72, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x77, 0x69, 0x72, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x32, 0x54, 0x0a, 0x03, 0x50,
	0x32, 0x50, 0x12, 0x4d, 0x0a, 0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x77, 0x69, 0x72, 0x65, 0x2e,
	0x4d, 0x65, 0x74, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x77, 0x69, 0x72, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30,
	0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_message_proto_goTypes = []interface{}{
	(*MetchainMessage)(nil),                               // 0: protowire.MetchainMessage
	(*NotifyNewBlockTemplateRequestMessage)(nil),          // 1: protowire.NotifyNewBlockTemplateRequestMessage
	(*NotifyNewBlockTemplateResponseMessage)(nil),         // 2: protowire.NotifyNewBlockTemplateResponseMessage
	(*NewBlockTemplateNotificationMessage)(nil),           // 3: protowire.NewBlockTemplateNotificationMessage
	(*GetBlockTemplateRequestMessage)(nil),                // 4: protowire.GetBlockTemplateRequestMessage
	(*GetBlockTemplateResponseMessage)(nil),               // 5: protowire.GetBlockTemplateResponseMessage
	(*GetBlockDagInfoRequestMessage)(nil),                 // 6: protowire.GetBlockDagInfoRequestMessage
	(*GetBlockDagInfoResponseMessage)(nil),                // 7: protowire.GetBlockDagInfoResponseMessage
	(*EstimateNetworkHashesPerSecondRequestMessage)(nil),  // 8: protowire.EstimateNetworkHashesPerSecondRequestMessage
	(*EstimateNetworkHashesPerSecondResponseMessage)(nil), // 9: protowire.EstimateNetworkHashesPerSecondResponseMessage
	(*GetBalanceByAddressRequestMessage)(nil),             // 10: protowire.GetBalanceByAddressRequestMessage
	(*GetBalanceByAddressResponseMessage)(nil),            // 11: protowire.GetBalanceByAddressResponseMessage
	(*SubmitBlockRequestMessage)(nil),                     // 12: protowire.SubmitBlockRequestMessage
	(*SubmitBlockResponseMessage)(nil),                    // 13: protowire.SubmitBlockResponseMessage
	(*P2PLatestBlockHeightRequestMessge)(nil),             // 14: protowire.P2PLatestBlockHeightRequestMessge
	(*P2PLatestBlockHeightResponseMessage)(nil),           // 15: protowire.P2PLatestBlockHeightResponseMessage
	(*P2PBlockWithTrustedDataRequestMessage)(nil),         // 16: protowire.P2PBlockWithTrustedDataRequestMessage
	(*P2PBlockWithTrustedDataResponseMessage)(nil),        // 17: protowire.P2PBlockWithTrustedDataResponseMessage
	(*SubmitSignedTXRequestMessage)(nil),                  // 18: protowire.SubmitSignedTXRequestMessage
	(*SubmitSignedTXResponseMessage)(nil),                 // 19: protowire.SubmitSignedTXResponseMessage
}
var file_message_proto_depIdxs = []int32{
	1,  // 0: protowire.MetchainMessage.notifyNewBlockTemplateRequest:type_name -> protowire.NotifyNewBlockTemplateRequestMessage
	2,  // 1: protowire.MetchainMessage.notifyNewBlockTemplateResponse:type_name -> protowire.NotifyNewBlockTemplateResponseMessage
	3,  // 2: protowire.MetchainMessage.newBlockTemplateNotification:type_name -> protowire.NewBlockTemplateNotificationMessage
	4,  // 3: protowire.MetchainMessage.getBlockTemplateRequest:type_name -> protowire.GetBlockTemplateRequestMessage
	5,  // 4: protowire.MetchainMessage.getBlockTemplateResponse:type_name -> protowire.GetBlockTemplateResponseMessage
	6,  // 5: protowire.MetchainMessage.getBlockDagInfoRequest:type_name -> protowire.GetBlockDagInfoRequestMessage
	7,  // 6: protowire.MetchainMessage.getBlockDagInfoResponse:type_name -> protowire.GetBlockDagInfoResponseMessage
	8,  // 7: protowire.MetchainMessage.estimateNetworkHashesPerSecondRequest:type_name -> protowire.EstimateNetworkHashesPerSecondRequestMessage
	9,  // 8: protowire.MetchainMessage.estimateNetworkHashesPerSecondResponse:type_name -> protowire.EstimateNetworkHashesPerSecondResponseMessage
	10, // 9: protowire.MetchainMessage.getBalanceByAddressRequest:type_name -> protowire.GetBalanceByAddressRequestMessage
	11, // 10: protowire.MetchainMessage.getBalanceByAddressResponse:type_name -> protowire.GetBalanceByAddressResponseMessage
	12, // 11: protowire.MetchainMessage.submitBlockRequest:type_name -> protowire.SubmitBlockRequestMessage
	13, // 12: protowire.MetchainMessage.submitBlockResponse:type_name -> protowire.SubmitBlockResponseMessage
	14, // 13: protowire.MetchainMessage.P2PLatestBlockHeightRequest:type_name -> protowire.P2PLatestBlockHeightRequestMessge
	15, // 14: protowire.MetchainMessage.P2PLatestBlockHeightResponse:type_name -> protowire.P2PLatestBlockHeightResponseMessage
	16, // 15: protowire.MetchainMessage.P2PBlockWithTrustedDataRequestMessage:type_name -> protowire.P2PBlockWithTrustedDataRequestMessage
	17, // 16: protowire.MetchainMessage.P2PBlockWithTrustedDataResponseMessage:type_name -> protowire.P2PBlockWithTrustedDataResponseMessage
	18, // 17: protowire.MetchainMessage.submitSignedTXRequestMessage:type_name -> protowire.SubmitSignedTXRequestMessage
	19, // 18: protowire.MetchainMessage.submitSignedTXResponseMessage:type_name -> protowire.SubmitSignedTXResponseMessage
	0,  // 19: protowire.RPC.MessageStream:input_type -> protowire.MetchainMessage
	0,  // 20: protowire.P2P.MessageStream:input_type -> protowire.MetchainMessage
	0,  // 21: protowire.RPC.MessageStream:output_type -> protowire.MetchainMessage
	0,  // 22: protowire.P2P.MessageStream:output_type -> protowire.MetchainMessage
	21, // [21:23] is the sub-list for method output_type
	19, // [19:21] is the sub-list for method input_type
	19, // [19:19] is the sub-list for extension type_name
	19, // [19:19] is the sub-list for extension extendee
	0,  // [0:19] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	file_rpc_proto_init()
	file_p2p_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetchainMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_message_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*MetchainMessage_NotifyNewBlockTemplateRequest)(nil),
		(*MetchainMessage_NotifyNewBlockTemplateResponse)(nil),
		(*MetchainMessage_NewBlockTemplateNotification)(nil),
		(*MetchainMessage_GetBlockTemplateRequest)(nil),
		(*MetchainMessage_GetBlockTemplateResponse)(nil),
		(*MetchainMessage_GetBlockDagInfoRequest)(nil),
		(*MetchainMessage_GetBlockDagInfoResponse)(nil),
		(*MetchainMessage_EstimateNetworkHashesPerSecondRequest)(nil),
		(*MetchainMessage_EstimateNetworkHashesPerSecondResponse)(nil),
		(*MetchainMessage_GetBalanceByAddressRequest)(nil),
		(*MetchainMessage_GetBalanceByAddressResponse)(nil),
		(*MetchainMessage_SubmitBlockRequest)(nil),
		(*MetchainMessage_SubmitBlockResponse)(nil),
		(*MetchainMessage_P2PLatestBlockHeightRequest)(nil),
		(*MetchainMessage_P2PLatestBlockHeightResponse)(nil),
		(*MetchainMessage_P2PBlockWithTrustedDataRequestMessage)(nil),
		(*MetchainMessage_P2PBlockWithTrustedDataResponseMessage)(nil),
		(*MetchainMessage_SubmitSignedTXRequestMessage)(nil),
		(*MetchainMessage_SubmitSignedTXResponseMessage)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}
