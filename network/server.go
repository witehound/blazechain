package network

type ServerOpts struct {
	Transports []Transport
}

type Server struct {
	ServerOpts
	rpcCh chan RPC
}

func NewServer(opts ServerOpts) *Server {
	return &Server{
		ServerOpts: opts,
		rpcCh:      make(chan RPC),
	}
}

func (s *Server) start() {
	s.initTranspose()
}

func (s *Server) initTranspose() {
	for _, tr := range s.Transports {
		go func(tr Transport) {
			for rpc := range tr.Consume() {

			}
		}(tr)
	}
}
