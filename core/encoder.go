package core

import "io"

type Encoder[T any] interface {
	Encode(io.Writer, T) error
}
