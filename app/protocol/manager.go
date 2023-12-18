package protocol

import (
	"fmt"
	"github.com/Metchain/Metblock/app/protocol/common"
	"github.com/Metchain/Metblock/app/protocol/flowcontext"
	peerpkg "github.com/Metchain/Metblock/app/protocol/peer"
	"sync"
)

func (m *Manager) runFlows(flows []*common.Flow, peer *peerpkg.Peer, errChan <-chan error, flowsWaitGroup *sync.WaitGroup) error {
	flowsWaitGroup.Add(len(flows))
	for _, flow := range flows {
		executeFunc := flow.ExecuteFunc // extract to new variable so that it's not overwritten
		spawn(fmt.Sprintf("flow-%s", flow.Name), func() {
			executeFunc(peer)
			flowsWaitGroup.Done()
		})
	}

	return <-errChan
}

// SetOnNewBlockTemplateHandler sets the onNewBlockTemplate handler
func (m *Manager) SetOnNewBlockTemplateHandler(onNewBlockTemplateHandler flowcontext.OnNewBlockTemplateHandler) {
	m.context.SetOnNewBlockTemplateHandler(onNewBlockTemplateHandler)
}
