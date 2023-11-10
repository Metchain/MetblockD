package blockchain

/*func (bc *Blockchain) ResolveConflicts() bool {
	var longestChain []*MiniBlock = nil
	maxLegth := len(bc.chain)

	for _, n := range bc.neighbors {
		endpoint := fmt.Sprintf("http://%s/chain", n)
		resp, _ := http.Get(endpoint)
		if resp.StatusCode == 200 {
			var bcResp Blockchain
			decoder := json.NewDecoder(resp.Body)
			_ = decoder.Decode(&bcResp)

			chain := bcResp.chain

			if len(chain) > maxLegth && bc.ValidChain(chain) {
				maxLegth = len(chain)
				longestChain = chain
			}
		}
	}

	if longestChain != nil {
		bc.chain = longestChain
		log.Print("Resolve conflicts replaced")
		return true
	}
	log.Printf("Resolve conflicts not replaced")
	return true
}
*/
