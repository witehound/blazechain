package core

import (
	"io"

	"github.com/witehound/blazechain/crypto"
	"github.com/witehound/blazechain/types"
)

type Header struct {
	Version      uint32
	PrevlockHash types.Hash
	TimeStamp    int64
	Height       uint32
	DataHash     types.Hash
}

type Block struct {
	Header
	Transactions []Transaction
	hash         types.Hash
	Validator    crypto.PublicKey
	Signature    crypto.Signature
}

func (b *Block) Hash(hasher Hasher[*Block]) types.Hash {
	if b.hash.FindCachedHash() {
		b.hash = hasher.Hash(b)
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

func (b *Block) Sign() {

}
