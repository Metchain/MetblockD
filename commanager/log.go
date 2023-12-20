package commanager

import (
	"github.com/Metchain/MetblockD/utils/logger"
	"github.com/Metchain/MetblockD/utils/panics"
)

var log = logger.RegisterSubSystem("MET-CMGR")
var spawn = panics.GoroutineWrapperFunc(log)
