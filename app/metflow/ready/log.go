package ready

import (
	"github.com/Metchain/MetblockD/utils/logger"
	"github.com/Metchain/MetblockD/utils/panics"
)

var log = logger.RegisterSubSystem("METD-Ready")
var spawn = panics.GoroutineWrapperFunc(log)
