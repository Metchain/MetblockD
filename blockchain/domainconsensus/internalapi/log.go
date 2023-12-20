package internalapi

import (
	"github.com/Metchain/MetblockD/utils/logger"
	"github.com/Metchain/MetblockD/utils/panics"
)

var log = logger.RegisterSubSystem("B-Concensus")
var spawn = panics.GoroutineWrapperFunc(log)
