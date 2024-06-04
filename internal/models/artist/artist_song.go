package artist

import "kcloudb1/internal/config"

type ArtistSong struct {
	ID           int64 `json:"ID" gorm:"primary_key"`
	ProfileID    int64 `json:"profile_id"`
	ArtistBandID int64 `json:"artist_band_id"`
	SongID       int64 `json:"song_id"`
}

func (c *ArtistSong) TableName() string {
	return "artist_song"
}

func (c *ArtistSong) Create() error {

	return config.DB.Create(c).Error
}

func (c *ArtistSong) GetList() ([]ArtistSong, error) {

	var list []ArtistSong
	err := config.DB.Find(&list).Error
	return list, err
}

func (c *ArtistSong) Update() error {

	return config.DB.Updates(c).Error
}

func (c *ArtistSong) Delete() error {

	return config.DB.Delete(c).Error
}

func (c *ArtistSong) GetListBySongID() ([]ArtistSong, error) {

	var list []ArtistSong

	err := config.DB.Where("song_id = ?", c.SongID).Find(&list).Error

	return list, err
}

func (c *ArtistSong) GetListByProfileID() ([]ArtistSong, error) {

	var list []ArtistSong

	err := config.DB.Where("profile_id = ?", c.ProfileID).Find(&list).Error

	return list, err
}

func (c *ArtistSong) Get() error {

	return config.DB.Where("id = ?", c.ID).First(c).Error
}
