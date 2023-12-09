package protocol

import (
	"sync"
)

type Manager struct {
	routersWaitGroup sync.WaitGroup
	isClosed         uint32
}
