package flowcontext

import (
	"github.com/Metchain/Metblock/external"
	"sync"
)

type SharedRequestedBlocks struct {
	blocks map[external.DomainHash]struct{}
	sync.Mutex
}
