package core

import "github.com/witehound/blazechain/types"

type Header struct {
	Version   uint32
	Prevlock  types.Hash
	TimeStamp uint64
	Height    uint32
	nonce     uint64
}

type Block struct {
	Header
	Transactions []Transaction
}
