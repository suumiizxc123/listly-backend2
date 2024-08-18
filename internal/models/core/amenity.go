package core

import "time"

type Amenity struct {
	UID       string    `json:"uid"`
	Slug      string    `json:"slug"`
	NameMn    string    `json:"name_mn"`
	NameEn    string    `json:"name_en"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (a *Amenity) TableName() string {
	return "amenities"
}
