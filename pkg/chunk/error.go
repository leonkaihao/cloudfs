package chunk

import "errors"

var (
	ErrAddSliceToChunk     = errors.New("failed to add slice to Chunk")
	ErrCheckChunkHashValue = errors.New("failed to check chunk hash value")
)
