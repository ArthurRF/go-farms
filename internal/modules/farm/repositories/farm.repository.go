package repository

import (
	"go-farms/internal/entity"

	"gorm.io/gorm"
)

type FarmRepositoryInterface interface {
	Create(farm *entity.Farm, tx *gorm.DB) (*entity.Farm, error)
}

type FarmRepository struct {
	db *gorm.DB
}

func GetFarmRepository(db *gorm.DB) FarmRepositoryInterface {
	return &FarmRepository{
		db: db,
	}
}

func (f *FarmRepository) Create(farm *entity.Farm, tx *gorm.DB) (*entity.Farm, error) {
	err := f.db.Create(farm).Error
	if err != nil {
		return nil, err
	}
	return farm, nil
}
