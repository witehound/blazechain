package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeypair_Sign_Verify_Succes(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.GetPublicKey()

	msg := []byte("hello world")

	sig, err := privKey.Sign(msg)

	assert.Nil(t, err)

	tr := sig.Verify(pubKey, msg)

	assert.True(t, tr)
}

func TestKeypair_Sign_Verify_Failed(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.GetPublicKey()

	privKeyTwo := GeneratePrivateKey()
	pubKeyTwo := privKeyTwo.GetPublicKey()

	msg := []byte("hello world")

	msgTwo := []byte("get ddown")

	sig, err := privKey.Sign(msg)
	sigTwo, errTwo := privKeyTwo.Sign(msgTwo)

	assert.Nil(t, err)
	assert.Nil(t, errTwo)

	tr := sig.Verify(pubKey, msgTwo)
	trTwo := sigTwo.Verify(pubKeyTwo, msg)

	assert.False(t, tr)
	assert.False(t, trTwo)

	assert.True(t, sig.Verify(pubKey, msg))
	assert.True(t, sigTwo.Verify(pubKeyTwo, msgTwo))
}
