package positionlevel

type Service interface {
	FindAll() ([]PositionLevel, error)
	FindById(ID int) (*PositionLevel, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]PositionLevel, error) {
	positionLevels, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return positionLevels, nil
}

func (s *service) FindById(ID int) (*PositionLevel, error) {
	positionLevel, err := s.repository.FindById(ID)
	if err != nil {
		return nil, err
	}
	return positionLevel, err
}
