package protocol

import (
	"github.com/Metchain/Metblock/utils/logger"
	"github.com/Metchain/Metblock/utils/panics"
)

var log = logger.RegisterSubSystem("METD-Portocol")
var spawn = panics.GoroutineWrapperFunc(log)
