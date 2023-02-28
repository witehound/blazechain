package main

import (
	"bytes"

	"math/rand"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/witehound/blazechain/core"
	"github.com/witehound/blazechain/crypto"
	"github.com/witehound/blazechain/network"
)

func main() {
	trLocal := network.NewLocalTransport("LOCAL")
	trRemote := network.NewLocalTransport("REMOTE")

	trLocal.Connect(trRemote)
	trRemote.Connect(trLocal)

	go func() {
		for {
			// trRemote.SendMessage(trLocal.Addr(), []byte("hello world"))
			if err := SendTransaction(trRemote, trLocal.Addr()); err != nil {
				logrus.Error(err)
			}
			time.Sleep(1 * time.Second)
		}
	}()

	privKey := crypto.GeneratePrivateKey()

	opts := network.ServerOpts{
		PrivateKey: privKey,
		Transports: []network.Transport{trLocal},
		ID:         "LOCAL",
	}

	s, err := network.NewServer(opts)

	if err != nil {
		logrus.Error(err)
	}

	s.Start()

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
