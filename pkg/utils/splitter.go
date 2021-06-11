package utils

import (
	"context"
	"io"
)

const (
	DefaultSliceSize = 1024 * 256
)

type Splitter struct {
	MaxSliceSize int
}

func NewSplitter(maxSliceSize int) *Splitter {
	if maxSliceSize == 0 {
		maxSliceSize = DefaultSliceSize
	}
	return &Splitter{
		MaxSliceSize: maxSliceSize,
	}
}

func (spltr *Splitter) Split(ctx context.Context, reader io.Reader) (<-chan []byte, <-chan error) {
	slices := make(chan []byte)
	errChan := make(chan error)
	go func() {
		defer close(slices)
		defer close(errChan)
		for {
			buf := make([]byte, spltr.MaxSliceSize)
			n, err := reader.Read(buf)
			if err != nil {
				if err != io.EOF {
					errChan <- err
				}
				return
			}

			select {
			case slices <- buf[:n]:
			case <-ctx.Done():
				errChan <- ctx.Err()
				return
			}
		}
	}()

	return slices, errChan
}
