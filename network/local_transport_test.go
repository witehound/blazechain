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
