package network

import (
	"crypto"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/witehound/blazechain/core"
)

var defaultBlockTime = 5 * time.Second

type ServerOpts struct {
	Transports []Transport
	PrivateKey *crypto.PrivateKey
	BlockTime  time.Duration
	RPCHandler
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

	if opts.BlockTime == time.Duration(0) {
		opts.BlockTime = defaultBlockTime
	}

	s := &Server{
		ServerOpts:  opts,
		rpcCh:       make(chan RPC),
		isValidator: opts.PrivateKey != nil,
		quitCh:      make(chan struct{}, 1),
		MemePool:    NewMemePool(),
		BlockTime:   opts.BlockTime,
	}

	if opts.RPCHandler == nil {
		s.ServerOpts.RPCHandler = NewDefaultRPCHandler(s)
	}

	return s
}

func (s *Server) Start() {
	s.InitTransports()
	ticker := time.NewTicker(s.BlockTime)

free:
	for {
		select {
		case rpc := <-s.rpcCh:
			if err := s.RPCHandler.HandleRPC(rpc); err != nil {
				logrus.Error(err)
			}
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

func (s *Server) ProcessTransaction(from NetAdd, tx *core.Transaction) error {
	hash := tx.Hash(core.TxHasher{})

	if s.MemePool.Has(hash) {
		logrus.WithField("HandleTransaction", logrus.Fields{
			"hash": hash,
		}).Info("tx already in the memepool")
		return nil
	}

	if err := tx.VerifyTx(); err != nil {
		return err
	}

	tx.SetFirstSeen(time.Now().UnixNano())

	logrus.WithField("HandleTransaction", logrus.Fields{
		"hash": hash,
	}).Info("adding new tx  to memepool")

	return s.MemePool.AddTx(hash, tx)
}

func (s *Server) CreateNewBlock() error {
	fmt.Println("creating a new block")

	// s.MemePool.Flush()

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
