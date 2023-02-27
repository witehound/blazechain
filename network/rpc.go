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

type DecodedMsg struct {
	from NetAdd
	Data any
}

type RPCDecodeFunc func(RPC) (*DecodedMsg, error)

type RPCProcessor interface {
	ProcessMessage(*DecodedMsg) error
}

func DefaultRPCDecodeFunc(rpc RPC) (*DecodedMsg, error) {

	msg := Message{}
	if err := gob.NewDecoder(rpc.Payload).Decode(&msg); err != nil {
		return nil, fmt.Errorf("failed to decode message from this pair %s : %s", rpc.Payload, err)
	}

	switch msg.Header {
	case MessageTypeTx:
		tx := new(core.Transaction)
		if err := tx.Decode(core.NewGobTxDecoder(bytes.NewReader(msg.Data))); err != nil {
			return nil, err
		}

		return &DecodedMsg{
			from: rpc.From,
			Data: tx,
		}, nil

	default:
		return nil, fmt.Errorf("invalid message header")
	}

}

func NewMessage(t MessageType, data []byte) *Message {
	return &Message{
		Header: t,
		Data:   data,
	}
}

func (msg *Message) Bytes() []byte {

	buf := &bytes.Buffer{}
	gob.NewEncoder(buf).Encode(msg)
	return buf.Bytes()
}
