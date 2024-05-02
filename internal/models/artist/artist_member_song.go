package artist

import "kcloudb1/internal/config"

type ArtistMemberSong struct {
	ID             int64 `json:"ID" gorm:"primary_key"`
	ArtistMemberID int64 `json:"artist_member_id"`
	SongID         int64 `json:"song_id"`
}

func (c *ArtistMemberSong) TableName() string {
	return "artist_member_song"
}

func (c *ArtistMemberSong) Create() error {
	return config.DB.Create(c).Error
}

func (c *ArtistMemberSong) GetList() ([]ArtistMemberSong, error) {
	var list []ArtistMemberSong
	err := config.DB.Find(&list).Error
	return list, err
}

func (c *ArtistMemberSong) Update() error {
	return config.DB.Updates(c).Error
}

func (c *ArtistMemberSong) Delete() error {
	return config.DB.Delete(c).Error
}

func (c *ArtistMemberSong) Get() error {
	return config.DB.Where("id = ?", c.ID).First(c).Error
}

func (c *ArtistMemberSong) GetListByArtistMemberID() ([]ArtistMemberSong, error) {
	var list []ArtistMemberSong
	err := config.DB.Where("artist_member_id = ?", c.ArtistMemberID).Find(&list).Error
	return list, err
}

func (c *ArtistMemberSong) GetListBySongID() ([]ArtistMemberSong, error) {
	var list []ArtistMemberSong
	err := config.DB.Where("song_id = ?", c.SongID).Find(&list).Error
	return list, err
}
