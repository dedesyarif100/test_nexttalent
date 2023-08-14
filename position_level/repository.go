package positionlevel

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]PositionLevel, error)
	FindById(ID int) (*PositionLevel, error)
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

func (r *repository) FindAll() ([]PositionLevel, error) {
	var positionLevels []PositionLevel
	err := r.db.Find(&positionLevels).Error
	return positionLevels, err
}

func (r *repository) FindById(ID int) (*PositionLevel, error) {
	var positionLevel PositionLevel
	err := r.db.Find(&positionLevel, ID).Error
	return &positionLevel, err
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
