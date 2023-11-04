package nodes

import (
	"net"
	"os"
)

func GetHost() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "127.0.0.1"
	}
	address, err := net.LookupHost(hostname)

	if err != nil {
		return "127.0.0.1"
	}
	//fmt.Printf(address[3])
	// address[0] will return your IP address
	return address[0]
}
