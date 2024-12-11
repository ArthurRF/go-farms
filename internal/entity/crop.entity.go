package entity

import "time"

type Crop struct {
	ID          int       `json:"id" gorm:"primaryKey;type:int;autoIncrement"`
	FarmID      int       `json:"farm_id" gorm:"not null"`
	CropType    CropType  `json:"crop_type" gorm:"not null"`
	IsIrrigated bool      `json:"is_irrigated" gorm:"not null"`
	IsInsured   bool      `json:"is_insured" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`

	Farm Farm `json:"-" gorm:"foreignKey:FarmID;references:ID"`
}

type CropType string

const (
	RICE     CropType = "RICE"
	BEANS    CropType = "BEANS"
	CORN     CropType = "CORN"
	COFFEE   CropType = "COFFEE"
	SOYBEANS CropType = "SOYBEANS"
)

func (Crop) TableName() string {
	return "crops"
}
