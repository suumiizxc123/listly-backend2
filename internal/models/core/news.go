package core

import "time"

type News struct {
	UID              string    `json:"uid"`
	AgencyUID        string    `json:"agency_uid"`
	Category         string    `json:"category"`
	Content          string    `json:"content"`
	Image            string    `json:"image"`
	ShortDescription string    `json:"short_description"`
	Title            string    `json:"title"`
	CreatedAt        time.Time `json:"created_at"`
}

func (n *News) TableName() string {
	return "news"
}
