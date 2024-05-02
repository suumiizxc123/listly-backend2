package song

import "kcloudb1/internal/config"

type SongCategoryCombination struct {
	ID             int64 `json:"ID" gorm:"primary_key"`
	SongCategoryID int64 `json:"song_category_id"`
	SongID         int64 `json:"song_id"`
}

func (c *SongCategoryCombination) TableName() string {
	return "song_category_combination"
}

func (c *SongCategoryCombination) Create() error {
	return config.DB.Create(c).Error
}

func (c *SongCategoryCombination) GetList() ([]SongCategoryCombination, error) {
	var list []SongCategoryCombination
	err := config.DB.Find(&list).Error
	return list, err
}

func (c *SongCategoryCombination) Update() error {
	return config.DB.Updates(c).Error
}

func (c *SongCategoryCombination) Delete() error {
	return config.DB.Delete(c).Error
}

func (c *SongCategoryCombination) Get() error {
	return config.DB.Where("id = ?", c.ID).First(c).Error
}

func (c *SongCategoryCombination) GetListBySongCategoryID() ([]SongCategoryCombination, error) {
	var list []SongCategoryCombination
	err := config.DB.Where("song_category_id = ?", c.SongCategoryID).Find(&list).Error
	return list, err
}

func (c *SongCategoryCombination) GetListBySongID() ([]SongCategoryCombination, error) {
	var list []SongCategoryCombination
	err := config.DB.Where("song_id = ?", c.SongID).Find(&list).Error
	return list, err
}
