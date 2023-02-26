package network

import (
	"math/rand"
	"strconv"
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

func TestSortMemePool(t *testing.T) {
	mp := NewMemePool()

	txLen := 100

	for i := 0; i < txLen; i++ {
		tx := core.NewTransactionWithSig(strconv.Itoa(i))
		hash := tx.Hash(core.TxHasher{})
		tx.SetFirstSeen(int64(i * rand.Intn(10000)))
		assert.Nil(t, mp.AddTx(hash, tx))

	}

	assert.Equal(t, txLen, mp.Len())

	tsx := mp.AllTransactions()

	for i := 0; i < len(tsx)-1; i++ {
		assert.True(t, tsx[i].FirstSeen() < tsx[i+1].FirstSeen())
	}
}
