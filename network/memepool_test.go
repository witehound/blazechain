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

	assert.Nil(t, mp.AddTx(tx))

	assert.Equal(t, mp.Len(), 1)
}
