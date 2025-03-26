package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// Company モデル
type Company struct {
	ID           string         `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name         string         `gorm:"type:varchar(255);not null"`
	URL          datatypes.JSON `gorm:"type:jsonb"` // JSONB で保存
	PhoneNumber  string         `gorm:"type:varchar(20)"`
	PostalCodeID int            `gorm:"type:integer"`
	PostalCode   PostalCode     `gorm:"foreignKey:PostalCodeID;references:ID"`
	Address      string         `gorm:"type:varchar(255)"`
	CreatedAt    time.Time      `gorm:"default:now()"`
	UpdatedAt    time.Time      `gorm:"default:now()"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`

	// 多対多のリレーション
	Tags []Tag `gorm:"many2many:tag_companies;"`
}
