package admin

import (
	"kcloudb1/internal/config"
	"time"
)

type Image struct {
	ID        int64     `json:"id"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}

func (i *Image) TableName() string {
	return "one_image"
}

func (i *Image) Create() error {
	return config.DB.Create(i).Error
}

func (i *Image) Get() error {
	return config.DB.Where("id = ?", i.ID).First(i).Error
}

func (i *Image) Update() error {
	return config.DB.Save(i).Error
}

func (i *Image) Delete() error {
	return config.DB.Delete(i).Error
}
