package models

// TagCompany (企業とタグのリレーション)
type TagCompany struct {
	CompanyID string `gorm:"type:uuid;not null;primaryKey;foreignKey:CompanyID;references:ID"`
	TagID     string `gorm:"type:uuid;not null;primaryKey;foreignKey:TagID;references:ID"`
}

