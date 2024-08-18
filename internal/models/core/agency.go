package core

import "time"

type Agency struct {
	UID             string `json:"uid"`
	Address         string `json:"address"`
	AgencyWhiteLogo string `json:"agency_white_logo"`
	Cover           string `json:"cover"`

	Description string    `json:"description"`
	Email       string    `json:"email"`
	Franchise   string    `json:"franchise"`
	IsActive    bool      `json:"is_active"`
	Logo        string    `json:"logo"`
	Name        string    `json:"name"`
	Phone       string    `json:"phone"`
	Type        string    `json:"type"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (a *Agency) TableName() string {
	return "agencies"
}
