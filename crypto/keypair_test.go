package crypto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenarate_privateKey(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.GetPublicKey()
	address := pubKey.Address()

	msg := []byte("hello world")

	sig, err := privKey.Sign(msg)

	assert.Nil(t, err)

	tr := sig.Verify(pubKey, msg)

	assert.True(t, tr)

	fmt.Println(address)
}
