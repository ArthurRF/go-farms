package entity

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Farm struct {
	ID            int       `json:"id" gorm:"primaryKey;type:int;autoIncrement"`
	FarmName      string    `json:"farm_name" gorm:"not null"`
	LandArea      int       `json:"land_area" gorm:"not null"`
	UnitOfMeasure string    `json:"unit_of_measure" gorm:"not null"`
	Address       string    `json:"address" gorm:"not null"`
	CreatedAt     time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`

	Crops []Crop `json:"crops" gorm:"OnDelete:CASCADE;foreignKey:FarmID;references:ID"`
}

func (f *Farm) AfterDelete(tx *gorm.DB) (err error) {
	tx.Clauses(clause.Returning{}).Where("farm_id = ?", f.ID).Delete(&Crop{})
	return
}

func (Farm) TableName() string {
	return "farms"
}
