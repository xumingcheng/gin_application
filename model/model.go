package model

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	gorm.DeletedAt `gorm:"index"`
}
