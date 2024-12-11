//go:build wireinject
// +build wireinject

package injector

import (
	farmRepo "go-farms/internal/modules/farm/repositories"
	farmService "go-farms/internal/modules/farm/services"

	"github.com/google/wire"
	"gorm.io/gorm"
)

var FarmServiceSet = wire.NewSet(
	farmRepo.GetFarmRepository,
	farmService.GetFarmService,
)

func InitializeFarmService(db *gorm.DB) farmService.FarmServiceInterface {
	wire.Build(FarmServiceSet)
	return nil
}
