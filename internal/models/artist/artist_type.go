package artist

import "kcloudb1/internal/config"

type ArtistType struct {
	ID   int64  `json:"ID" gorm:"primary_key"`
	Type string `json:"type"`
}

func (c *ArtistType) TableName() string {
	return "artist_type"
}

func (c *ArtistType) Create() error {
	return config.DB.Create(c).Error
}

func (c *ArtistType) GetList() ([]ArtistType, error) {
	var list []ArtistType
	err := config.DB.Find(&list).Error
	return list, err
}

func (c *ArtistType) GetListByType() ([]ArtistType, error) {
	var list []ArtistType
	err := config.DB.Where("type = ?", c.Type).Find(&list).Error
	return list, err
}

func (c *ArtistType) Get() error {
	return config.DB.Where("id = ?", c.ID).First(c).Error
}

func (c *ArtistType) Update() error {
	return config.DB.Updates(c).Error
}

func (c *ArtistType) Delete() error {
	return config.DB.Delete(c).Error
}
