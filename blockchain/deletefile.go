package blockchain

func (bc *Blockchain) CreateTempBlock(i int64, ts int64) []byte {
	new := new(MiniBlock)
	new.height = uint64(i)
	new.timestamp = ts
	s, _ := new.MarshalJSON()
	return s

}
