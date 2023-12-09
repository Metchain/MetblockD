package transactionhelper

import "github.com/Metchain/Metblock/mconfig/externalapi"

func NewSubnetworkTransaction(version uint16, inputs []*externalapi.DomainTransactionInput,
	outputs []*externalapi.DomainTransactionOutput, subnetworkID *externalapi.DomainSubnetworkID,
	gas uint64, payload []byte) *externalapi.DomainTransaction {

	return &externalapi.DomainTransaction{
		Version:      version,
		Inputs:       inputs,
		Outputs:      outputs,
		LockTime:     0,
		SubnetworkID: *subnetworkID,
		Gas:          gas,
		Payload:      payload,
		Fee:          0,
		Mass:         0,
	}
}
