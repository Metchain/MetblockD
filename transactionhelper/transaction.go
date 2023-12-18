package transactionhelper

import (
	"github.com/Metchain/Metblock/external"
)

func NewSubnetworkTransaction(version uint16, inputs []*external.DomainTransactionInput,
	subnetworkID *external.DomainSubnetworkID,
	gas uint64, payload []byte) *external.DomainTransaction {

	return &external.DomainTransaction{
		Version: version,
		Inputs:  inputs,

		LockTime: 0,

		Payload: payload,
		Fee:     0,
	}
}
