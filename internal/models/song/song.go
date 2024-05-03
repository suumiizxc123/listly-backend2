package song

import (
	"kcloudb1/internal/config"
	"kcloudb1/internal/models/common"
	"time"

	"gorm.io/gorm/clause"
)

type Song struct {
	ID int64 `json:"ID" gorm:"primary_key"`
	// CategoryID  int64     `json:"category_id"`
	UUID        string    `json:"uuid"`
	Name        string    `json:"name"`
	Url         string    `json:"url"`
	Duration    float32   `json:"duration"`
	Thumbnail   string    `json:"thumbnail"`
	LanguageID  int64     `json:"language_id"`
	ReleaseDate time.Time `json:"release_date"`
	IsActive    int64     `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
}

func (c *Song) TableName() string {
	return "song"
}

func (c *Song) Create() error {
	return config.DB.Create(c).Error
}

func (c *Song) Update() error {
	return config.DB.Updates(c).Error
}

func (c *Song) Delete() error {
	return config.DB.Delete(c).Error
}

func (c *Song) Get() error {
	return config.DB.Where("id = ?", c.ID).First(c).Error
}

func (c *Song) GetList() ([]Song, error) {
	var list []Song
	return list, config.DB.Find(&list).Error
}

type SongExtend struct {
	ID int64 `json:"ID" gorm:"primary_key"`
	// CategoryID  int64           `json:"category_id"`
	// Category    SongCategory    `json:"category" gorm:"foreignKey:CategoryID; references:ID"`
	Name        string          `json:"name"`
	Url         string          `json:"url"`
	Duration    float32         `json:"duration"`
	Thumbnail   string          `json:"thumbnail"`
	LanguageID  int64           `json:"language_id"`
	Language    common.Language `json:"language" gorm:"foreignKey:LanguageID; references:ID"`
	ReleaseDate time.Time       `json:"release_date"`
	IsActive    int64           `json:"is_active"`
	CreatedAt   time.Time       `json:"created_at"`
}

func (c *SongExtend) TableName() string {
	return "song"
}

func (c *SongExtend) Get() error {
	return config.DB.Preload(clause.Associations).Where("id = ?", c.ID).First(c).Error
}

func (c *SongExtend) GetList() ([]SongExtend, error) {
	var list []SongExtend
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
