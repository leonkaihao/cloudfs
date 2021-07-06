package repository

type Hashes interface {
	HashItemByPath(path string) (HasheItem, error)
}
