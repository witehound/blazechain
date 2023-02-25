package types

import "fmt"

type Address [20]uint8

func (a Address) ToSlice() []byte {
	b := make([]byte, 20)
	for i := 0; i < 20; i++ {
		b[i] = a[i]
	}
	return b
}

func NewAddressFromByte(b []byte) Address {
	if len(b) != 32 {
		msg := fmt.Sprintf("length should be 20 bubt got %d", len(b))
		panic(msg)
	}

	var value [20]uint8

	for i := 0; i < 32; i++ {
		value[i] = b[i]
	}

	return Address(value)
}
