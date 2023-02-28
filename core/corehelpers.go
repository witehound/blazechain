package core

import (
	"time"

	"github.com/witehound/blazechain/crypto"
	"github.com/witehound/blazechain/types"
)

func StartNewBlockChainWithGenesis(privKey crypto.PrivateKey) (*BlockChain, error) {

	b := RandomBlock(0)

	dataHash, err := CalculateDataHash(b.Transactions)

	if err != nil {
		return nil, err
	}

	b.Header.DataHash = dataHash

	b.Sign(privKey)

	bc, err := NewBlockChain(b)

	if err != nil {
		return nil, err
	}

	return bc, nil
}

func RandomBlock(height uint32) *Block {
	h := &Header{
		Version:       1,
		PrevBlockHash: types.RandomHash(),
		Height:        height,
		TimeStamp:     time.Now().UnixNano(),
	}

	tx := NewTransactionWithSig("food")

	return NewBlock(Header(*h), []Transaction{*tx})
}
