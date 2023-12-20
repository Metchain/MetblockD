package serialization

import (
	"github.com/Metchain/MetblockD/external"
	"github.com/Metchain/MetblockD/utils/binaryserializer"
	"github.com/pkg/errors"
	"io"
)

func WriteElement(w io.Writer, element interface{}) error {
	// Attempt to write the element based on the concrete type via fast
	// type assertions first.
	switch e := element.(type) {
	case []byte:
		err := WriteElement(w, uint64(len(e)))
		if err != nil {
			return err
		}
		_, err = w.Write(e)
		if err != nil {
			return err
		}
		return nil
	case int16:
		err := binaryserializer.PutUint16(w, uint16(e))
		if err != nil {
			return err
		}
		return nil
	case uint16:
		err := binaryserializer.PutUint16(w, e)
		if err != nil {
			return err
		}
		return nil
	case int32:
		err := binaryserializer.PutUint32(w, uint32(e))
		if err != nil {
			return err
		}
		return nil

	case uint32:
		err := binaryserializer.PutUint32(w, e)
		if err != nil {
			return err
		}
		return nil

	case int64:
		err := binaryserializer.PutUint64(w, uint64(e))
		if err != nil {
			return err
		}
		return nil

	case uint64:
		err := binaryserializer.PutUint64(w, e)
		if err != nil {
			return err
		}
		return nil

	case uint8:
		err := binaryserializer.PutUint8(w, e)
		if err != nil {
			return err
		}
		return nil

	case bool:
		var err error
		if e {
			err = binaryserializer.PutUint8(w, 0x01)
		} else {
			err = binaryserializer.PutUint8(w, 0x00)
		}
		if err != nil {
			return err
		}
		return nil

	case external.DomainHash:
		_, err := w.Write(e.ByteSlice())
		if err != nil {
			return err
		}
		return nil

	case *external.DomainHash:
		_, err := w.Write(e.ByteSlice())
		if err != nil {
			return err
		}
		return nil

	case external.DomainTransactionID:
		_, err := w.Write(e.ByteSlice())
		if err != nil {
			return err
		}
		return nil

	case external.DomainSubnetworkID:
		_, err := w.Write(e[:])
		if err != nil {
			return err
		}
		return nil

	case *external.DomainSubnetworkID:
		_, err := w.Write(e[:])
		if err != nil {
			return err
		}
		return nil
	}

	return errors.Wrapf(errNoEncodingForType, "couldn't find a way to write type %T", element)
}

func WriteElements(w io.Writer, elements ...interface{}) error {
	for _, element := range elements {
		err := WriteElement(w, element)
		if err != nil {
			return err
		}
	}
	return nil
}
