package core

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/witehound/blazechain/types"
)

func TestHeader_decode_encode(t *testing.T) {
	h := &Header{
		Version:   1,
		Prevlock:  types.RandomHash(),
		TimeStamp: time.Now().UnixNano(),
		Height:    10,
		Nonce:     985677,
	}

	buf := &bytes.Buffer{}
	assert.Nil(t, h.EncodeBinary(buf))

	hDecode := &Header{}
	assert.Nil(t, hDecode.DecodeBinary(buf))

	assert.Equal(t, h, hDecode)
}

func TestBlock_decode_encode(t *testing.T) {
	h := Header{
		Version:   1,
		Prevlock:  types.RandomHash(),
		TimeStamp: time.Now().UnixNano(),
		Height:    10,
		Nonce:     985677,
	}

	b := &Block{
		Header:       h,
		Transactions: nil,
	}

	buf := &bytes.Buffer{}
	assert.Nil(t, b.EncodeBinary(buf))

	bDecode := &Block{}
	assert.Nil(t, bDecode.DecodeBinary(buf))

	assert.Equal(t, b, bDecode)
}
