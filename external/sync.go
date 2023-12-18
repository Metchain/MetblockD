package external

// SyncInfo holds info about the current sync state of the consensus
type SyncInfo struct {
	HeaderCount uint64
	BlockCount  uint64
}
