package core

import (
	"crypto/sha256"

	"github.com/witehound/blazechain/types"
)

type Hasher[T any] interface {
	Hash(T) types.Hash
}

type BlockHasher struct {
}

func (BlockHasher) Hash(b *Block) types.Hash {

	e := sha256.Sum256(b.HeaderData())

	return types.Hash(e)
}
