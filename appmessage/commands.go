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
)
