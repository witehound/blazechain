package main

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"math/rand"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/witehound/blazechain/core"
	"github.com/witehound/blazechain/crypto"
	"github.com/witehound/blazechain/network"
)

func main() {
	trLocal := network.NewLocalTransport("LOCAL")
	trRemoteA := network.NewLocalTransport("REMOTE_A")
	trRemoteB := network.NewLocalTransport("REMOTE_B")
	trRemoteC := network.NewLocalTransport("REMOTE_C")

	trLocal.Connect(trRemoteA)
	trRemoteA.Connect(trRemoteB)
	trRemoteB.Connect(trRemoteC)

	trRemoteA.Connect(trLocal)

	MakeRemoteServers([]network.Transport{trRemoteA, trRemoteB, trRemoteC})

	go func() {
		for {
			if err := SendTransaction(trRemoteA, trLocal.Addr()); err != nil {
				logrus.Error(err)
			}
			time.Sleep(2 * time.Second)
		}

	}()

	privKey := crypto.GeneratePrivateKey()

	loacalServer := MakeServer("LOCAL", trLocal, &privKey)

	loacalServer.Start()

}

func SendTransaction(tr network.Transport, to network.NetAdd) error {
	tx := core.NewTransactionWithSig(strconv.Itoa(rand.Intn(1000000000)))

	buf := &bytes.Buffer{}

	if err := tx.Encode(core.NewGobTxEncoder(buf)); err != nil {
		return err
	}

	msg := network.NewMessage(network.MessageTypeTx, buf.Bytes())

	return tr.SendMessage(to, msg.Bytes())

}

func MakeServer(id string, tr network.Transport, privKey *crypto.PrivateKey) *network.Server {
	opts := network.ServerOpts{
		PrivateKey: privKey,
		Transports: []network.Transport{tr},
		ID:         id,
	}

	s, err := network.NewServer(opts)
	if err != nil {
		log.Fatal(err)
	}

	return s
}

func MakeRemoteServers(trs []network.Transport) *network.Server {
	for i := 0; i < len(trs); i++ {
		id := fmt.Sprintf("REMOTE_%d", i+1)
		server := MakeServer(id, trs[i], nil)
		go server.Start()
	}
	return nil
}
