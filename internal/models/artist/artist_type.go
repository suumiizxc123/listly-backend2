package artist

type ArtistType struct {
	ID   int64  `json:"ID" gorm:"primary_key"`
	Type string `json:"type"`
}

func (c *ArtistType) TableName() string {
	return "artist_type"
}
