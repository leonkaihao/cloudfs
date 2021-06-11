package utils

import (
	"crypto/md5"
	"hash"
)

type Encoder interface {
	Consume([]byte) error
	Hash() []byte
}

type encoder struct {
	h hash.Hash
}

func (enc *encoder) Consume(data []byte) error {
	_, err := enc.h.Write(data)
	return err
}

func (enc *encoder) Hash() []byte {
	return enc.h.Sum(nil)
}

func NewEncoder() Encoder {
	return &encoder{
		h: md5.New(),
	}
}
