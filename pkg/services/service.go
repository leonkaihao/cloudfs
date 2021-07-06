package services

import "errors"

type Service interface {
	Repo
	Parser
	History
}

type service struct {
}

func New() Service {
	return &service{}
}

func (svc *service) GetHistory(latest int) ([]string, error) {
	return nil, errors.New("Not implemented")
}

func (svc *service) ParseDir(path string) (int, error) {
	return 0, errors.New("Not implemented")
}

func (svc *service) ParseFile(path string) ([]byte, error) {
	return nil, errors.New("Not implemented")
}

func (svc *service) Init(path string) error {
	return errors.New("Not implemented")
}

func (svc *service) MetaInfo() (map[string]string, error) {
	return nil, errors.New("Not implemented")
}
