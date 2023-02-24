package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra.Connet(trb)
	trb.Connet(tra)

	assert.Equal(t, tra.peers[trb.addr], trb)

	assert.Equal(t, trb.peers[tra.addr], tra)
}

func TestSendMessagge(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra.Connet(trb)
	trb.Connet(tra)

	msg := []byte("heLlo B")

	assert.Nil(t, tra.SendMessage(trb.addr, msg))

	rpc := <-trb.Consume()

	assert.Equal(t, rpc.Payload, msg)
	assert.Equal(t, rpc.From, tra.addr)

}
