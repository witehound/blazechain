package network

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/witehound/blazechain/core"
	"github.com/witehound/blazechain/util"
)

func TestTxMaxLength(t *testing.T) {
	p := NewMemePool(1)
	p.Add(util.NewRandomTransaction(strconv.Itoa(rand.Intn(1000000000))))
	assert.Equal(t, 1, p.all.Count())

	p.Add(util.NewRandomTransaction(strconv.Itoa(rand.Intn(1000000000))))
	p.Add(util.NewRandomTransaction(strconv.Itoa(rand.Intn(1000000000))))
	p.Add(util.NewRandomTransaction(strconv.Itoa(rand.Intn(1000000000))))
	tx := util.NewRandomTransaction(strconv.Itoa(rand.Intn(1000000000)))
	p.Add(tx)
	assert.Equal(t, 1, p.all.Count())
	assert.True(t, p.Contains(tx.Hash(core.TxHasher{})))
}

func TestTxPoolAdd(t *testing.T) {
	p := NewMemePool(11)
	n := 10

	for i := 1; i <= n; i++ {
		tx := util.NewRandomTransaction(strconv.Itoa(rand.Intn(1000000000)))
		p.Add(tx)
		// cannot add twice
		p.Add(tx)

		assert.Equal(t, i, p.PendingCount())
		assert.Equal(t, i, p.pending.Count())
		assert.Equal(t, i, p.all.Count())
	}
}

func TestTxPoolMaxLength(t *testing.T) {
	maxLen := 10
	p := NewMemePool(maxLen)
	n := 100
	txx := []*core.Transaction{}

	for i := 0; i < n; i++ {
		tx := util.NewRandomTransaction(strconv.Itoa(rand.Intn(1000000000)))
		p.Add(tx)

		if i > n-(maxLen+1) {
			txx = append(txx, tx)
		}
	}

	assert.Equal(t, p.all.Count(), maxLen)
	assert.Equal(t, len(txx), maxLen)

	for _, tx := range txx {
		assert.True(t, p.Contains(tx.Hash(core.TxHasher{})))
	}
}

func TestTxSortedMapFirst(t *testing.T) {
	m := NewTxSortedMap()
	first := util.NewRandomTransaction(strconv.Itoa(rand.Intn(1000000000)))
	m.Add(first)
	m.Add(util.NewRandomTransaction(strconv.Itoa(rand.Intn(1000000000))))
	m.Add(util.NewRandomTransaction(strconv.Itoa(rand.Intn(1000000000))))
	m.Add(util.NewRandomTransaction(strconv.Itoa(rand.Intn(1000000000))))
	m.Add(util.NewRandomTransaction(strconv.Itoa(rand.Intn(1000000000))))
	assert.Equal(t, first, m.First())
}

func TestTxSortedMapAdd(t *testing.T) {
	m := NewTxSortedMap()
	n := 100

	for i := 0; i < n; i++ {
		tx := util.NewRandomTransaction(strconv.Itoa(rand.Intn(1000000000)))
		m.Add(tx)
		// cannot add the same twice
		m.Add(tx)

		assert.Equal(t, m.Count(), i+1)
		assert.True(t, m.Contains(tx.Hash(core.TxHasher{})))
		assert.Equal(t, len(m.lookup), m.txx.Len())
		assert.Equal(t, m.Get(tx.Hash(core.TxHasher{})), tx)
	}

	m.Clear()
	assert.Equal(t, m.Count(), 0)
	assert.Equal(t, len(m.lookup), 0)
	assert.Equal(t, m.txx.Len(), 0)
}

func TestTxSortedMapRemove(t *testing.T) {
	m := NewTxSortedMap()

	tx := util.NewRandomTransaction(strconv.Itoa(rand.Intn(1000000000)))
	m.Add(tx)
	assert.Equal(t, m.Count(), 1)

	m.Remove(tx.Hash(core.TxHasher{}))
	assert.Equal(t, m.Count(), 0)
	assert.False(t, m.Contains(tx.Hash(core.TxHasher{})))
}
