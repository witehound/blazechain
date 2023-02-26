package network

import (
	"crypto"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/witehound/blazechain/core"
)

type ServerOpts struct {
	Transports []Transport
	PrivateKey *crypto.PrivateKey
	BlockTime  time.Duration
}

type Server struct {
	ServerOpts
	isValidator bool
	rpcCh       chan RPC
	quitCh      chan struct{}
	MemePool    *MemePool
	BlockTime   time.Duration
}

func NewServer(opts ServerOpts) *Server {

	return &Server{
		ServerOpts:  opts,
		rpcCh:       make(chan RPC),
		isValidator: opts.PrivateKey != nil,
		quitCh:      make(chan struct{}, 1),
		MemePool:    NewMemePool(),
		BlockTime:   opts.BlockTime,
	}
}

func (s *Server) Start() {
	s.InitTransports()
	ticker := time.NewTicker(s.BlockTime)

free:
	for {
		select {
		case rpc := <-s.rpcCh:
			fmt.Printf("%v\n", rpc)
		case <-s.quitCh:
			break free
		case <-ticker.C:
			if s.isValidator {
				s.CreateNewBlock()
			}

		}
	}

	fmt.Println("Server Shutdown")
}

func (s *Server) HandleTransaction(tx *core.Transaction) error {
	hash := tx.Hash(core.TxHasher{})
	if err := tx.VerifyTx(); err != nil {
		return err
	}

	if s.MemePool.Has(hash) {
		logrus.WithField("HandleTransaction", logrus.Fields{
			"hash": hash,
		}).Info("tx already in the memepool")
		return nil
	}

	logrus.WithField("HandleTransaction", logrus.Fields{
		"hash": hash,
	}).Info("adding new tx  to memepool")

	return s.MemePool.AddTx(hash, tx)
}

func (s *Server) CreateNewBlock() error {
	fmt.Println("creating a new block")

	// mp := s.MemePool

	s.MemePool.Flush()

	return nil
}

func (s *Server) InitTransports() {
	for _, tr := range s.Transports {
		go func(tr Transport) {
			for rpc := range tr.Consume() {
				s.rpcCh <- rpc
			}
		}(tr)
	}
}
