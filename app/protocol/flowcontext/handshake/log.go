package handshake

import (
	"github.com/Metchain/Metblock/utils/logger"
	"github.com/Metchain/Metblock/utils/panics"
)

var log = logger.RegisterSubSystem("METD-Handshake")
var spawn = panics.GoroutineWrapperFunc(log)
