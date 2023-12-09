package addressmanager

import (
	"github.com/Metchain/Metblock/mconfig/infraconfig"
	"net"
)

// Config is a descriptor which specifies the AddressManager instance configuration.
type Config struct {
	AcceptUnroutable bool
	DefaultPort      string
	ExternalIPs      []string
	Listeners        []string
	Lookup           func(string) ([]net.IP, error)
}

// NewConfig returns a new address manager Config.
func NewConfig(cfg *infraconfig.Config) *Config {
	return &Config{
		DefaultPort: cfg.NetParams().DefaultPort,
		ExternalIPs: cfg.ExternalIPs,
		Listeners:   cfg.Listeners,
		Lookup:      cfg.Lookup,
	}
}
