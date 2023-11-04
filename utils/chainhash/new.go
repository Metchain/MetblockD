package chainhash

import (
	"crypto/sha256"
	"fmt"
	"github.com/Metchain/Metblock/mconfig"
	"strings"
)

type Merkle struct {
	m string
	p string
	n string
	t string
}

func GenrateMerkle() *Merkle {
	M := new(Merkle)
	M.m = fmt.Sprintf("%x", Rectify(mconfig.Hexmsg))
	M.p = fmt.Sprintf("%x", mconfig.Prevhash)
	M.n = fmt.Sprintf("%x", mconfig.NonceHex)
	M.t = fmt.Sprintf("%x", mconfig.TimestampHex)
	return M
}

func (m *Merkle) GenrateMerkle() [32]byte {
	msg := []string{m.m, m.p, m.n, m.t}
	mh := sha256.Sum256([]byte(strings.Join(msg, "!!!!!")))

	return mh
}

func GenrateCoinBase() [64]byte {
	m := GenrateMerkle()
	mr := fmt.Sprintf("%x", m.GenrateMerkle())
	ms := fmt.Sprintf("%x", Rectify(mconfig.Hexmsg))
	p := fmt.Sprintf("%x", mconfig.Prevhash)
	n := fmt.Sprintf("%x", mconfig.NonceHex)
	t := fmt.Sprintf("%x", mconfig.TimestampHex)
	msg := []string{mr, ms, p, n, t}
	CoinBasehash := GenrateChainBig((strings.Join(msg, "*****")))
	return CoinBasehash
}
