package chunk

import (
	"bytes"
	"crypto/md5"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Slice struct {
	SliceHeader
	Data []byte
}

func NewSlice(seq int64, data []byte, hashValue []byte) (*Slice, error) {
	return &Slice{
		SliceHeader: SliceHeader{
			ID:        uuid.New(),
			Size:      int32(len(data)),
			Seq:       seq,
			HashValue: hashValue,
			Metadata:  make(map[string]interface{}),
		},
		Data: data,
	}, nil
}

func (slc *Slice) AddMetadata(key string, val interface{}) {
	slc.Metadata[key] = val
}

func (slc *Slice) Verify() error {
	if slc == nil {
		return errors.New("Slice is nil")
	}
	hashValue := md5.Sum(slc.Data)
	if bytes.Compare(hashValue[:], slc.HashValue) != 0 {
		return errors.Errorf("Slice %v has MD5 checksum error", slc.Seq)
	}
	return nil
}
