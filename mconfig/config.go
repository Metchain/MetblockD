package mconfig

import (
	"fmt"
	"github.com/Metchain/Metblock/utils/appdir"
	"log"
	"net"
	"strconv"
)

const (
	AppName        = "Metchain"
	DbDir          = "data"
	DbDirTestnet   = "dataverify"
	MainnetPortRPC = 14031
	MainnetPortP2p = 14041
	RPCMaxClients  = 125
	MinimumStaking = 15000
	Lock3Month     = 131400
	Lock6Month     = 262800
	Lock9Month     = 350400
	Lock12Month    = 525600
	DeadWallet     = "metchain:000000000000000000000000000000000000000000000000000000000000DEAD"
)

type Config struct {
	RPCPORT       []string
	RPCListeners  []string
	RPCMaxClients int
	Listeners     []string
}

func GetCfg() *Config {
	cfg := &Config{}
	cfg.RPCPORT = []string{":" + strconv.Itoa(MainnetPortRPC)}
	cfg.RPCListeners = []string{":" + strconv.Itoa(MainnetPortRPC)}
	cfg.Listeners = []string{":" + strconv.Itoa(MainnetPortP2p)}
	cfg.RPCMaxClients = RPCMaxClients
	return cfg
}
func GetDatadir() string {
	d := appdir.AppDataDir(AppName, false)
	return d
}

func GetDBDir() string {
	d := appdir.AppDataDir(AppName, false) + "\\" + DbDir
	log.Println(d)
	return d
}

func GetDBDirTestnet() string {
	d := appdir.AppDataDir(AppName, false) + "\\" + DbDirTestnet
	log.Println(d)
	return d
}

//func GetRPCNodes() []string {
//	nodes := []string{}
//}

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil || ipnet.IP.To16 != nil {
				// print available addresses
				//fmt.Println(ipnet.IP.String())
			}
		}
	}
	for _, addr := range addrs {
		var ip net.IP
		switch v := addr.(type) {
		case *net.IPAddr:
			ip = v.IP
		case *net.IPNet:
			ip = v.IP
		default:
			continue
		}
		// print the available ip addresses
		log.Println(ip.String())
	}
	return ""
}

func GetLocalIPs() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddress := conn.LocalAddr().(*net.UDPAddr)

	return localAddress.IP
}

func VerifyDomainIp(domain string) string {

	ips, err := net.LookupIP(domain)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	for _, ip := range ips {
		return ip.String()
	}
	return ""
}
