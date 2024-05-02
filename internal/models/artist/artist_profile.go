package artist

import (
	"kcloudb1/internal/config"
	"time"
)

type ArtistProfile struct {
	ID            int64     `json:"ID" gorm:"primary_key"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	Description   string    `json:"description"`
	Gender        int64     `json:"gender"`
	Email         string    `json:"email"`
	Phone         string    `json:"phone"`
	BankID        int64     `json:"bank_id"`
	AccountNumber string    `json:"account_number"`
	IsActive      int64     `json:"is_active"`
	CreatedAt     time.Time `json:"created_at"`
}

func (c *ArtistProfile) TableName() string {
	return "artist_profile"
}

func (c *ArtistProfile) Create() error {
	return config.DB.Create(c).Error
}

func (c *ArtistProfile) Get() error {
	return config.DB.Where("id = ?", c.ID).First(c).Error
}

func (c *ArtistProfile) GetList() ([]ArtistProfile, error) {
	var list []ArtistProfile
	return list, config.DB.Find(&list).Error
}

func (c *ArtistProfile) Update() error {
	return config.DB.Updates(c).Error
}

func (c *ArtistProfile) Delete() error {
	return config.DB.Delete(c).Error
}
