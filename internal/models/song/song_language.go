package song

import "kcloudb1/internal/config"

type SongLanguage struct {
	ID       int64  `json:"ID" gorm:"primary_key"`
	Language string `json:"language"`
}

func (c *SongLanguage) TableName() string {
	return "song_language"
}

func (c *SongLanguage) Create() error {
	return config.DB.Create(c).Error
}

func (c *SongLanguage) Update() error {
	return config.DB.Updates(c).Error
}

func (c *SongLanguage) Delete() error {
	return config.DB.Delete(c).Error
}

func (c *SongLanguage) Get() error {
	return config.DB.Where("id = ?", c.ID).First(c).Error
}

func (c *SongLanguage) GetList() ([]SongLanguage, error) {
	var orgs []SongLanguage

	if err := config.DB.Find(&orgs).Error; err != nil {
		return nil, err
	}

	return orgs, nil
}
