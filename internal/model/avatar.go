package model

import (
	"time"

	"gorm.io/gorm"
)

type Avatar struct {
	Id        uint   `gorm:"primarykey"`
	Hash      string `gorm:"unique;not null"`
	ImageData []byte `gorm:"type:mediumblob"`
	ImageURL  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (m *Avatar) TableName() string {
	return "avatar"
}
