package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
)

type PrivateKey struct {
	Key *ecdsa.PrivateKey
}

type PublicKey struct {
	Key *ecdsa.PublicKey
}

type Signature struct {
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

func (k PrivateKey) GeneratePublicKey() PublicKey {
	return PublicKey{
		Key: &k.Key.PublicKey,
	}
}
