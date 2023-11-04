package chainhash

import (
	"crypto/sha256"
	"crypto/sha512"
	"strconv"
)

func GenrateChainInt(input int) [32]byte {
	n := strconv.Itoa(input)
	nhex := sha256.Sum256([]byte(n))

	return nhex
}

func GenrateChainString(input string) [32]byte {
	nhex := sha256.Sum256([]byte(input))

	return nhex
}

func GenrateChainBig(input string) [64]byte {
	nhex := sha512.Sum512([]byte(input))
	return nhex
}

func Rectify(i []byte) []byte {
	n := []byte{}
	p := false
	for _, j := range i {
		if j != 213 {
			n = append(n, j)
			p = true
		}
	}
	if !p {
		return i
	}
	return n
}
