package chunk

import (
	"bytes"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/leonkaihao/cloudfs/pkg/utils"
)

type ChunkHeader struct {
	ID        uuid.UUID // used for identifying a chunk
	CreateAt  time.Time // create time
	Size      int64     // chunk size
	HashValue []byte    // 16 bytes hash for the whole chunk
}

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
		return fmt.Errorf("%w: %v", ErrAddSliceToChunk, err)
	}
	if chk.slices[slice.Seq] != nil {
		return fmt.Errorf("%w: slice %v has already been used", ErrAddSliceToChunk, slice.Seq)
	}
	chk.slices[slice.Seq] = slice
	return nil
}

func (chk *Chunk) CheckHashValue() error {
	enc := utils.NewEncoder()
	for i, slice := range chk.slices {
		if slice == nil {
			return fmt.Errorf("%w: slice %v cannot be nil when calculating chunk hash value", ErrCheckChunkHashValue, i)
		}
		if err := enc.Consume(slice.Data); err != nil {
			return fmt.Errorf("%w: cannot consume slice %v", ErrCheckChunkHashValue, i)
		}
	}
	if bytes.Compare(enc.Hash(), chk.HashValue) != 0 {
		return fmt.Errorf("%w: chunk %v has MD5 checksum error", ErrCheckChunkHashValue, chk.ID)
	}
	return nil
}
