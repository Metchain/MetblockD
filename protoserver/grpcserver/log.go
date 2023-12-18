package grpcserver

import (
	"github.com/Metchain/Metblock/utils/logger"
	"github.com/Metchain/Metblock/utils/panics"
)

var log = logger.RegisterSubSystem("TXMP")
var spawn = panics.GoroutineWrapperFunc(log)
