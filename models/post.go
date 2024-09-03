package models

import (
	"gorm.io/gorm"
	"time"
)

type Post struct {
	ID        uint `gorm:"primarykey"`
	Title     string
	Content   string
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at"`
}
