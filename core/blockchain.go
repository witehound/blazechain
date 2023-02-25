package core

type BlockChain struct {
	Headers   []*Header
	Store     Storage
	Validator Validator
}

func NewBlockChain() *BlockChain {
	bc := &BlockChain{
		Headers: []*Header{},
	}

	bc.Validator = NewBlockValidator(bc)

	return bc
}

func (bc *BlockChain) SetValidator(v Validator) {
	bc.Validator = v
}

func (bc *BlockChain) AddBlock() error {
	return nil
}

func (bc *BlockChain) Height() uint32 {
	return uint32(len(bc.Headers) - 1)
}
