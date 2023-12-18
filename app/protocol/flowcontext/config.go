package flowcontext

import "github.com/Metchain/Metblock/mconfig/infraconfig"

func (f *FlowContext) Config() *infraconfig.Config {
	return f.cfg
}
