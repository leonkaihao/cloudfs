package chunk

import (
	"bytes"
	"crypto/md5"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type SliceHeader struct {
	ID        uuid.UUID // used for identifying a slice
	Size      int32     // slice size
	Seq       int64     // sequence number of the slice in a chunk
	HashValue []byte    // 16 bytes hash for the slice

	Metadata map[string]interface{} // can include like partitionID
}
type SliceData struct {
	Data []byte
}
type SliceStorage struct {
	Path string
}
type Slice struct {
	SliceHeader
	SliceData
	SliceStorage
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
		SliceData: SliceData{
			Data: data,
		},
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
