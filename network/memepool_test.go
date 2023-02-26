package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/witehound/blazechain/core"
)

func TestNewMemepool(t *testing.T) {
	mp := NewMemePool()
	assert.Equal(t, mp.Len(), 0)

	tx := core.NewTransactionWithSig("first")

	hash := tx.Hash(core.TxHasher{})

	assert.Nil(t, mp.AddTx(hash, tx))

	assert.Equal(t, mp.Len(), 1)

	mp.Flush()
	assert.Equal(t, mp.Len(), 0)
}
