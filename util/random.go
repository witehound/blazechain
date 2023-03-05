package util

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/witehound/blazechain/core"
	"github.com/witehound/blazechain/crypto"
	"github.com/witehound/blazechain/types"
)

func RandomBytes(size int) []byte {
	token := make([]byte, size)
	rand.Read(token)
	return token
}

func RandomHash() types.Hash {
	return types.HashFromBytes(RandomBytes(32))
}

// NewRandomTransaction return a new random transaction whithout signature.
func NewRandomTransaction(size string) *core.Transaction {
	return core.NewTransactionWithSig(size)
}

func NewRandomTransactionWithSignature(t *testing.T, privKey crypto.PrivateKey, size string) *core.Transaction {
	tx := NewRandomTransaction(size)
	assert.Nil(t, tx.SignTx(privKey))
	return tx
}

func NewRandomBlock(t *testing.T, height uint32, prevBlockHash types.Hash) *core.Block {
	txSigner := crypto.GeneratePrivateKey()
	tx := NewRandomTransactionWithSignature(t, txSigner, "foo")
	header := &core.Header{
		Version:       1,
		PrevBlockHash: prevBlockHash,
		Height:        height,
		TimeStamp:     time.Now().UnixNano(),
	}
	b, err := core.NewBlock(*header, []*core.Transaction{tx})
	assert.Nil(t, err)
	dataHash, err := core.CalculateDataHash(b.Transactions)
	assert.Nil(t, err)
	b.Header.DataHash = dataHash

	return b
}

func NewRandomBlockWithSignature(t *testing.T, pk crypto.PrivateKey, height uint32, prevHash types.Hash) *core.Block {
	b := NewRandomBlock(t, height, prevHash)
	assert.Nil(t, b.Sign(pk))

	return b
}
