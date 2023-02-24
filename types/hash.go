package types

import (
	"crypto/rand"
	"fmt"
)

type Hash [32]uint8

func HashFromByte(b []byte) Hash {
	if len(b) != 32 {
		msg := fmt.Sprintf("length should be 32 bubt got %d", len(b))
		panic(msg)
	}

	var value [32]uint8

	for i := 0; i < 32; i++ {
		value[i] = b[i]
	}

	return Hash(value)
}

func RandomByte(size int) []byte {
	token := make([]byte, size)
	rand.Read(token)
	return token
}

func RandomHash() Hash {
	return HashFromByte(RandomByte(32))
}
