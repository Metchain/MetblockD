package protoserver

func (c *NetConnection) setOnDisconnectedHandler(onDisconnectedHandler OnDisconnectedHandler) {
	c.onDisconnectedHandler = onDisconnectedHandler
}
