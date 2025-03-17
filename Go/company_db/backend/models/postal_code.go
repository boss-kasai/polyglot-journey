package models

import (
	"time"

	"gorm.io/gorm"
)

// PostalCode テーブル
type PostalCode struct {
	ID         int            `gorm:"primaryKey;autoIncrement"`
	PostalCode string         `gorm:"type:varchar(10);not null;unique"`
	Address    string         `gorm:"type:varchar(255);not null"`
	CreatedAt  time.Time      `gorm:"default:now()"`
	UpdatedAt  time.Time      `gorm:"default:now()"`
	DeletedAt  gorm.DeletedAt `gorm:"index"` // ソフトデリート対応
}

