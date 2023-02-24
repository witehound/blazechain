package main

import (
	"time"

	"github.com/witehound/blazechain/network"
)

func main() {
	trLocal := network.NewLocalTransport("lpone")
	trRemote := network.NewLocalTransport("rppone")

	trLocal.Connect(trRemote)
	trRemote.Connect(trLocal)

	go func() {
		for {
			trRemote.SendMessage(trLocal.Addr(), []byte("hello local"))
			time.Sleep(1 * time.Second)
		}
	}()

	opts := network.ServerOpts{
		Transports: []network.Transport{trLocal},
	}

	s := network.NewServer(opts)
	s.Start()
}
