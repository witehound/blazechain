package core

import (
	"fmt"

	"github.com/witehound/blazechain/crypto"
)

type Transaction struct {
	Data      []byte
	From      crypto.PublicKey
	Signature *crypto.Signature
}

func (tx *Transaction) SignTx(key crypto.PrivateKey) error {
	sig, err := key.Sign(tx.Data)
	if err != nil {
		return err
	}
	tx.From = key.GetPublicKey()
	tx.Signature = sig

	return nil
}

func (tx *Transaction) VerifyTx() error {
	if tx.Signature == nil {
		return fmt.Errorf("transaction has no signature")
	}

	if !tx.Signature.Verify(tx.From, tx.Data) {
		return fmt.Errorf("invalid transaction")
	}

	return nil
}
