package core

import (
	"crypto"

	"github.com/witehound/blazechain/types"
)

type Header struct {
	Version      uint32
	PrevlockHash types.Hash
	TimeStamp    int64
	Height       uint32
	DataHash     types.Hash
}

type Block struct {
	Header
	Transactions []Transaction
	hash         types.Hash
	Validator    crypto.PublicKey
	Signature    *types.Signature
}

func (b *Block) Hash(hasher Hasher[*Block]) types.Hash {
	if b.hash.FindCachedHash() {
		b.hash = hasher.Hash(b)
	}

	return b.hash
}

// func (h *Header) EncodeBinary(w io.Writer) error {
// 	if err := binary.Write(w, binary.LittleEndian, &h.Version); err != nil {
// 		return err
// 	}
// 	if err := binary.Write(w, binary.LittleEndian, &h.PrevlockHash); err != nil {
// 		return err
// 	}
// 	if err := binary.Write(w, binary.LittleEndian, &h.TimeStamp); err != nil {
// 		return err
// 	}
// 	if err := binary.Write(w, binary.LittleEndian, &h.Height); err != nil {
// 		return err
// 	}

// 	return binary.Write(w, binary.LittleEndian, &h.Nonce)
// }

// func (h *Header) DecodeBinary(r io.Reader) error {
// 	if err := binary.Read(r, binary.LittleEndian, &h.Version); err != nil {
// 		return err
// 	}
// 	if err := binary.Read(r, binary.LittleEndian, &h.PrevlockHash); err != nil {
// 		return err
// 	}
// 	if err := binary.Read(r, binary.LittleEndian, &h.TimeStamp); err != nil {
// 		return err
// 	}
// 	if err := binary.Read(r, binary.LittleEndian, &h.Height); err != nil {
// 		return err
// 	}

// 	return binary.Read(r, binary.LittleEndian, &h.Nonce)

// }

// func (b *Block) EncodeBinary(w io.Writer) error {
// 	if err := b.Header.EncodeBinary(w); err != nil {
// 		return err
// 	}
// 	for _, tx := range b.Transactions {
// 		if err := tx.EncodeBinary(w); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// func (b *Block) DecodeBinary(r io.Reader) error {
// 	if err := b.Header.DecodeBinary(r); err != nil {
// 		return err
// 	}
// 	for _, tx := range b.Transactions {
// 		if err := tx.DecodeBinary(r); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }
