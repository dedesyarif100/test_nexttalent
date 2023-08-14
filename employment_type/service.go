package employmenttype

type Service interface {
	FindAll() ([]EmploymentType, error)
	FindById(ID int) (*EmploymentType, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]EmploymentType, error) {
	employeeTypes, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return employeeTypes, nil
}

func (s *service) FindById(ID int) (*EmploymentType, error) {
	employeeType, err := s.repository.FindById(ID)
	if err != nil {
		return nil, err
	}
	return employeeType, nil
}
