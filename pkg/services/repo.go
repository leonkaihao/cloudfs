package services

type Repo interface {
	Init(path string) error
	MetaInfo() (map[string]string, error)
}
