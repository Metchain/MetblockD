package protocol

import (
	"github.com/Metchain/MetblockD/utils/logger"
	"github.com/Metchain/MetblockD/utils/panics"
)

var log = logger.RegisterSubSystem("METD-Portocol")
var spawn = panics.GoroutineWrapperFunc(log)
