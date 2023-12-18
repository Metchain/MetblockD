package dnsseed

import (
	"github.com/Metchain/Metblock/utils/logger"
	"github.com/Metchain/Metblock/utils/panics"
)

var log = logger.RegisterSubSystem("CMGR")
var spawn = panics.GoroutineWrapperFunc(log)
