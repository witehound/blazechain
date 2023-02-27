package network

import "github.com/witehound/blazechain/core"

type MessageType byte

const (
	MessageTypeTx MessageType = 0x1
	MessageTypeBlock
)

type RPC struct {
	From    NetAdd
	Payload []byte
}

type Message struct {
	Header MessageType
	Data   []byte
}

type RPCHandler interface {
	HandleRPC(rpc RPC) error
}

type RPCProcessor interface {
	ProcessTransaction(NetAdd, *core.Transaction) error
}

type DefaultRPCHandler struct {
	p RPCProcessor
}
