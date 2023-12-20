package grpcclient

import (
	"github.com/Metchain/MetblockD/utils/logger"
	"github.com/Metchain/MetblockD/utils/panics"
)

var log = logger.RegisterSubSystem("MET-RPCC")
var spawn = panics.GoroutineWrapperFunc(log)
