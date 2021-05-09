package chunk

import (
	"context"
	"crypto/md5"
	"io"
)

const (
	DefaultSliceSize = 1024 * 256
)

type Splitter struct {
	MaxSliceSize int
}

func NewSplitter(maxSliceSize int) *Splitter {
	return &Splitter{
		MaxSliceSize: maxSliceSize,
	}
}

func (spltr *Splitter) SplitMemData(ctx context.Context, data []byte, startSeq int64) (<-chan *Slice, <-chan error) {
	slices := make(chan *Slice)
	errChan := make(chan error)
	go func() {
		defer close(slices)
		index := startSeq
		for start := 0; start < len(data); start += spltr.MaxSliceSize {
			end := start + spltr.MaxSliceSize
			if end > len(data) {
				end = len(data)
			}
			dataSeg := data[start:end]
			hashValue := md5.New().Sum(dataSeg)
			slice, err := NewSlice(index, dataSeg, hashValue)
			if err != nil {
				errChan <- err
				return
			}
			index++

			select {
			case slices <- slice:
			case <-ctx.Done():
				errChan <- ctx.Err()
				return
			}
		}
	}()
	return slices, errChan
}

func (spltr *Splitter) SplitIOData(ctx context.Context, reader io.Reader, startSeq int64) (<-chan *Slice, <-chan error) {
	slices := make(chan *Slice)
	errChan := make(chan error)
	go func() {
		defer close(slices)
		buf := make([]byte, spltr.MaxSliceSize)
		index := startSeq
		_, err := reader.Read(buf)
		for err != nil {
			hashValue := md5.New().Sum(buf)

			slice, err := NewSlice(index, buf, hashValue)
			if err != nil {
				errChan <- err
				return
			}

			select {
			case slices <- slice:
			case <-ctx.Done():
				errChan <- ctx.Err()
				return
			}

			index++
			buf = make([]byte, spltr.MaxSliceSize)
			_, err = reader.Read(buf)
		}
	}()

	return slices, errChan
}
