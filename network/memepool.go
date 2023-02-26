package network

import (
	"github.com/witehound/blazechain/core"
	"github.com/witehound/blazechain/types"
)

type MemePool struct {
	Transactions map[types.Hash]*core.Transaction
}

func NewMemePool() *MemePool {
	return &MemePool{
		Transactions: make(map[types.Hash]*core.Transaction),
	}
}

func (mp *MemePool) Len() int {
	return len(mp.Transactions)
}

func (mp *MemePool) Flush() {
	mp.Transactions = make(map[types.Hash]*core.Transaction)
}

func (mp *MemePool) AddTx(hash types.Hash, tx *core.Transaction) error {

	mp.Transactions[hash] = tx
	return nil
}

func (mp *MemePool) Has(h types.Hash) bool {
	_, ok := mp.Transactions[h]
	return ok
}
