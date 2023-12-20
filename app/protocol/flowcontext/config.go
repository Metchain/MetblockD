package flowcontext

import "github.com/Metchain/MetblockD/mconfig/infraconfig"

func (f *FlowContext) Config() *infraconfig.Config {
	return f.cfg
}
