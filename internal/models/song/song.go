package song

import "time"

type Song struct {
	ID          int64     `json:"ID" gorm:"primary_key"`
	CategoryID  int64     `json:"category_id"`
	UUID        string    `json:"uuid"`
	Name        string    `json:"name"`
	Url         string    `json:"url"`
	Duration    float32   `json:"duration"`
	Thumbnail   string    `json:"thumbnail"`
	LanguageID  int64     `json:"language_id"`
	ReleaseDate time.Time `json:"release_date"`
	IsActive    int64     `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
}

func (c *Song) TableName() string {
	return "song"
}
