package internalapi

import (
	"github.com/Metchain/Metblock/utils/logger"
	"github.com/Metchain/Metblock/utils/panics"
)

var log = logger.RegisterSubSystem("B-Concensus")
var spawn = panics.GoroutineWrapperFunc(log)
