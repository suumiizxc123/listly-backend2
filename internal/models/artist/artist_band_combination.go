package artist

import "kcloudb1/internal/config"

type ArtistBandCombination struct {
	ID       int64 `json:"ID" gorm:"primary_key"`
	BandID   int64 `json:"band_id"`
	ArtistID int64 `json:"artist_id"`
}

func (c *ArtistBandCombination) TableName() string {
	return "artist_band_combination"
}

func (c *ArtistBandCombination) Create() error {
	return config.DB.Create(c).Error
}

func (c *ArtistBandCombination) Update() error {
	return config.DB.Updates(c).Error
}

func (c *ArtistBandCombination) Delete() error {
	return config.DB.Delete(c).Error
}

func (c *ArtistBandCombination) Get() error {
	return config.DB.Where("id = ?", c.ID).First(c).Error
}

func (c *ArtistBandCombination) GetList() ([]ArtistBandCombination, error) {
	var orgs []ArtistBandCombination

	if err := config.DB.Find(&orgs).Error; err != nil {
		return nil, err
	}

	return orgs, nil
}
