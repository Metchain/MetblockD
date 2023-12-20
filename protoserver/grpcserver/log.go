package grpcserver

import (
	"github.com/Metchain/MetblockD/utils/logger"
	"github.com/Metchain/MetblockD/utils/panics"
)

var log = logger.RegisterSubSystem("TXMP")
var spawn = panics.GoroutineWrapperFunc(log)
