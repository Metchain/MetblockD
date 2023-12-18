package domain

import (
	"github.com/Metchain/Metblock/blockchain/consensus"
	"github.com/Metchain/Metblock/blockchain/consensus/consensusreference"
	"github.com/Metchain/Metblock/blockchain/domainconsensus/miningmanager"
	"github.com/Metchain/Metblock/blockchain/mempool"
	"github.com/Metchain/Metblock/db/database"
	"github.com/Metchain/Metblock/external"
)

// Domain provides a reference to the domain's external aps
type Domain interface {
	Consensus() external.Consensus

	ConsensusEventsChannel() chan external.ConsensusEvent

	MiningManager() miningmanager.MiningManager
}

type domain struct {
	consensus *external.Consensus

	db                     database.Database
	consensusEventsChannel chan external.ConsensusEvent
	miningManager          miningmanager.MiningManager
}

func New(consensusConfig *consensus.Config, mempoolConfig *mempool.Config, db database.Database) (Domain, error) {

	consensusEventsChan := make(chan external.ConsensusEvent, 100e3)
	consensusFactory := consensus.NewFactory()
	consensusInstance, err := consensusFactory.NewConsensus(consensusConfig, db, consensusEventsChan)

	if err != nil {
		return nil, err
	}
	miningManagerFactory := miningmanager.NewFactory()
	domainInstance := &domain{
		consensus:              &consensusInstance,
		db:                     db,
		consensusEventsChannel: consensusEventsChan,
	}
	domainInstance.miningManager = miningManagerFactory.NewMiningManager(consensusreference.NewConsensusReference(&domainInstance.consensus), &consensusConfig.Params, mempoolConfig)

	return domainInstance, nil
}

func (d *domain) Consensus() external.Consensus {
	return *d.consensus
}

func (d *domain) ConsensusEventsChannel() chan external.ConsensusEvent {
	return d.consensusEventsChannel
}

func (d *domain) MiningManager() miningmanager.MiningManager {
	return d.miningManager
}
