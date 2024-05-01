package artist

type ArtistMember struct {
	ID       int64  `json:"ID" gorm:"primary_key"`
	ArtistID int64  `json:"artist_id"`
	Name     string `json:"name"`
	SongID   int64  `json:"song_id"`
}

func (ArtistMember) TableName() string {
	return "artist_member"
}

type ArtistMemberSong struct {
	ID             int64 `json:"ID" gorm:"primary_key"`
	ArtistMemberID int64 `json:"artist_member_id"`
	SongID         int64 `json:"song_id"`
}

func (ArtistMemberSong) TableName() string {
	return "artist_member_song"
}
