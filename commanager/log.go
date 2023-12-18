package commanager

import (
	"github.com/Metchain/Metblock/utils/logger"
	"github.com/Metchain/Metblock/utils/panics"
)

var log = logger.RegisterSubSystem("MET-CMGR")
var spawn = panics.GoroutineWrapperFunc(log)
