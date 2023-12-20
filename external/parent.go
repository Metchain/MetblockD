package external

type BlockLevelParents []*DomainHash

// CloneParents creates a clone of the given BlockLevelParents slice
func CloneParents(parents []BlockLevelParents) []BlockLevelParents {
	clone := make([]BlockLevelParents, len(parents))
	for i, blockLevelParents := range parents {
		clone[i] = blockLevelParents.Clone()
	}
	return clone
}

// Clone creates a clone of this BlockLevelParents
func (sl BlockLevelParents) Clone() BlockLevelParents {
	return CloneHashes(sl)
}
