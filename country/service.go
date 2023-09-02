package country

type Service interface {
	FindById(name string) (*Country, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindById(name string) (*Country, error) {
	country, err := s.repository.FindById(name)
	if err != nil {
		return nil, err
	}
	return country, err
}
