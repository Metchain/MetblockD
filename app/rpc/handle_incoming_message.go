package rpc

import "github.com/Metchain/Metblock/protoserver/routerpkg"

func (m *Manager) handleIncomingMessages(router *routerpkg.Router, incomingRoute *routerpkg.Route) error {
	outgoingRoute := router.OutgoingRoute()
	for {
		request, err := incomingRoute.Dequeue()
		if err != nil {
			return err
		}
		handler, ok := handlers[request.Command()]
		if !ok {
			return err
		}
		response, err := handler(m.context, router, request)
		if err != nil {
			return err
		}

		err = outgoingRoute.Enqueue(response)
		if err != nil {
			return err
		}
	}
}
