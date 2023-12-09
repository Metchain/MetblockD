package profiling

import (
	"fmt"
	"time"
)

var heapDumpFileName = fmt.Sprintf("heap-%s.pprof", time.Now().Format("01-02-2006T15.04.05"))
