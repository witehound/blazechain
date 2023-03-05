package core

import (
	"bytes"
	"math/rand"
	"strconv"
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

func TestTxEncodeDecode(t *testing.T) {
	privkey := crypto.GeneratePrivateKey()

	tx := NewTransactionWithSig(strconv.Itoa(rand.Intn(1000000000)))
	buf := &bytes.Buffer{}
	assert.Nil(t, tx.SignTx(privkey))

	assert.Nil(t, tx.Encode(NewGobTxEncoder(buf)))
	txDecoded := new(Transaction)
	assert.Nil(t, txDecoded.Decode(NewGobTxDecoder(buf)))
	assert.Equal(t, tx, txDecoded)
}
