package flowcontext

import (
	"github.com/Metchain/MetblockD/external"
	"sync"
)

type SharedRequestedBlocks struct {
	blocks map[external.DomainHash]struct{}
	sync.Mutex
}
