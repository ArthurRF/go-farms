package service

import (
	"go-farms/internal/entity"
	repository "go-farms/internal/modules/farm/repositories"

	"gorm.io/gorm"
)

type FarmServiceInterface interface {
	Create(farm *entity.Farm, tx *gorm.DB) (*entity.Farm, error)
}

type FarmService struct {
	farmRepository repository.FarmRepositoryInterface
}

func GetFarmService(repo repository.FarmRepositoryInterface) FarmServiceInterface {
	return &FarmService{
		farmRepository: repo,
	}
}

func (f *FarmService) Create(farm *entity.Farm, tx *gorm.DB) (*entity.Farm, error) {
	return f.farmRepository.Create(farm, tx)
}
