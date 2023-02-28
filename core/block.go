package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"

	"github.com/witehound/blazechain/crypto"
	"github.com/witehound/blazechain/types"
)

type Header struct {
	Version       uint32
	PrevBlockHash types.Hash
	TimeStamp     int64
	Height        uint32
	DataHash      types.Hash
}

type Block struct {
	Header       Header
	Transactions []Transaction
	hash         types.Hash
	Validator    crypto.PublicKey
	Signature    *crypto.Signature
}

func (b *Block) Hash(hasher Hasher[*Header]) types.Hash {
	if b.hash.IsZero() {
		b.hash = hasher.Hash(&b.Header)
	}

	return b.hash
}

func (b *Block) Decode(dec Decoder[*Block]) error {
	return dec.Decode(b)
}

func (b *Block) Encode(enc Encoder[*Block]) error {
	return enc.Encode(b)
}

func NewBlock(h Header, txs []Transaction) *Block {
	return &Block{
		Header:       h,
		Transactions: txs,
	}
}

func (b *Block) AddTransaction(tx *Transaction) {
	b.Transactions = append(b.Transactions, *tx)
}

func (b *Block) Sign(key crypto.PrivateKey) error {
	sig, err := key.Sign(b.Header.Bytes())

	if err != nil {
		return err
	}

	b.Validator = key.GetPublicKey()
	b.Signature = sig

	return nil
}

func (b *Block) Verify() error {
	if b.Signature == nil {
		return fmt.Errorf("block has no signature")
	}

	if !b.Signature.Verify(b.Validator, b.Header.Bytes()) {
		return fmt.Errorf("block has invalid signature")
	}

	for _, tx := range b.Transactions {
		if err := tx.VerifyTx(); err != nil {
			return err
		}
	}

	dataHash, err := CalculateDataHash(b.Transactions)

	if err != nil {
		return err
	}

	if dataHash != b.Header.DataHash {
		return fmt.Errorf("block has invalid transaction datahash")
	}

	return nil
}

func (h *Header) Bytes() []byte {
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	enc.Encode(h)

	return buf.Bytes()
}

func CalculateDataHash(tsx []Transaction) (hash types.Hash, err error) {
	var (
		buf = &bytes.Buffer{}
	)

	for _, tx := range tsx {
		if err = tx.Encode(NewGobTxEncoder(buf)); err != nil {
			return
		}
	}

	hash = sha256.Sum256(buf.Bytes())
	return
}
