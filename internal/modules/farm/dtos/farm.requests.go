package dtos

import (
	"go-farms/internal/entity"
)

type CreateFarmRequest struct {
	FarmName      string              `json:"farm_name" validate:"required"`
	LandArea      int                 `json:"land_area" validate:"required"`
	UnitOfMeasure string              `json:"unit_of_measure" validate:"required"`
	Address       string              `json:"address" validate:"required"`
	Crops         []CreateCropRequest `json:"crops"`
}

type CreateCropRequest struct {
	CropType    entity.CropType `json:"crop_type" validate:"required"`
	IsIrrigated bool            `json:"is_irrigated" validate:"required"`
	IsInsured   bool            `json:"is_insured" validate:"required"`
}

func ValidateCropType(ccr *CreateCropRequest) bool {
	for _, cropType := range entity.CropTypes {
		if cropType == ccr.CropType {
			return true
		}
	}
	return false
}
