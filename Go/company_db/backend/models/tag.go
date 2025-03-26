package models

import (
	"time"

	"gorm.io/gorm"
)

// Tag テーブル
type Tag struct {
	ID        string         `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string         `gorm:"type:varchar(100);not null;unique"`
	CreatedAt time.Time      `gorm:"default:now()"`
	UpdatedAt time.Time      `gorm:"default:now()"`
	DeletedAt gorm.DeletedAt `gorm:"index"` // ソフトデリート対応

	// 多対多のリレーション
	Companies []Company `gorm:"many2many:tag_companies;"`
}
