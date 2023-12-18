package commanager

import (
	"github.com/Metchain/Metblock/protoserver"
)

type connectionSet map[string]*protoserver.NetConnection

func (cs connectionSet) add(connection *protoserver.NetConnection) {
	cs[connection.Address()] = connection
}

func (cs connectionSet) remove(connection *protoserver.NetConnection) {
	delete(cs, connection.Address())
}

func (cs connectionSet) get(address string) (*protoserver.NetConnection, bool) {
	connection, ok := cs[address]
	return connection, ok
}

func convertToSet(connections []*protoserver.NetConnection) connectionSet {
	connSet := make(connectionSet, len(connections))

	for _, connection := range connections {
		connSet[connection.Address()] = connection
	}

	return connSet
}
