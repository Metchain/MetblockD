package mempool

import (
	"github.com/Metchain/MetblockD/mconfig/dagconfig"
	"github.com/Metchain/MetblockD/mconfig/util"
	"time"
)

// Config represents a mempool configuration
type Config struct {
	MaximumTransactionCount               uint64
	TransactionExpireIntervalDAAScore     uint64
	TransactionExpireScanIntervalDAAScore uint64
	TransactionExpireScanIntervalSeconds  uint64
	OrphanExpireIntervalDAAScore          uint64
	OrphanExpireScanIntervalDAAScore      uint64
	MaximumOrphanTransactionMass          uint64
	MaximumOrphanTransactionCount         uint64
	AcceptNonStandard                     bool
	MaximumMassPerBlock                   uint64
	MinimumRelayTransactionFee            util.Amount
	MinimumStandardTransactionVersion     uint16
	MaximumStandardTransactionVersion     uint16
}

// DefaultConfig returns the default mempool configuration
func DefaultConfig(dagParams *dagconfig.Params) *Config {
	targetBlocksPerSecond := time.Second.Seconds() / dagParams.TargetTimePerBlock.Seconds()

	return &Config{
		MaximumTransactionCount:               defaultMaximumTransactionCount,
		TransactionExpireIntervalDAAScore:     uint64(float64(defaultTransactionExpireIntervalSeconds) / targetBlocksPerSecond),
		TransactionExpireScanIntervalDAAScore: uint64(float64(defaultTransactionExpireScanIntervalSeconds) / targetBlocksPerSecond),
		TransactionExpireScanIntervalSeconds:  defaultTransactionExpireScanIntervalSeconds,
		OrphanExpireIntervalDAAScore:          uint64(float64(defaultOrphanExpireIntervalSeconds) / targetBlocksPerSecond),
		OrphanExpireScanIntervalDAAScore:      uint64(float64(defaultOrphanExpireScanIntervalSeconds) / targetBlocksPerSecond),
		MaximumOrphanTransactionMass:          defaultMaximumOrphanTransactionMass,
		MaximumOrphanTransactionCount:         defaultMaximumOrphanTransactionCount,
		MinimumRelayTransactionFee:            defaultMinimumRelayTransactionFee,
		MinimumStandardTransactionVersion:     defaultMinimumStandardTransactionVersion,
		MaximumStandardTransactionVersion:     defaultMaximumStandardTransactionVersion,
	}
}
