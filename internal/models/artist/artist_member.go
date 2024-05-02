package artist

import "kcloudb1/internal/config"

type ArtistMember struct {
	ID       int64  `json:"ID" gorm:"primary_key"`
	ArtistID int64  `json:"artist_id"`
	Name     string `json:"name"`
	SongID   int64  `json:"song_id"`
}

func (ArtistMember) TableName() string {
	return "artist_member"
}

func (c *ArtistMember) Create() error {
	return config.DB.Create(c).Error
}

func (c *ArtistMember) GetList() ([]ArtistMember, error) {
	var artistMembers []ArtistMember
	err := config.DB.Find(&artistMembers).Error
	return artistMembers, err
}

func (c *ArtistMember) Update() error {
	return config.DB.Updates(c).Error
}

func (c *ArtistMember) Delete() error {
	return config.DB.Delete(c).Error
}

func (c *ArtistMember) Get() error {
	return config.DB.Where("id = ?", c.ID).First(c).Error
}

func (c *ArtistMember) GetListByArtistID(artistID int64) ([]ArtistMember, error) {
	var artistMembers []ArtistMember
	err := config.DB.Where("artist_id = ?", artistID).Find(&artistMembers).Error
	return artistMembers, err
}

func (c *ArtistMember) GetListBySongID(songID int64) ([]ArtistMember, error) {
	var artistMembers []ArtistMember
	err := config.DB.Where("song_id = ?", songID).Find(&artistMembers).Error
	return artistMembers, err
}

func (c *ArtistMember) GetListByName() ([]ArtistMember, error) {
	var artistMembers []ArtistMember
	err := config.DB.Where("name = ?", c.Name).Find(&artistMembers).Error
	return artistMembers, err
}
