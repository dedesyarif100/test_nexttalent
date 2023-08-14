package employmenttype

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]EmploymentType, error)
	FindById(ID int) (*EmploymentType, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]EmploymentType, error) {
	var employmentTypes []EmploymentType
	err := r.db.Find(&employmentTypes).Error
	if err != nil {
		return nil, err
	}
	return employmentTypes, nil
}

func (r *repository) FindById(ID int) (*EmploymentType, error) {
	var employmentType EmploymentType
	err := r.db.Find(&employmentType, ID).Error
	if employmentType.ID == 0 {
		err = errors.New("data not found")
		return nil, err
	}
	return &employmentType, nil
}
