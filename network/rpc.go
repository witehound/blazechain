package network

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"

	"github.com/witehound/blazechain/core"
)

type MessageType byte

const (
	MessageTypeTx MessageType = 0x1
	MessageTypeBlock
)

type RPC struct {
	From    NetAdd
	Payload io.Reader
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

func NewDefaultRPCHandler(p RPCProcessor) *DefaultRPCHandler {
	return &DefaultRPCHandler{
		p: p,
	}
}

func (rh *DefaultRPCHandler) HandleRPC(rpc RPC) error {
	msg := Message{}
	if err := gob.NewDecoder(rpc.Payload).Decode(&msg); err != nil {
		return err
	}

	switch msg.Header {
	case MessageTypeTx:
		tx := new(core.Transaction)
		if err := tx.Decode(core.NewGobTxDecoder(bytes.NewReader(msg.Data))); err != nil {
			return nil
		}
		return rh.p.ProcessTransaction(rpc.From, tx)
	default:
		return fmt.Errorf("invalid message header")
	}

}
