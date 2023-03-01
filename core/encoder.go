package core

import (
	"crypto/elliptic"
	"encoding/gob"
	"io"
)

type Encoder[T any] interface {
	Encode(T) error
}

type GobTxEncoder struct {
	w io.Writer
}

func NewGobTxEncoder(w io.Writer) *GobTxEncoder {
	gob.Register(elliptic.P256())
	return &GobTxEncoder{
		w: w,
	}
}

func (e *GobTxEncoder) Encode(tx *Transaction) error {
	return gob.NewEncoder(e.w).Encode(tx)

}

type GobBlockEncoder struct {
	w io.Writer
}

func NewGobBlockEncoder(w io.Writer) *GobBlockEncoder {
	gob.Register(elliptic.P256())
	return &GobBlockEncoder{
		w: w,
	}
}

func (e *GobBlockEncoder) Encode(b *Block) error {
	return gob.NewEncoder(e.w).Encode(b)

}
