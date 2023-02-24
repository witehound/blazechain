package network

import (
	"fmt"
	"sync"
)

type LocalTransport struct {
	addr        NetAdd
	peers       map[NetAdd]*LocalTransport
	lock        sync.RWMutex
	consumeChan chan RPC
}

func NewLocalTransport(addr NetAdd) *LocalTransport {
	return &LocalTransport{
		addr:        addr,
		consumeChan: make(chan RPC, 1024),
		peers:       make(map[NetAdd]*LocalTransport),
	}
}

func (t *LocalTransport) Consume() <-chan RPC {
	return t.consumeChan
}

func (t *LocalTransport) Connet(tr LocalTransport) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.peers[tr.Addr()] = tr

	return nil
}

func (t *LocalTransport) SendMessage(to NetAdd, payLoad []byte) error {
	t.lock.RLock()
	defer t.lock.RUnlock()

	peer, ok := t.peers[to]
	if !ok {
		return fmt.Errorf("%s, Could not send message to person %s", t.addr, to)
	}
	peer.consumeChan <- RPC{
		From:    t.addr,
		Payload: payLoad,
	}
	return nil
}

func (t *LocalTransport) Addr() NetAdd {
	return t.addr
}
