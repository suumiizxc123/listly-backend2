package admin

import (
	"kcloudb1/internal/config"
	"time"
)

type Slider struct {
	ID          int64     `json:"id"`
	Image       string    `json:"image"`
	Title       string    `json:"title"`
	Subtitle    string    `json:"subtitle"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

func (c *Slider) TableName() string {
	return "one_slider"
}

func (c *Slider) Create() error {
	return config.DB.Create(c).Error
}

func (c *Slider) Get() error {
	return config.DB.First(c).Error
}

func (c *Slider) Update() error {
	return config.DB.Updates(c).Error
}

func (c *Slider) Delete() error {
	return config.DB.Delete(c).Error
}

func (c *Slider) GetAll() ([]Slider, error) {
	var sliders []Slider
	err := config.DB.Order("created_at desc").Find(&sliders).Error
	return sliders, err
}
