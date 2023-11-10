package consensus

import (
	"fmt"
	"github.com/Metchain/Metblock/mconfig"
	"github.com/Metchain/Metblock/utils/chainhash"
	"os"
	"strings"
)

func VerfiyNonce(i int) (string, bool) {
	os.Stdout = nil
	mn := fmt.Sprintf("%x", mconfig.NonceHex)
	n := fmt.Sprintf("%x", chainhash.GenrateChainInt(i))
	if mn == n {
		return "Successfully verified Nonce", false
	}
	return "Error Verifying Nonce", true
}

func VerifyTimestamp(i int) (string, bool) {
	mn := fmt.Sprintf("%x", mconfig.TimestampHex)
	n := fmt.Sprintf("%x", chainhash.GenrateChainInt(i))
	if mn == n {
		return "Successfully verified TimeStamp", false
	}
	return "Error Verifying TimeStamp", true
}

func VerifyMessage() (string, bool) {
	mn := fmt.Sprintf("%x", chainhash.Rectify(mconfig.Hexmsg))
	n := fmt.Sprintf("%x", chainhash.GenrateChainString("This is just the beginning of the blockchain. Just getting started."))
	if mn == n {
		return "Successfully verified Message", false
	}
	return "Error Verifying Message", true
}

func VerifyMerkleRoot() (string, bool) {
	m := chainhash.GenrateMerkle()
	mn := fmt.Sprintf("%x", mconfig.MerkleHex)
	mr := fmt.Sprintf("%x", m.GenrateMerkle())
	if mn == mr {
		return "Successfully Verified MerkleRoot", false
	}
	return "Error Verifing MerkleRoot", true
}

func VerifyPreviousHash() (string, bool) {
	mn := fmt.Sprintf("%x", mconfig.Prevhash)
	n := strings.Repeat("0", 64)

	if mn == n {
		return "Successfully verified PreviousHash", false
	}
	return "Error Verifying PreviousHash", true
}

func VerifyTransaction() (string, bool) {
	mn := fmt.Sprintf("%x", chainhash.Rectify(mconfig.CoinBaseHex))
	n := fmt.Sprintf("%x", chainhash.GenrateCoinBase())

	if mn == n {
		return "Successfully verified Coinbase TX", false
	}
	return "Error Verifying Coinbase TX", true
}
