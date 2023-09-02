package country

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	FindByName(name string) (*Country, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByName(name string) (*Country, error) {
	var country Country
	err := r.db.Find(&country, "name", name).Error
	fmt.Println("ERR : ", err)
	if country.Name == "" {
		err = errors.New("data not found")
		return nil, err
	}
	return &country, nil
}
