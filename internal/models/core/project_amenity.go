package core

type ProjectAmenity struct {
	UID string `json:"uid"`

	ProjectUID string `json:"project_uid"`

	AmenityUID string `json:"amenity_uid"`

	IsActive bool `json:"is_active"`
}

func (c *ProjectAmenity) TableName() string {
	return "project_amenities"
}
