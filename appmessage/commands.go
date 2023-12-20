package appmessage

const (
	// protocol
	CmdVersion MessageCommand = iota
	CmdVerAck
	CmdRequestAddresses
	CmdAddresses
	CmdReady
	CmdReject
	CmdNotifyNewBlockTemplateResponseMessage
	CmdNewBlockTemplateNotificationMessage
	CmdNotifyNewBlockTemplateRequestMessage
	CmdBlock
	CmdGetBlockTemplateRequestMessage
	CmdSubmitBlockResponseMessage
	CmdSubmitBlockRequestMessage
	CmdGetBlockTemplateResponseMessage
	CmdGetBlockDAGInfoResponseMessage
	CmdGetBlockDAGInfoRequestMessage
	CmdGetBalanceByAddressResponseMessage
	CmdGetBalanceByAddressRequestMessage
	CmdEstimateNetworkHashesPerSecondResponseMessage
	CmdEstimateNetworkHashesPerSecondRequestMessage

	//RPC
	CmdGetInfoRequestMessage
	CmdGetInfoResponseMessage
)

// Add all RPC commands here
var RPCMessageCommandToString = map[MessageCommand]string{}
