package consensus

import (
	"github.com/Metchain/MetblockD/external"
)

func (c *consensus) BuildBlockTemplate(coinbaseData *external.DomainCoinbaseData) (*external.DomainBlockTemplate, error) {

	c.lock.Lock()
	defer c.lock.Unlock()

	block, _, err := c.blockBuilder.BuildBlockTemplate(coinbaseData, []*external.DomainTransaction{})

	if err != nil {
		return nil, err
	}

	isNearlySynced, err := c.isNearlySyncedNoLock()
	if err != nil {
		return nil, err
	}

	return &external.DomainBlockTemplate{
		Block:        block,
		CoinbaseData: coinbaseData,

		IsNearlySynced: isNearlySynced,
	}, nil

}

func (s *consensus) isNearlySyncedNoLock() (bool, error) {
	/*stagingArea := model.NewStagingArea()
	virtualGHOSTDAGData, err := s.ghostdagDataStores[0].Get(s.databaseContext, stagingArea, model.VirtualBlockHash, false)
	if err != nil {
		return false, err
	}

	if virtualGHOSTDAGData.SelectedParent().Equal(s.genesisHash) {
		return false, nil
	}*/

	return true, nil
}
