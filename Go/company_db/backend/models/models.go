package models

import "gorm.io/gorm"

// MigrateAll - すべてのモデルをマイグレーション
func MigrateAll(db *gorm.DB) error {
	return db.AutoMigrate(&Company{}, &PostalCode{}, &Tag{}, &TagCompany{})
}

