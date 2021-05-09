package chunk

import (
	"fmt"
)

const (
	ErrCreateChunk int = iota
	ErrAddSliceToChunk
	ErrCheckChunkHashValue
)

type ChunkError struct {
	Code int
	Err  error
}

func (r *ChunkError) Error() string {
	return fmt.Sprintf("ChunkError, Code: %v, Error: %v", r.Code, r.Err)
}
