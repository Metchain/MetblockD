package app

import (
	"github.com/Metchain/MetblockD/commanager"
	"github.com/Metchain/MetblockD/network/addressmanager"
	"github.com/Metchain/MetblockD/protoserver"
)

func (metApp *metchainApp) Verify(ps *protoserver.NetAdapter, cnmager *commanager.ConnectionManager, addrmngr *addressmanager.AddressManager) error {

	log.Infof("Output", metApp.cfg.ActiveNetParams.GRPCSeeds)

	return nil
}
