package core

import (
	"github.com/witehound/blazechain/crypto"
)

type Transaction struct {
	Data      []byte
	PublicKey crypto.PublicKey
	Signature *crypto.Signature
}

func (tx *Transaction) SignTx(key crypto.PrivateKey) error {

	return nil
}
