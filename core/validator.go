package core

import "fmt"

type Validator interface {
	ValidateBlock(*Block) error
}

type BlockValidator struct {
	bc *BlockChain
}

func NewBlockValidator(bc *BlockChain) *BlockValidator {
	return &BlockValidator{
		bc: bc,
	}
}

func (bv *BlockValidator) ValidateBlock(b *Block) error {

	if bv.bc.Height() != b.Header.Height-1 {
		return fmt.Errorf("invalid block height")
	}

	prevHeader, err := bv.bc.BlockHeader(b.Header.Height - 1)

	if err != nil {
		return fmt.Errorf("invalid block hash")
	}

	hash := BlockHasher{}.Hash(prevHeader)

	if hash != b.Header.PrevBlockHash {
		return fmt.Errorf("invalid hash of previous block")
	}

	if err := b.Verify(); err != nil {
		return err
	}

	return nil
}
