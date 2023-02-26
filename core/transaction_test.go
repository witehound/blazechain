package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/witehound/blazechain/crypto"
)

func TestSignTx(t *testing.T) {
	privkey := crypto.GeneratePrivateKey()
	data := []byte("foo")

	tx := &Transaction{
		Data: data,
	}

	assert.Nil(t, tx.SignTx(privkey))

	assert.NotNil(t, tx.Signature)
}

func TestVerifyTx(t *testing.T) {
	privkey := crypto.GeneratePrivateKey()
	data := []byte("foo")

	tx := &Transaction{
		Data: data,
	}

	assert.Nil(t, tx.SignTx(privkey))

	assert.Nil(t, tx.VerifyTx())

	privkeyTwo := crypto.GeneratePrivateKey()

	tx.From = privkeyTwo.GetPublicKey()

	assert.NotNil(t, tx.VerifyTx())

}

func RandomTxwIthSig(t *testing.T) *Transaction {
	privkey := crypto.GeneratePrivateKey()
	tx := &Transaction{
		Data: []byte("foo"),
	}

	assert.Nil(t, tx.SignTx(privkey))

	return tx
}
