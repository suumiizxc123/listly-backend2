package common

import "kcloudb1/internal/config"

type Language struct {
	ID       int64  `json:"ID" gorm:"primary_key"`
	Language string `json:"language"`
}

func (c *Language) TableName() string {
	return "language"
}

func (c *Language) Create() error {
	return config.DB.Create(c).Error
}

func (c *Language) GetList() ([]Language, error) {
	var languages []Language

	if err := config.DB.Find(&languages).Error; err != nil {
		return nil, err
	}

	return languages, nil
}

func (c *Language) Update() error {
	return config.DB.Updates(c).Error
}

func (c *Language) Delete() error {
	return config.DB.Delete(c).Error
}
