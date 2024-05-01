package artist

import "time"

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

func (ArtistProfile) TableName() string {
	return "artist_profile"
}
