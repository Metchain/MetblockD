package commanager

import "github.com/Metchain/MetblockD/appmessage"

// checkOutgoingConnections goes over all activeOutgoing and makes sure they are still active.
// Then it opens connections so that we have targetOutgoing active connections
func (c *ConnectionManager) checkOutgoingConnections(connSet connectionSet) {
	for address := range c.activeOutgoing {
		connection, ok := connSet.get(address)
		if ok { // connection is still connected
			connSet.remove(connection)
			continue
		}

		// if connection is dead - remove from list of active ones
		delete(c.activeOutgoing, address)
	}

	connections := c.netAdapter.P2PConnections()
	connectedAddresses := make([]*appmessage.NetAddress, len(connections))
	for i, connection := range connections {
		connectedAddresses[i] = connection.NetAddress()
	}

	liveConnections := len(c.activeOutgoing)
	if c.targetOutgoing == liveConnections {
		return
	}

	log.Debugf("Have got %d outgoing connections out of target %d, adding %d more",
		liveConnections, c.targetOutgoing, c.targetOutgoing-liveConnections)

}
