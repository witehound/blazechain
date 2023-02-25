package core

type BlockChain struct {
	Headers []*Header
	Store   Storage
}

func (bc *BlockChain) Height() uint32 {
	return uint32(len(bc.Headers) - 1)
}
