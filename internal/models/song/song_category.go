package song

import (
	"kcloudb1/internal/config"
)

type SongCategory struct {
	ID       int64  `json:"ID" gorm:"primary_key"`
	Ordering int64  `json:"ordering"`
	Category string `json:"category"`
}

func (c *SongCategory) TableName() string {
	return "song_category"
}

func (c *SongCategory) Create() error {
	return config.DB.Create(c).Error
}

func (c *SongCategory) GetList() ([]SongCategory, error) {
	var list []SongCategory
	err := config.DB.Find(&list).Error
	return list, err
}

func (c *SongCategory) Get() error {
	return config.DB.Where("id = ?", c.ID).First(c).Error
}

func (c *SongCategory) Update() error {
	return config.DB.Updates(c).Error
}

func (c *SongCategory) Delete() error {
	return config.DB.Delete(c).Error
}

type SongCategorySong struct {
	ID       int64        `json:"ID" gorm:"primary_key"`
	Ordering int64        `json:"ordering"`
	Category string       `json:"category"`
	Songs    []SongExtend `json:"songs" gorm:"foreignKey:CategoryID; references:ID"`
}

func (c *SongCategorySong) TableName() string {
	return "song_category"
}

func (c *SongCategorySong) Get() error {
	return config.DB.
		Preload("Songs").
		Preload("Songs.Category").
		Preload("Songs.Language").
		Where("id = ?", c.ID).First(c).Error
}
