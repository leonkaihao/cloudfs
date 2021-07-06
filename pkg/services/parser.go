package services

type Parser interface {
	ParseDir(path string) (int, error)
	ParseFile(path string) ([]byte, error)
}
