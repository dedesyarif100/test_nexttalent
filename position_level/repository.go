package positionlevel

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]PositionLevel, error)
	FindById(ID int) (*PositionLevel, error)
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
	if err != nil {
		return nil, err
	}
	return positionLevels, err
}

func (r *repository) FindById(ID int) (*PositionLevel, error) {
	var positionLevel PositionLevel
	err := r.db.Find(&positionLevel, ID).Error
	if positionLevel.ID == 0 {
		err = errors.New("data not found")
		return nil, err
	}
	return &positionLevel, nil
}
