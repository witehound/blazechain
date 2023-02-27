package network

type NetAdd string

type Transport interface {
	Connect(Transport) error
	Consume() <-chan RPC
	SendMessage(NetAdd, []byte) error
	Addr() NetAdd
}
