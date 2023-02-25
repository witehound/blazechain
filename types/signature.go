package types

import "math/big"

type Signature struct {
	r, s *big.Int
}
