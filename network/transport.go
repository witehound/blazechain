package network

type NetAdd string

type RPC struct {
	From    NetAdd
	Payload []byte
}

type Transport interface {
	Connect(Transport) error
	Consume() <-chan RPC
	SendMessage(NetAdd, []byte) error
	Addr() NetAdd
}
