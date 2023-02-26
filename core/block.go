package core

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"

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
	if b.hash.FindCachedHash() {
		b.hash = hasher.Hash(&b.Header)
	}

	return b.hash
}

func (b *Block) Decode(r io.Reader, dec Decoder[*Block]) error {
	return dec.Decode(r, b)
}

func (b *Block) Encode(w io.Writer, enc Encoder[*Block]) error {
	return enc.Encode(w, b)
}

func NewBlock(h Header, txs []Transaction) *Block {
	return &Block{
		Header:       h,
		Transactions: txs,
	}
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

	return nil
}

func (h *Header) Bytes() []byte {
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	enc.Encode(h)

	return buf.Bytes()
}
