package core

import "time"

type Project struct {
	UID          string    `json:"uid"`
	Address      string    `json:"address"`
	BlockNumber  string    `json:"block_number"`
	Category     string    `json:"category"`
	City         string    `json:"city"`
	District     string    `json:"district"`
	Khoroo       string    `json:"khoroo"`
	Description  string    `json:"description"`
	FacebookURL  string    `json:"facebook_url"`
	InstagramURL string    `json:"instagram_url"`
	IsActive     bool      `json:"is_active"`
	IsFeature    bool      `json:"is_feature"`
	LaunchDate   time.Time `json:"launch_date"`
	Parking      int64     `json:"parking"`
	Progress     float32   `json:"progress"`
	StartPrice   float32   `json:"start_price"`
	Title        string    `json:"title"`
	TotalFloors  int64     `json:"total_floors"`
	TotalUnits   int64     `json:"total_units"`
	TotalRoomMax int64     `json:"total_room_max"`
	TotalRoomMin int64     `json:"total_room_min"`
	Type         string    `json:"type"`
	VideoLink    string    `json:"video_link"`
	YouTubeURL   string    `json:"youtube_url"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (p *Project) TableName() string {
	return "projects"
}
