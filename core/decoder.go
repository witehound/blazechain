package core

import (
	"encoding/gob"
	"io"
)

type Decoder[T any] interface {
	Decode(T) error
}

type GobTxDecoder struct {
	r io.Reader
}

type GobBlockDecoder struct {
	r io.Reader
}

func NewGobTxDecoder(r io.Reader) *GobTxDecoder {

	return &GobTxDecoder{
		r: r,
	}
}

func (e *GobTxDecoder) Decode(tx *Transaction) error {
	return gob.NewDecoder(e.r).Decode(tx)

}

func NewGobBlockDecoder(r io.Reader) *GobBlockDecoder {

	return &GobBlockDecoder{
		r: r,
	}
}

func (e *GobBlockDecoder) Decode(tx *Block) error {
	return gob.NewDecoder(e.r).Decode(tx)

}
