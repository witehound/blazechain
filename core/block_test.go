package core

import (
	"fmt"
	"testing"
	"time"

	"github.com/witehound/blazechain/types"
)

func RandomBlock(height uint32) *Block {
	h := &Header{
		Version:      1,
		PrevlockHash: types.RandomHash(),
		Height:       height,
		TimeStamp:    time.Now().UnixNano(),
	}

	txs := Transaction{
		Data: []byte("foo"),
	}

	return NewBlock(Header(*h), []Transaction{txs})
}

func TestBlock_Hash(t *testing.T) {
	b := RandomBlock(0)
	fmt.Println(b.Hash(BlockHasher{}))
}
