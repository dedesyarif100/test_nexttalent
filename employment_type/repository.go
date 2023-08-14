package employmenttype

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]EmploymentType, error)
	FindById(ID int) (*EmploymentType, error)
	// Create(employmentType EmploymentType) (EmploymentType, error)
	// Update(employmentType EmploymentType) (EmploymentType, error)
	// Delete(employmentType EmploymentType) (EmploymentType, error)
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
	return employmentTypes, err
}

func (r *repository) FindById(ID int) (*EmploymentType, error) {
	var position EmploymentType
	err := r.db.Find(&position, ID).Error
	return &position, err
}

// func (r *repository) Create(book EmploymentType) (EmploymentType, error) {
// 	err := r.db.Create(&book).Error
// 	return book, err
// }

// func (r *repository) Update(book EmploymentType) (EmploymentType, error) {
// 	err := r.db.Save(&book).Error
// 	return book, err
// }

// func (r *repository) Delete(book EmploymentType) (EmploymentType, error) {
// 	err := r.db.Delete(&book).Error
// 	return book, err
// }
