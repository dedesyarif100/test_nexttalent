package position

type Service interface {
	FindAll() ([]Position, error)
	FindById(ID int) (*Position, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Position, error) {
	positions, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return positions, nil
}

func (s *service) FindById(ID int) (*Position, error) {
	position, err := s.repository.FindById(ID)
	if err != nil {
		return nil, err
	}
	return position, err
}
