package metwallet

func WalletRecover(seed string) *Wallet {
	// 1. Creating ECDSA private key (32 bytes) public key (64 bytes)

	// 2. Perform SHA-256 hashing on the public key (32 bytes).
	address := WalletSigs(seed)
	w := new(Wallet)
	w.seedphrase = seed
	w.blockchainAddress = address
	return w
}
