package network

import (
	"bytes"
	"crypto"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/witehound/blazechain/core"
)

var defaultBlockTime = 5 * time.Second

type ServerOpts struct {
	Transports    []Transport
	PrivateKey    *crypto.PrivateKey
	BlockTime     time.Duration
	RPCDecodeFunc RPCDecodeFunc
	RPCProcessor  RPCProcessor
}

type Server struct {
	ServerOpts
	isValidator bool
	rpcCh       chan RPC
	quitCh      chan struct{}
	MemePool    *MemePool
}

func NewServer(opts ServerOpts) *Server {

	if opts.BlockTime == time.Duration(0) {
		opts.BlockTime = defaultBlockTime
	}

	if opts.RPCDecodeFunc == nil {
		opts.RPCDecodeFunc = DefaultRPCDecodeFunc
	}

	s := &Server{
		ServerOpts:  opts,
		rpcCh:       make(chan RPC),
		isValidator: opts.PrivateKey != nil,
		quitCh:      make(chan struct{}, 1),
		MemePool:    NewMemePool(),
	}

	if s.ServerOpts.RPCProcessor == nil {
		s.ServerOpts.RPCProcessor = s
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
			msg, err := s.RPCDecodeFunc(rpc)
			if err != nil {
				logrus.Error(err)
			}
			if err := s.RPCProcessor.ProcessMessage(msg); err != nil {
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

func (s *Server) ProcessTransaction(tx *core.Transaction) error {
	hash := tx.Hash(core.TxHasher{})

	if s.MemePool.Has(hash) {
		logrus.WithFields(logrus.Fields{
			"hash": hash,
		}).Info("tx already in the memepool")
		return nil
	}

	if err := tx.VerifyTx(); err != nil {
		return err
	}

	tx.SetFirstSeen(time.Now().UnixNano())

	logrus.WithFields(logrus.Fields{
		"hash":             hash,
		"memepool lenggth": s.MemePool.Len(),
	}).Info("adding new tx  to memepool")

	go s.BroadCastTx(tx)

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

func (s *Server) ProcessMessage(msg *DecodedMsg) error {

	switch t := msg.Data.(type) {
	case *core.Transaction:
		return s.ProcessTransaction(t)
	}

	return nil
}

func (s *Server) BroadCasting(msg []byte) error {
	for _, tr := range s.Transports {
		if err := tr.BroadCast(msg); err != nil {
			logrus.Error(err)
			return err
		}
	}
	return nil
}

func (s *Server) BroadCastTx(tx *core.Transaction) error {
	buf := &bytes.Buffer{}
	if err := tx.Encode(core.NewGobTxEncoder(buf)); err != nil {
		return err
	}

	msg := NewMessage(MessageTypeTx, buf.Bytes())
	return s.BroadCasting(msg.Bytes())
}
