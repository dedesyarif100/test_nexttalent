package position

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Position, error)
	FindById(ID int) (*Position, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Position, error) {
	var position []Position
	err := r.db.Find(&position).Error
	if err != nil {
		return nil, err
	}
	return position, err
}

func (r *repository) FindById(ID int) (*Position, error) {
	var position Position
	err := r.db.Find(&position, ID).Error
	if position.ID == 0 {
		err = errors.New("data not found")
		return nil, err
	}
	return &position, nil
}
