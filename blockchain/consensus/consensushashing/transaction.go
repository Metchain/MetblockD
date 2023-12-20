package consensushashing

import (
	"github.com/Metchain/MetblockD/external"
	"github.com/Metchain/MetblockD/utils/binaryserializer"
	"github.com/Metchain/MetblockD/utils/hashes"
	"github.com/Metchain/MetblockD/utils/serialization"
	"io"

	"github.com/pkg/errors"
)

// txEncoding is a bitmask defining which transaction fields we
// want to encode and which to ignore.
type txEncoding uint8

const (
	txEncodingFull txEncoding = 0

	txEncodingExcludeSignatureScript = 1 << iota
)

// TransactionHash returns the transaction hash.
func TransactionHash(tx *external.DomainTransaction) *external.DomainHash {
	// Encode the header and hash everything prior to the number of
	// transactions.
	writer := hashes.NewTransactionHashWriter()
	err := serializeTransaction(writer, tx, txEncodingFull)
	if err != nil {
		// It seems like this could only happen if the writer returned an error.
		// and this writer should never return an error (no allocations or possible failures)
		// the only non-writer error path here is unknown types in `WriteElement`
		panic(errors.Wrap(err, "TransactionHash() failed. this should never fail for structurally-valid transactions"))
	}

	return writer.Finalize()
}

// TransactionID generates the Hash for the transaction without the signature script and payload field.
func TransactionID(tx *external.DomainTransaction) *external.DomainTransactionID {
	// If transaction ID is already cached, return it
	if tx.ID != nil {
		return tx.ID
	}

	// Encode the transaction, replace signature script with zeroes, cut off
	// payload and hash the result.
	var encodingFlags txEncoding
	// Metchainupdate needed
	/*if !transactionhelper.IsCoinBase(tx) {
		encodingFlags = txEncodingExcludeSignatureScript
	}*/
	writer := hashes.NewTransactionIDWriter()
	err := serializeTransaction(writer, tx, encodingFlags)
	if err != nil {
		// this writer never return errors (no allocations or possible failures) so errors can only come from validity checks,
		// and we assume we never construct malformed transactions.
		panic(errors.Wrap(err, "TransactionID() failed. this should never fail for structurally-valid transactions"))
	}
	transactionID := external.DomainTransactionID(*writer.Finalize())

	tx.ID = &transactionID

	return tx.ID
}

// TransactionIDs converts the provided slice of DomainTransactions to a corresponding slice of TransactionIDs
func TransactionIDs(txs []*external.DomainTransaction) []*external.DomainTransactionID {
	txIDs := make([]*external.DomainTransactionID, len(txs))
	for i, tx := range txs {
		txIDs[i] = TransactionID(tx)
	}
	return txIDs
}

func serializeTransaction(w io.Writer, tx *external.DomainTransaction, encodingFlags txEncoding) error {
	err := binaryserializer.PutUint16(w, tx.Version)
	if err != nil {
		return err
	}

	count := uint64(len(tx.Inputs))
	err = serialization.WriteElement(w, count)
	if err != nil {
		return err
	}

	for _, ti := range tx.Inputs {
		err = writeTransactionInput(w, ti, encodingFlags)
		if err != nil {
			return err
		}
	}

	err = binaryserializer.PutUint64(w, tx.LockTime)
	if err != nil {
		return err
	}

	err = writeVarBytes(w, tx.Payload)
	if err != nil {
		return err
	}

	return nil
}

// writeTransactionInput encodes ti to the Metchain protocol encoding for a transaction
// input to w.
func writeTransactionInput(w io.Writer, ti *external.DomainTransactionInput, encodingFlags txEncoding) error {
	err := writeOutpoint(w, &ti.PreviousOutpoint)
	if err != nil {
		return err
	}

	if encodingFlags&txEncodingExcludeSignatureScript != txEncodingExcludeSignatureScript {
		err = writeVarBytes(w, ti.SignatureScript)
		if err != nil {
			return err
		}

		_, err = w.Write([]byte{ti.SigOpCount})
		if err != nil {
			return err
		}
	} else {
		err = writeVarBytes(w, []byte{})
		if err != nil {
			return err
		}
	}

	return binaryserializer.PutUint64(w, ti.Sequence)
}

func writeOutpoint(w io.Writer, outpoint *external.DomainOutpoint) error {
	_, err := w.Write(outpoint.TransactionID.ByteSlice())
	if err != nil {
		return err
	}

	return binaryserializer.PutUint32(w, outpoint.Index)
}

func writeVarBytes(w io.Writer, data []byte) error {
	dataLength := uint64(len(data))
	err := serialization.WriteElement(w, dataLength)
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	return err
}
