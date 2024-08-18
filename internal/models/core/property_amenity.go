package core

type PropertyAmenity struct {
	UID         string `json:"uid"`
	PropertyUID string `json:"property_uid"`
	AmenityUID  string `json:"amenity_uid"`
}

func (p *PropertyAmenity) TableName() string {
	return "property_amenities"
}
