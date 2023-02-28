package core

import (
	"time"

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
		PrevBlockHash: types.RandomHash(),
		Height:        0,
		TimeStamp:     time.Now().UnixNano(),
	}

	b, err := NewBlock(Header(*h), []*Transaction{})
	if err != nil {
		return nil, err
	}

	return b, nil
}

func CheckOptsToSignBlock(key *crypto.PrivateKey) crypto.PrivateKey {
	if key != nil {
		return crypto.GeneratePrivateKey()
	}
	return *key
}
