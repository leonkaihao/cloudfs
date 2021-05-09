package chunk

import (
	"bytes"
	"crypto/md5"
	"io"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Chunk struct {
	ChunkHeader
	slices []*Slice
}

func NewChunk(size int64, sliceCount int64, hashValue []byte) (*Chunk, error) {
	return &Chunk{
		ChunkHeader: ChunkHeader{
			ID:        uuid.New(),
			CreateAt:  time.Now(),
			Size:      size,
			HashValue: hashValue,
		},
		slices: make([]*Slice, sliceCount),
	}, nil
}

func (chk *Chunk) AddSlice(slice *Slice) error {
	if err := slice.Verify(); err != nil {
		return &ChunkError{ErrAddSliceToChunk, err}
	}
	if chk.slices[slice.Seq] != nil {
		return &ChunkError{
			ErrAddSliceToChunk,
			errors.Errorf("Slice %v has already been used", slice.Seq),
		}
	}
	chk.slices[slice.Seq] = slice
	return nil
}

func (chk *Chunk) CheckHashValue() error {
	h := md5.New()
	for i, slice := range chk.slices {
		if slice == nil {
			return &ChunkError{
				ErrCheckChunkHashValue,
				errors.Errorf("Slice %v cannot be nil when calculating chunk hash value", i),
			}
		}
		io.WriteString(h, string(slice.Data))
	}
	if bytes.Compare(h.Sum(nil), chk.HashValue) != 0 {
		return &ChunkError{
			ErrCheckChunkHashValue,
			errors.Errorf("Chunk %v has MD5 checksum error", chk.ID),
		}
	}
	return nil
}
