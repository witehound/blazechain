package core

import (
	"fmt"

	"github.com/witehound/blazechain/crypto"
)

type Transaction struct {
	Data      []byte
	PublicKey crypto.PublicKey
	Signature *crypto.Signature
}

func (tx *Transaction) SignTx(key crypto.PrivateKey) error {
	sig, err := key.Sign(tx.Data)
	if err != nil {
		return err
	}
	tx.PublicKey = key.GetPublicKey()
	tx.Signature = sig

	return nil
}

func (tx *Transaction) VerifyTx() error {
	if tx.Signature == nil {
		return fmt.Errorf("transaction has no signature")
	}

	if !tx.Signature.Verify(tx.PublicKey, tx.Data) {
		return fmt.Errorf("invalid transaction")
	}

	return nil
}
