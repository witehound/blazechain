package core

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/witehound/blazechain/crypto"
)

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
