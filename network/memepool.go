package network

import (
	"sort"

	"github.com/witehound/blazechain/core"
	"github.com/witehound/blazechain/types"
)

type TxMapSorter struct {
	transactions []*core.Transaction
}

func NewTxMapSorter(txMap map[types.Hash]*core.Transaction) *TxMapSorter {
	tsx := make([]*core.Transaction, len(txMap))

	i := 0
	for _, value := range txMap {
		tsx[i] = value
		i++
	}

	s := &TxMapSorter{transactions: tsx}

	sort.Sort(s)

}

func (t *TxMapSorter) Len() int {
	return len(t.transactions)
}

func (t *TxMapSorter) Swap(i, j int) {
	t.transactions[i], t.transactions[j] = t.transactions[j], t.transactions[i]
}
func (t *TxMapSorter) Less(i, j int) {
	t.transactions[i], t.transactions[j] = t.transactions[j], t.transactions[i]
}

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

func (mp *MemePool) AllTransactions() []core.Transaction {

}
