package core

import (
	"bytes"
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
	b, err := RandomBlock(0)
	assert.Nil(t, err)

	assert.Nil(t, b.Sign(privkey))

	assert.NotNil(t, b.Signature)
}

func TestBlock_Verifying(t *testing.T) {
	privkey := crypto.GeneratePrivateKey()
	b, err := RandomBlock(0)
	assert.Nil(t, err)

	dataHash, err := CalculateDataHash(b.Transactions)

	assert.Nil(t, err)
	b.Header.DataHash = dataHash

	assert.Nil(t, b.Sign(privkey))

	assert.Nil(t, b.Verify())

	v, err := RandomBlock(1)

	assert.Nil(t, err)

	assert.NotNil(t, v.Verify())

}

func TestBlock_Decode_Encode(t *testing.T) {
	b, err := RandomBlock(0)

	assert.Nil(t, err)

	buf := &bytes.Buffer{}

	assert.Nil(t, b.Encode(NewGobBlockEncoder(buf)))

	bDecode := new(Block)

	assert.Nil(t, bDecode.Decode(NewGobBlockDecoder(buf)))
}
