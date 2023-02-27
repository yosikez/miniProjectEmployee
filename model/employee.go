package model

import (
	"time"

	"gorm.io/gorm"
)

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

type Employee struct {
	ID        int64  `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"type:varchar(100)" json:"name"`
	Email     string `gorm:"type:varchar(100)" json:"email"`
	Address   string `gorm:"type:text" json:"address"`
	Phone     string `gorm:"type:varchar(13)" json:"phone"`
	Gender    Gender `gorm:"type:enum('male','female');enum:'male,female'" json:"gender"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
}

func (em *Employee) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()
	em.CreatedAt = now
	em.UpdatedAt = now
	return nil
}

func (em *Employee) BeforeUpdate(tx *gorm.DB) error {
	em.UpdatedAt = time.Now()
	return nil
}