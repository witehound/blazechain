package core

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/witehound/blazechain/crypto"
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

func (bc *BlockChain) RandomBlockWithSig(height uint32) *Block {
	b := RandomBlock(height)
	privkey := crypto.GeneratePrivateKey()
	b.Sign(privkey)

	return b
}

func TestBlock_Signing(t *testing.T) {
	privkey := crypto.GeneratePrivateKey()
	b := RandomBlock(0)

	assert.Nil(t, b.Sign(privkey))

	assert.NotNil(t, b.Signature)
}

func TestBlock_Verifying(t *testing.T) {
	privkey := crypto.GeneratePrivateKey()
	b := RandomBlock(0)

	assert.Nil(t, b.Sign(privkey))

	assert.Nil(t, b.Verify())

	v := RandomBlock(1)

	assert.NotNil(t, v.Verify())

}
