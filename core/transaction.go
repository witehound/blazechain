package core

import (
	"crypto"

	"github.com/witehound/blazechain/types"
)

type Transaction struct {
	Data      []byte
	PublicKey crypto.PublicKey
	Signature *types.Signature
}

// func (tx *Transaction) EncodeBinary(w io.Writer) error {
// 	return nil
// }

// func (tx *Transaction) DecodeBinary(r io.Reader) error {
// 	return nil
// }
