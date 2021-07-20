package repository

type HashSchema struct {
}

type Hashes interface {
	HashItemByPath(path string) (*HashSchema, error)
}
