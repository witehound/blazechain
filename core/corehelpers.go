package core

import (
	"time"

	"github.com/go-kit/log"
	"github.com/witehound/blazechain/crypto"
	"github.com/witehound/blazechain/types"
)

func StartNewBlockChainWithGenesis(privKey crypto.PrivateKey) (*BlockChain, error) {

	b, err := GenesisBlock()
	if err != nil {
		return nil, err
	}

	b.Sign(privKey)

	bc, err := NewBlockChain(b)

	if err != nil {
		return nil, err
	}

	return bc, nil
}

func RandomBlock(height uint32) (*Block, error) {
	h := &Header{
		Version:       1,
		PrevBlockHash: types.RandomHash(),
		Height:        height,
		TimeStamp:     time.Now().UnixNano(),
	}

	tx := NewTransactionWithSig("food")
	b, err := NewBlock(Header(*h), []*Transaction{tx})
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GenesisBlock() (*Block, error) {
	h := &Header{
		Version:       1,
		PrevBlockHash: types.Hash{},
		Height:        0,
		TimeStamp:     0000000,
	}

	b, err := NewBlock(Header(*h), []*Transaction{})
	if err != nil {
		return nil, err
	}

	return b, nil
}

func StartNewBlockChainGenesisLogger(privKey crypto.PrivateKey, logger log.Logger) (*BlockChain, error) {

	b, err := GenesisBlock()
	if err != nil {
		return nil, err
	}

	b.Sign(privKey)

	bc := &BlockChain{
		Headers: []*Header{},
		Store:   NewMemoryStote(),
		Logger:  logger,
	}

	bc.Validator = NewBlockValidator(bc)

	bc.AddBlockWithoutValidator(b)

	bc.Logger = logger

	if err != nil {
		return nil, err
	}

	return bc, nil
}
