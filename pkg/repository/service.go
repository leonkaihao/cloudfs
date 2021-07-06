package repository

type Service interface {
	Actions() Actions
	Hashes() Hashes
	Watches() Watches
}

type service struct {
}

func (svc *service) Validate(path string) error {
	return nil
}

func (svc *service) Init(path string) error {
	return nil
}

func (svc *service) Load(path string) error {
	return nil
}

func (svc *service) Actions() Actions {
	return nil
}

func (svc *service) Hashes() Hashes {
	return nil
}

func (svc *service) Watches() Watches {
	return nil
}

func New(path string) Service {
	svc := &service{}
	return svc
}
