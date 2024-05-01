package song

type SongCategory struct {
	ID       int64  `json:"ID" gorm:"primary_key"`
	Ordering int64  `json:"ordering"`
	Category string `json:"category"`
}

func (c *SongCategory) TableName() string {
	return "song_category"
}

type SongCategoryCombination struct {
	ID             int64 `json:"ID" gorm:"primary_key"`
	SongCategoryID int64 `json:"song_category_id"`
	SongID         int64 `json:"song_id"`
}
