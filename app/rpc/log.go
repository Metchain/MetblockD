package rpc

import (
	"github.com/Metchain/Metblock/utils/logger"
	"github.com/Metchain/Metblock/utils/panics"
)

var log = logger.RegisterSubSystem("METD-RPCS")
var spawn = panics.GoroutineWrapperFunc(log)
