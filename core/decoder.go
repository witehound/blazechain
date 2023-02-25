package core

import "io"

type Decoder[T any] interface {
	Decode(io.Reader, T) error
}
