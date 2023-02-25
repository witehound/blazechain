package core

import "github.com/witehound/blazechain/types"

type Hasher[T any] interface {
	Hash(T) types.Hash
}
