package model

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Opportunity struct {
	Id        int64          `gorm:"primaryKey" json:"id"`
	Data      datatypes.JSON `gorm:"type:jsonb" json:"data"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

func (op *Opportunity) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()
	op.CreatedAt = now
	op.UpdatedAt = now	
	return nil
}

func (op *Opportunity) BeforeUpdate(tx *gorm.DB) error {
	op.UpdatedAt = time.Now()
	return nil
}