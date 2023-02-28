package network

import (
	"bytes"

	"os"
	"time"

	"github.com/go-kit/log"

	"github.com/sirupsen/logrus"
	"github.com/witehound/blazechain/core"
	"github.com/witehound/blazechain/crypto"
)

var defaultBlockTime = 5 * time.Second

type ServerOpts struct {
	Transports    []Transport
	PrivateKey    *crypto.PrivateKey
	BlockTime     time.Duration
	RPCDecodeFunc RPCDecodeFunc
	RPCProcessor  RPCProcessor
	Logger        log.Logger
	ID            string
}

type Server struct {
	ServerOpts
	isValidator bool
	rpcCh       chan RPC
	quitCh      chan struct{}
	MemePool    *MemePool
	Chain       *core.BlockChain
}

func NewServer(opts ServerOpts) (*Server, error) {

	if opts.BlockTime == time.Duration(0) {
		opts.BlockTime = defaultBlockTime
	}

	if opts.RPCDecodeFunc == nil {
		opts.RPCDecodeFunc = DefaultRPCDecodeFunc
	}

	if opts.Logger == nil {
		opts.Logger = log.NewLogfmtLogger(os.Stderr)
		opts.Logger = log.With(opts.Logger, "ID", opts.ID)
	}

	var chain *core.BlockChain

	if opts.PrivateKey == nil {
		priv := crypto.GeneratePrivateKey()
		bc, err := core.StartNewBlockChainGenesisLogger(priv, opts.Logger)
		if err != nil {
			return nil, err
		}
		chain = bc
	} else {
		bc, err := core.StartNewBlockChainGenesisLogger(*opts.PrivateKey, opts.Logger)
		if err != nil {
			return nil, err
		}
		chain = bc
	}

	s := &Server{
		ServerOpts:  opts,
		rpcCh:       make(chan RPC),
		isValidator: opts.PrivateKey != nil,
		quitCh:      make(chan struct{}, 1),
		MemePool:    NewMemePool(),
		Chain:       chain,
	}

	if s.ServerOpts.RPCProcessor == nil {
		s.ServerOpts.RPCProcessor = s
	}

	if s.isValidator {
		go s.Validator()
	}

	return s, nil
}

func (s *Server) Start() {
	s.InitTransports()

free:
	for {
		select {
		case rpc := <-s.rpcCh:
			msg, err := s.RPCDecodeFunc(rpc)
			if err != nil {
				s.Logger.Log("error", err)
			}
			if err := s.RPCProcessor.ProcessMessage(msg); err != nil {
				s.Logger.Log("error", err)
			}
		case <-s.quitCh:
			break free

		}
	}

	s.Logger.Log("msg", "server is shutting down")
}

func (s *Server) ProcessTransaction(tx *core.Transaction) error {
	hash := tx.Hash(core.TxHasher{})

	if s.MemePool.Has(hash) {

		return nil
	}

	if err := tx.VerifyTx(); err != nil {
		return err
	}

	tx.SetFirstSeen(time.Now().UnixNano())

	s.Logger.Log("msg", "adding new tx  to memepool",
		"hash", hash,
		"memepool length", s.MemePool.Len())

	go s.BroadCastTx(tx)

	return s.MemePool.AddTx(hash, tx)
}

func (s *Server) CreateNewBlock() error {

	currHeader, err := s.Chain.BlockHeader(s.Chain.Height())
	if err != nil {
		return err
	}

	tsx := s.MemePool.AllTransactions()

	block, err := core.NewBlockFromPrevHeader(currHeader, tsx)
	if err != nil {
		return err
	}

	if err := block.Sign(*s.ServerOpts.PrivateKey); err != nil {
		return err
	}

	if err := s.Chain.AddBlock(block); err != nil {
		return err
	}

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

func (s *Server) Validator() {
	ticker := time.NewTicker(s.BlockTime)

	s.Logger.Log("msg", "starting validator loop", "blocktime", s.BlockTime)

	for {
		<-ticker.C
		s.CreateNewBlock()
	}

}
