package core

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/witehound/blazechain/crypto"
	"github.com/witehound/blazechain/types"
)

func RandomBlock(height uint32) *Block {
	h := &Header{
		Version:       1,
		PrevBlockHash: types.RandomHash(),
		Height:        height,
		TimeStamp:     time.Now().UnixNano(),
	}

	tx := NewTransactionWithSig("food")

	return NewBlock(Header(*h), []Transaction{*tx})
}

func (bc *BlockChain) RandomBlockWithSig(t *testing.T, height uint32) (*Block, error) {

	b, err := bc.BlockWithHash(t, height)

	if err != nil {
		return nil, err
	}

	privkey := crypto.GeneratePrivateKey()
	tx := NewTransactionWithSig("food")
	b.AddTransaction(tx)

	fmt.Println(b.Transactions)

	dataHash, err := CalculateDataHash(b.Transactions)

	assert.Nil(t, err)

	b.Header.DataHash = dataHash

	assert.Nil(t, b.Sign(privkey))

	return b, nil
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

	dataHash, err := CalculateDataHash(b.Transactions)

	assert.Nil(t, err)
	b.Header.DataHash = dataHash

	assert.Nil(t, b.Sign(privkey))

	assert.Nil(t, b.Verify())

	v := RandomBlock(1)

	assert.NotNil(t, v.Verify())

}
