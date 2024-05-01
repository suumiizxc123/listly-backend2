package artist

type ArtistSong struct {
	ID           int64 `json:"ID" gorm:"primary_key"`
	ProfileID    int64 `json:"profile_id"`
	ArtistTypeID int64 `json:"artist_type_id"`
	SongID       int64 `json:"song_id"`
}
