package core

import "time"

type PropertyImage struct {
	UID        string `json:"uid"`
	PropertyID string `json:"property_id"`
	Image      string `json:"image"`
	Ordering   int64  `json:"ordering"`

	CreatedAt time.Time `json:"created_at"`
}

func (p *PropertyImage) TableName() string {
	return "property_images"
}
