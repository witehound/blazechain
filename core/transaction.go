package core

import (
	"fmt"

	"github.com/witehound/blazechain/crypto"
	"github.com/witehound/blazechain/types"
)

type Transaction struct {
	Data      []byte
	From      crypto.PublicKey
	Signature *crypto.Signature
	hash      types.Hash
	firstSeen int64
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

func (tx *Transaction) Hash(hasher Hasher[*Transaction]) types.Hash {
	if tx.hash.IsZero() {
		tx.hash = hasher.Hash(tx)
	}
	return tx.hash
}

func NewTransactionWithSig(data string) *Transaction {
	privkey := crypto.GeneratePrivateKey()
	tx := &Transaction{
		Data: []byte(data),
	}

	tx.SignTx(privkey)

	return tx

}

func (tx *Transaction) Decode(dec Decoder[*Transaction]) error {
	return dec.Decode(tx)
}

func (tx *Transaction) Encode(enc Encoder[*Transaction]) error {
	return enc.Encode(tx)
}

func (tx *Transaction) SetFirstSeen(t int64) {
	tx.firstSeen = t
}

func (tx *Transaction) FirstSeen() int64 {
	return tx.firstSeen
}
