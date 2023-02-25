package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"math/big"

	"github.com/witehound/blazechain/types"
)

type Signature struct {
	r, s *big.Int
}

type PrivateKey struct {
	Key *ecdsa.PrivateKey
}

type PublicKey struct {
	Key *ecdsa.PublicKey
}

func GeneratePrivateKey() PrivateKey {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	return PrivateKey{
		Key: key,
	}
}

func (k PrivateKey) GetPublicKey() PublicKey {
	return PublicKey{
		Key: &k.Key.PublicKey,
	}
}

func (k PrivateKey) Sign(data []byte) (*Signature, error) {
	r, s, err := ecdsa.Sign(rand.Reader, k.Key, data)
	if err != nil {
		panic(err)
	}

	return &Signature{
		r: r, s: s,
	}, nil
}

func (k PublicKey) ToSlice() []byte {
	return elliptic.MarshalCompressed(k.Key, k.Key.X, k.Key.Y)
}

func (k PublicKey) Address() types.Address {
	h := sha256.Sum256(k.ToSlice())
	return types.AddressFromByte(h[len(h)-20:])
}

func (s Signature) Verify(pub PublicKey, data []byte) bool {
	return ecdsa.Verify(pub.Key, data, s.r, s.s)
}
