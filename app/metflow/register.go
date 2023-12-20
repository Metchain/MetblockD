package metflow

import (
	"github.com/Metchain/MetblockD/app/metflow/addressexchange"
	"github.com/Metchain/MetblockD/app/protocol/common"
	peerpkg "github.com/Metchain/MetblockD/app/protocol/peer"
	"github.com/Metchain/MetblockD/appmessage"
	"github.com/Metchain/MetblockD/protoserver/routerpkg"
)

func registerAddressFlows(m protocolManager, router *routerpkg.Router, isStopping *uint32, errChan chan error) []*common.Flow {
	outgoingRoute := router.OutgoingRoute()

	return []*common.Flow{
		m.RegisterFlow("SendAddresses", router, []appmessage.MessageCommand{appmessage.CmdRequestAddresses}, isStopping, errChan,
			func(incomingRoute *routerpkg.Route, peer *peerpkg.Peer) error {
				return addressexchange.SendAddresses(m.Context(), incomingRoute, outgoingRoute)
			},
		),

		m.RegisterOneTimeFlow("ReceiveAddresses", router, []appmessage.MessageCommand{appmessage.CmdAddresses}, isStopping, errChan,
			func(incomingRoute *routerpkg.Route, peer *peerpkg.Peer) error {
				return addressexchange.ReceiveAddresses(m.Context(), incomingRoute, outgoingRoute, peer)
			},
		),
	}
}
