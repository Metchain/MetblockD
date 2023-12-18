package external

type BlockLevelChildern []*DomainHash

// CloneParents creates a clone of the given BlockLevelParents slice
func CloneChildren(parents []BlockLevelChildern) []BlockLevelChildern {
	clone := make([]BlockLevelChildern, len(parents))
	for i, blockLevelChildern := range parents {
		clone[i] = blockLevelChildern.Clone()
	}
	return clone
}

// Clone creates a clone of this BlockLevelParents
func (sl BlockLevelChildern) Clone() BlockLevelChildern {
	return CloneHashes(sl)
}
