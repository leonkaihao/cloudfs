package services

type History interface {
	GetHistory(latest int) ([]string, error)
}
