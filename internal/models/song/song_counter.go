package song

import "kcloudb1/internal/config"

type SongCounter struct {
	ID        int64 `json:"ID" gorm:"primary_key"`
	SongID    int64 `json:"song_id"`
	KaraokeID int64 `json:"karaoke_id"`
	RoomID    int64 `json:"room_id"`
}

func (c *SongCounter) TableName() string {
	return "song_counter"
}

func (c *SongCounter) Create() error {
	return config.DB.Create(c).Error
}

func (c *SongCounter) Update() error {
	// return config.DB.Updates(c).Error
	return config.DB.Where("id = ?", c.ID).First(c).Error
}

func (c *SongCounter) Delete() error {
	return config.DB.Delete(c).Error
}

func (c *SongCounter) Get() error {
	return config.DB.Where("id = ?", c.ID).First(c).Error
}

func (c *SongCounter) GetList() ([]SongCounter, error) {
	var orgs []SongCounter

	if err := config.DB.Find(&orgs).Error; err != nil {
		return nil, err
	}

	return orgs, nil
}
