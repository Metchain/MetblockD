package app

import (
	"github.com/Metchain/MetblockD/utils/logger"
	"github.com/Metchain/MetblockD/utils/panics"
)

var log = logger.RegisterSubSystem("METD")
var spawn = panics.GoroutineWrapperFunc(log)
