package protoserver

import (
	"github.com/Metchain/Metblock/protoserver/routerpkg"
	"sync/atomic"
)

func newNetConnection(connection Connection, routerInitializer RouterInitializer, name string) *NetConnection {
	router := routerpkg.NewRouter(name)

	netConnection := &NetConnection{
		connection: connection,
		router:     router,
	}

	netConnection.connection.SetOnDisconnectedHandler(func() {
		log.Infof("Disconnected from %s", netConnection)
		// If the disconnection came because of a network error and not because of the application layer, we
		// need to close the router as well.
		if atomic.AddUint32(&netConnection.isRouterClosed, 1) == 1 {
			netConnection.router.Close()
		}
		netConnection.onDisconnectedHandler()
	})

	routerInitializer(router, netConnection)

	return netConnection
}
