package main

import (
	"github.com/witehound/blazechain/network"
)

func main() {
	trLocal := network.NewLocalTransport("lpone")

	opts := network.ServerOpts{
		Transports: []network.Transport{trLocal},
	}

	s := network.NewServer(opts)
	s.Start()
}
