package consensus

import (
	"github.com/Metchain/Metblock/app/model"
	"github.com/Metchain/Metblock/external"
)

type PastMedianTimeManagerConstructor func(int, model.DBReader, model.BlockHeaderStore, *external.DomainHash) model.PastMedianTimeManager
