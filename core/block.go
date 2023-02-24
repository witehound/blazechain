package core

import (
	"encoding/binary"
	"io"

	"github.com/witehound/blazechain/types"
)

type Header struct {
	Version   uint32
	Prevlock  types.Hash
	TimeStamp int64
	Height    uint32
	Nonce     uint64
}

type Block struct {
	Header
	Transactions []Transaction
}

func (h *Header) EncodeBinary(w io.Writer) error {
	if err := binary.Write(w, binary.LittleEndian, &h.Version); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, &h.Prevlock); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, &h.TimeStamp); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, &h.Height); err != nil {
		return err
	}

	return binary.Write(w, binary.LittleEndian, &h.Nonce)
}

func (h *Header) DecodeBinary(r io.Reader) error {
	if err := binary.Read(r, binary.LittleEndian, &h.Version); err != nil {
		return err
	}
	if err := binary.Read(r, binary.LittleEndian, &h.Prevlock); err != nil {
		return err
	}
	if err := binary.Read(r, binary.LittleEndian, &h.TimeStamp); err != nil {
		return err
	}
	if err := binary.Read(r, binary.LittleEndian, &h.Height); err != nil {
		return err
	}

	return binary.Read(r, binary.LittleEndian, &h.Nonce)

}

func (b *Block) EncodeBinary(w io.Writer) error {
	if err := b.Header.EncodeBinary(w); err != nil {
		return err
	}
	for _, tx := range b.Transactions {
		if err := tx.EncodeBinary(w); err != nil {
			return err
		}
	}

	return nil
}

func (b *Block) DecodeBinary(r io.Reader) error {
	if err := b.Header.DecodeBinary(r); err != nil {
		return err
	}
	for _, tx := range b.Transactions {
		if err := tx.DecodeBinary(r); err != nil {
			return err
		}
	}

	return nil
}
