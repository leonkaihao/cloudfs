package chunk

import (
	"time"

	"github.com/google/uuid"
)

const (
	MetadataPartitionKey = "Partition"
)

type Config struct {
	MaxSliceSize int `json:"maxSliceSize"`
	Partitions   int `json:"partitions"`
}

// ChunkHeader
type ChunkHeader struct {
	ID        uuid.UUID // used for identifying a chunk
	CreateAt  time.Time // create time
	Size      int64     // chunk size
	HashValue []byte    // 16 bytes hash for the whole chunk
}

type SliceHeader struct {
	ID        uuid.UUID // used for identifying a slice
	Size      int32     // slice size
	Seq       int64     // sequence number of the slice in a chunk
	HashValue []byte    // 16 bytes hash for the slice

	Metadata map[string]interface{} // can include like partitionID
}
