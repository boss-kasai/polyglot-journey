package models

import (
	"time"

	"gorm.io/gorm"
)

// PostalCode テーブル
type PostalCode struct {
	ID         int            `gorm:"primaryKey;autoIncrement" json:"id"`
	PostalCode string         `gorm:"type:varchar(10);not null;unique" json:"postal_code"`
	Address    string         `gorm:"type:varchar(255);not null" json:"address"`
	CreatedAt  time.Time      `gorm:"default:now()" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"default:now()" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index"` // ソフトデリート対応
}
