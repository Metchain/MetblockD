package consensus

import (
	"github.com/Metchain/MetblockD/app/model"
	"github.com/Metchain/MetblockD/external"
)

type PastMedianTimeManagerConstructor func(int, model.DBReader, model.BlockHeaderStore, *external.DomainHash) model.PastMedianTimeManager
