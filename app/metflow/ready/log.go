package ready

import (
	"github.com/Metchain/Metblock/utils/logger"
	"github.com/Metchain/Metblock/utils/panics"
)

var log = logger.RegisterSubSystem("METD-Ready")
var spawn = panics.GoroutineWrapperFunc(log)