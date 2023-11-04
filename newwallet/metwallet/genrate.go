package metwallet

import (
	"fmt"
	"github.com/blockchainspectre/go-bip32"
	"github.com/btcsuite/btcutil/base58"
	"github.com/tyler-smith/go-bip39"
	"log"
)

func createWallet() *walletk {

	// Generate a mnemonic for memorization or user-friendly seeds
	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	nseed := fmt.Sprintf("%v", mnemonic)

	// Generate a Bip32 HD wallet for the mnemonic and a user supplied password
	seed := bip39.NewSeed(nseed, "Secret Passphrase")

	masterKey, _ := bip32.NewMasterKey(seed)

	// m/44'
	key, err := masterKey.NewChildKey(2147483648 + 44)
	if err != nil {
		log.Fatal(err)
	}

	decoded := base58.Decode(key.B58Serialize())
	privateKey := decoded[46:78]
	log.Println(privateKey)

	// Hex private key to ECDSA private key
	privateKeyECDSA, err := toECDSA(privateKey)
	if err != nil {
		log.Fatal(err)
	}
	walletk := new(walletk)
	//keys, _ := ecdsa.GenerateKey(privateKeyECDSA, rand.Reader)

	walletk.privateKey = privateKeyECDSA
	walletk.mnemonic = nseed

	return walletk
	// ECDSA private key to hex private key

}

func createWalletFormSeed(mnemonic string) walletk {

	seed := bip39.NewSeed(mnemonic, "Secret Passphrase")

	masterKey, _ := bip32.NewMasterKey(seed)

	// m/44'
	key, err := masterKey.NewChildKey(2147483648 + 44)
	if err != nil {
		log.Fatal(err)
	}

	decoded := base58.Decode(key.B58Serialize())
	privateKey := decoded[46:78]

	// Hex private key to ECDSA private key
	privateKeyECDSA, err := toECDSA(privateKey)
	if err != nil {
		log.Fatal(err)
	}
	walletk := walletk{
		privateKey: privateKeyECDSA,
		address:    fmt.Sprintf("%x%x", privateKeyECDSA.X.Bytes(), privateKeyECDSA.Y.Bytes()),
	}
	//keys, _ := ecdsa.GenerateKey(privateKeyECDSA, rand.Reader)

	return walletk
	// ECDSA private key to hex private key

}
