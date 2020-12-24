package types

type MetaData struct {
	BaseModel
	Name        string `gorm:"not null;" json:"name"`
	Value       string `gorm:"not null;" json:"value"`
	Type        string `gorm:"not null;" json:"type"`
	IsActive    bool   `gorm:"not null;" json:"is_active"`
	PlatformKey string `gorm:"" json:"platform_key"`
}
