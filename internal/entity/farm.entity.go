package entity

import "time"

type Farm struct {
	ID            int       `json:"id" gorm:"primaryKey;type:int;autoIncrement"`
	FarmName      string    `json:"farm_name" gorm:"not null"`
	LandArea      int       `json:"land_area" gorm:"not null"`
	UnitOfMeasure string    `json:"unit_of_measure" gorm:"not null"`
	Address       string    `json:"address" gorm:"not null"`
	CreatedAt     time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`

	Crops []Crop `json:"crops" gorm:"foreignKey:FarmID;references:ID"`
}

func (Farm) TableName() string {
	return "farms"
}