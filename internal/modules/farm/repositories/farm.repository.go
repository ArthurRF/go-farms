package repository

import (
	"go-farms/internal/entity"

	"gorm.io/gorm"
)

type FarmRepositoryInterface interface {
	Create(farm *entity.Farm) (*entity.Farm, error)
	List() []entity.Farm
	Delete(id int) error
}

type FarmRepository struct {
	db *gorm.DB
}

func GetFarmRepository(db *gorm.DB) FarmRepositoryInterface {
	return &FarmRepository{
		db: db,
	}
}

func (f *FarmRepository) Create(farm *entity.Farm) (*entity.Farm, error) {
	err := f.db.Create(farm).Error
	if err != nil {
		return nil, err
	}
	return farm, nil
}

func (f *FarmRepository) List() []entity.Farm {
	var farms []entity.Farm
	f.db.Preload("Crops").Find(&farms)
	return farms
}

func (f *FarmRepository) Delete(id int) error {
	err := f.db.Where("farm_id = ?", id).Delete(&entity.Crop{}).Error
	if err != nil {
		return err
	}
	return f.db.Delete(&entity.Farm{}, id).Error
}
