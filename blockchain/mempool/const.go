package mempool

import (
	"github.com/Metchain/MetblockD/mconfig/constants"
	"github.com/Metchain/MetblockD/mconfig/util"
)

const (
	defaultMaximumTransactionCount = 1_000_000

	defaultTransactionExpireIntervalSeconds     uint64 = 60
	defaultTransactionExpireScanIntervalSeconds uint64 = 10
	defaultOrphanExpireIntervalSeconds          uint64 = 60
	defaultOrphanExpireScanIntervalSeconds      uint64 = 10

	defaultMaximumOrphanTransactionMass = 100000
	// defaultMaximumOrphanTransactionCount should remain small as long as we have recursion in
	// removeOrphans when removeRedeemers = true
	defaultMaximumOrphanTransactionCount = 50

	// defaultMinimumRelayTransactionFee specifies the minimum transaction fee for a transaction to be accepted to
	// the mempool and relayed. It is specified in sompi per 1kg (or 1000 grams) of transaction mass.
	defaultMinimumRelayTransactionFee = util.Amount(1000)

	// Standard transaction version range might be different from what consensus accepts, therefore
	// we define separate values in mempool.
	// However, currently there's exactly one transaction version, so mempool accepts the same version
	// as consensus.
	defaultMinimumStandardTransactionVersion = constants.MaxTransactionVersion
	defaultMaximumStandardTransactionVersion = constants.MaxTransactionVersion
)
