package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func StartNewBlockChainWithGenesis(t *testing.T) *BlockChain {
	bc, err := NewBlockChain(RandomBlock(0))
	assert.Nil(t, err)

	return bc
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
		bc.AddBlock(bc.RandomBlockWithSig(uint32(i + 1)))
	}

	assert.Equal(t, bc.Height(), ct)
}

func TestValidator(t *testing.T) {
	bc := StartNewBlockChainWithGenesis(t)

	assert.NotNil(t, bc.Validator)
	var ct uint32 = 0

	for i := 0; i < 10; i++ {
		ct++
		bc.AddBlock(bc.RandomBlockWithSig(uint32(i + 1)))
	}

	assert.Nil(t, bc.AddBlock(bc.RandomBlockWithSig(ct+1)))

	assert.NotNil(t, bc.AddBlock(bc.RandomBlockWithSig(3)))
	assert.NotNil(t, bc.AddBlock(bc.RandomBlockWithSig(14)))

}

func TestBlockHeeder(t *testing.T) {
	bc := StartNewBlockChainWithGenesis(t)

	var ct uint32 = 0

	for i := 0; i < 1; i++ {
		ct++
		tb := bc.RandomBlockWithSig(uint32(i + 1))
		bc.AddBlock(tb)
		h, err := bc.BlockHeader(tb.Header.Height)
		assert.Nil(t, err)
		assert.Equal(t, &tb.Header, h)

	}
}
