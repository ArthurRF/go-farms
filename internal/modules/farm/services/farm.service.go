package service

import (
	"go-farms/internal/entity"
	repository "go-farms/internal/modules/farm/repositories"
)

type FarmServiceInterface interface {
	Create(farm *entity.Farm) (*entity.Farm, error)
	List() []entity.Farm
	Delete(id int) error
}

type FarmService struct {
	farmRepository repository.FarmRepositoryInterface
}

func GetFarmService(repo repository.FarmRepositoryInterface) FarmServiceInterface {
	return &FarmService{
		farmRepository: repo,
	}
}

func (f *FarmService) Create(farm *entity.Farm) (*entity.Farm, error) {
	return f.farmRepository.Create(farm)
}

func (f *FarmService) List() []entity.Farm {
	return f.farmRepository.List()
}

func (f *FarmService) Delete(id int) error {
	return f.farmRepository.Delete(id)
}
