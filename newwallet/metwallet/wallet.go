package metwallet

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/Metchain/Metblock/heavyhash"
	"time"
)

type Wallet struct {
	privateKey        *ecdsa.PrivateKey
	publicKey         *ecdsa.PublicKey
	blockchainAddress string
	seedphrase        string
}

func NewWallet() *Wallet {
	// 1. Creating ECDSA private key (32 bytes) public key (64 bytes)
	nw := createWallet()
	w := new(Wallet)
	w.privateKey = nw.privateKey
	w.publicKey = &w.privateKey.PublicKey
	w.seedphrase = nw.mnemonic

	// 2. Perform SHA-256 hashing on the public key (32 bytes).
	address := WalletSigs(w.seedphrase)

	w.blockchainAddress = address
	return w
}

// For webwallet UI

type walletk struct {
	mnemonic   string
	privateKey *ecdsa.PrivateKey
	address    string
}

func WalletSigs(mnemonic string) string {
	address := [13]walletk{}
	seed := [12]string{}
	seed[1] = CreateFSeed(mnemonic)
	seed[2] = CreateSSeed(mnemonic)
	seed[3] = CreateTSeed(mnemonic)
	seed[4] = CreateFoSeed(mnemonic)
	seed[5] = CreateFiSeed(mnemonic)
	seed[6] = CreateSiSeed(mnemonic)
	seed[7] = CreateSeSeed(mnemonic)
	seed[8] = CreateEiSeed(mnemonic)
	seed[9] = CreateNSeed(mnemonic)
	seed[10] = CreateTeSeed(mnemonic)
	seed[11] = CreateElSeed(mnemonic)
	address[0] = createWalletFormSeed(mnemonic)
	address[1] = createWalletFormSeed(seed[1])
	address[2] = createWalletFormSeed(seed[2])
	address[3] = createWalletFormSeed(seed[3])
	address[4] = createWalletFormSeed(seed[4])
	address[5] = createWalletFormSeed(seed[5])
	address[6] = createWalletFormSeed(seed[6])
	address[7] = createWalletFormSeed(seed[7])
	address[8] = createWalletFormSeed(seed[8])
	address[9] = createWalletFormSeed(seed[9])
	address[10] = createWalletFormSeed(seed[10])
	address[11] = createWalletFormSeed(seed[11])
	addr := ""

	for _, val := range address {
		if val.address != "" {
			addr = addr + ":met:" + val.address
		}

	}
	i := time.Unix(16400000005022, 0)
	walletaddress := heavyhash.HeavyHash([]byte(addr), i.UnixMilli(), 10)

	return fmt.Sprintf("metchain:%x", walletaddress)
}
