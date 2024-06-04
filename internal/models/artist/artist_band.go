package artist

import "kcloudb1/internal/config"

type ArtistBand struct {
	ID   int64  `json:"ID" gorm:"primary_key"`
	Name string `json:"name"`
}

func (c *ArtistBand) TableName() string {
	return "artist_band"
}

func (c *ArtistBand) Create() error {
	return config.DB.Create(c).Error
}

func (c *ArtistBand) Update() error {
	return config.DB.Updates(c).Error
}

func (c *ArtistBand) Delete() error {
	return config.DB.Delete(c).Error
}

func (c *ArtistBand) Get() error {
	return config.DB.Where("id = ?", c.ID).First(c).Error
}

func (c *ArtistBand) GetList() ([]ArtistBand, error) {
	var orgs []ArtistBand

	if err := config.DB.Find(&orgs).Error; err != nil {
		return nil, err
	}

	return orgs, nil
}
