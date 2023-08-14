package employmenttype

type Service interface {
	FindAll() ([]EmploymentType, error)
	FindById(ID int) (*EmploymentType, error)
	// Create(employmentTypeRequest EmploymentTypeRequest) error
	// Update(ID int, employmentTypeRequest EmploymentTypeRequest) error
	// Delete(ID int) error
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
	return employeeType, err
}

// func (s *service) Create(employmentTypeRequest EmploymentTypeRequest) error {
// 	book := EmploymentType{
// 		Name: employmentTypeRequest.Name,
// 	}

// 	_, err := s.repository.Create(book)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (s *service) Update(ID int, employmentTypeRequest EmploymentTypeRequest) error {
// 	employmentType, err := s.repository.FindById(ID)
// 	if err != nil {
// 		return err
// 	}
// 	employmentType.Name = employmentTypeRequest.Name
// 	_, err = s.repository.Update(employmentType)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (s *service) Delete(ID int) error {
// 	book, err := s.repository.FindById(ID)
// 	if err != nil {
// 		return err
// 	}
// 	_, err = s.repository.Delete(book)
// 	if err != nil {
// 		return err
// 	}
// 	return err
// }
