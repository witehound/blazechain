package core

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func StartNewBlockChainWithGenesis(t *testing.T) *BlockChain {

	bc, err := NewBlockChain(RandomBlock(0))
	assert.Nil(t, err)

	return bc
}

func (bc *BlockChain) BlockWithHash(t *testing.T, height uint32) (*Block, error) {

	ph, err := bc.BlockHeader(height - 1)

	if err != nil {
		return nil, err
	}

	h := &Header{
		Version:       1,
		PrevBlockHash: BlockHasher{}.Hash(ph),
		Height:        height,
		TimeStamp:     time.Now().UnixNano(),
	}

	txs := RandomTxwIthSig(t)

	return NewBlock(Header(*h), []Transaction{*txs}), nil
}

func TestBlockChainInit(t *testing.T) {

	bc := StartNewBlockChainWithGenesis(t)

	assert.NotNil(t, bc.Validator)

	assert.Equal(t, bc.Height(), uint32(0))
}

func TestAddBlock(t *testing.T) {
	bc := StartNewBlockChainWithGenesis(t)

	assert.NotNil(t, bc.Validator)
	var ct uint32 = 0

	for i := 0; i < 1000; i++ {
		ct++
		tb, err := bc.RandomBlockWithSig(t, uint32(i+1))
		assert.Nil(t, err)
		bc.AddBlock(tb)
	}

	assert.Equal(t, bc.Height(), ct)
}

func TestValidator(t *testing.T) {
	bc := StartNewBlockChainWithGenesis(t)

	assert.NotNil(t, bc.Validator)
	var ct uint32 = 0

	for i := 0; i < 10; i++ {
		ct++
		tb, err := bc.RandomBlockWithSig(t, uint32(i+1))
		assert.Nil(t, err)
		bc.AddBlock(tb)
	}

	b1, err := bc.RandomBlockWithSig(t, ct+1)
	assert.Nil(t, err)

	assert.Nil(t, bc.AddBlock(b1))

	b2, err2 := bc.RandomBlockWithSig(t, 3)
	b3, err3 := bc.RandomBlockWithSig(t, 14)

	assert.Nil(t, err2)
	assert.NotNil(t, err3)

	assert.NotNil(t, bc.AddBlock(b2))
	assert.NotNil(t, bc.AddBlock(b3))

}

func TestBlockHeeder(t *testing.T) {
	bc := StartNewBlockChainWithGenesis(t)

	var ct uint32 = 0

	for i := 0; i < 1; i++ {
		ct++
		tb, err := bc.RandomBlockWithSig(t, uint32(i+1))
		assert.Nil(t, err)
		bc.AddBlock(tb)
		h, err := bc.BlockHeader(tb.Header.Height)
		assert.Nil(t, err)
		assert.Equal(t, &tb.Header, h)

	}
}
