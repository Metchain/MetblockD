package metwallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"math/big"
)

func (w *Wallet) PrivateKeyStr() string {
	//return as hexdecimal numeral
	return fmt.Sprintf("%x", w.privateKey.D.Bytes())
}
func (w *Wallet) PublicKeyStr() string {
	//return as hexdecimal numeral
	return fmt.Sprintf("%x%x", w.publicKey.X.Bytes(), w.publicKey.Y.Bytes())
}

func (w *Wallet) BlockchainAddress() string {
	return w.blockchainAddress
}

func (w *Wallet) Seedphrase() string {
	return w.seedphrase
}

func (w *Wallet) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		BlockchainAddress string `json:"blockchain_address"`
		Seedphrase        string `json:"seedphrase"`
	}{

		BlockchainAddress: w.BlockchainAddress(),
		Seedphrase:        w.Seedphrase(),
	})
}

func toECDSA(d []byte) (*ecdsa.PrivateKey, error) {
	priv := new(ecdsa.PrivateKey)
	priv.PublicKey.Curve = S256()
	if 8*len(d) != priv.Params().BitSize {
		return nil, fmt.Errorf("invalid length, need %d bits", priv.Params().BitSize)
	}
	priv.D = new(big.Int).SetBytes(d)

	// The priv.D must not be zero or negative.
	if priv.D.Sign() <= 0 {
		return nil, errors.New("invalid private key, zero or negative")
	}

	priv.PublicKey.X, priv.PublicKey.Y = priv.PublicKey.Curve.ScalarBaseMult(d)
	if priv.PublicKey.X == nil {
		return nil, errors.New("invalid private key")
	}
	return priv, nil
}

func S256() elliptic.Curve {
	return elliptic.P256()
}
