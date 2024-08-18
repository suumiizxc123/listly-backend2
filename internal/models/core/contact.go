package core

import "time"

type Contact struct {
	UID           string    `json:"uid"`
	Address       string    `json:"address"`
	AgencyUID     string    `json:"agency_uid"`
	CallsNum      int64     `json:"calls_num"`
	Category      string    `json:"category"`
	Email         string    `json:"email"`
	EmailNum      int64     `json:"email_num"`
	FacebookURL   string    `json:"facebook_url"`
	Image         string    `json:"image"`
	InstagramName string    `json:"instagram_name"`
	MetDate       string    `json:"met_date"`
	Name          string    `json:"name"`
	Note          string    `json:"note"`
	Phone         string    `json:"phone"`
	Rank          string    `json:"rank"`
	Status        string    `json:"status"`
	TextNum       int64     `json:"text_num"`
	Type          string    `json:"type"`
	UserUID       string    `json:"user_uid"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (c *Contact) TableName() string {
	return "contacts"
}
