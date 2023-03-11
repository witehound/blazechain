package core

import (
	"fmt"
	"sync"

	"github.com/go-kit/log"
)

type BlockChain struct {
	Lock      sync.RWMutex
	Headers   []*Header
	Store     Storage
	Validator Validator
	Logger    log.Logger
}

func NewBlockChain(genesis *Block) (*BlockChain, error) {
	bc := &BlockChain{
		Headers: []*Header{},
		Store:   NewMemoryStote(),
	}

	bc.Validator = NewBlockValidator(bc)

	bc.AddBlockWithoutValidator(genesis)

	return bc, nil
}

func (bc *BlockChain) SetValidator(v Validator) {
	bc.Validator = v
}

func (bc *BlockChain) AddBlock(b *Block) error {

	if b == nil {
		return fmt.Errorf("invalid block type")
	}

	if err := bc.Validator.ValidateBlock(b); err != nil {
		return err
	}

	bc.AddBlockWithoutValidator(b)

	return nil
}

func (bc *BlockChain) Height() uint32 {
	bc.Lock.RLock()
	defer bc.Lock.RUnlock()
	return uint32(len(bc.Headers) - 1)
}

func (bc *BlockChain) HasBlock(height uint32) bool {
	return height <= bc.Height()
}

func (bc *BlockChain) AddBlockWithoutValidator(b *Block) error {

	bc.Lock.RLock()
	bc.Headers = append(bc.Headers, &b.Header)
	bc.Lock.RUnlock()

	bc.Logger.Log("msg", "adding block", "height", b.Header.Height)

	return bc.Store.Put(b)

}

func (bc *BlockChain) BlockHeader(h uint32) (*Header, error) {
	if !bc.HasBlock(h) {

		return nil, fmt.Errorf("given height (%d) too high", h)
	}

	bc.Lock.RLock()
	defer bc.Lock.RUnlock()
	return bc.Headers[h], nil
}
