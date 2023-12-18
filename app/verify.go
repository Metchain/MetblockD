package app

import (
	"github.com/Metchain/Metblock/commanager"
	"github.com/Metchain/Metblock/network/addressmanager"
	"github.com/Metchain/Metblock/protoserver"
)

func (metApp *metchainApp) Verify(ps *protoserver.NetAdapter, cnmager *commanager.ConnectionManager, addrmngr *addressmanager.AddressManager) error {

	log.Infof("Output", metApp.cfg.ActiveNetParams.GRPCSeeds)

	return nil
}
