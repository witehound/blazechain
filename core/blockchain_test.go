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

	for i := 0; i < 5; i++ {
		ct++
		bc.AddBlock(RandomBlock(uint32(i + 1)))
	}

	assert.Equal(t, bc.Height(), ct)
}
