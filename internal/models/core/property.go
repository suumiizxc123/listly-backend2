package core

import "time"

type Property struct {
	UID                        string    `json:"uid"`
	ActionType                 string    `json:"action_type"`
	Address                    string    `json:"address"`
	AgencyUID                  string    `json:"agency_uid"`
	Area                       string    `json:"area"`
	BuildDate                  time.Time `json:"build_date"`
	City                       string    `json:"city"`
	Description                string    `json:"description"`
	District                   string    `json:"district"`
	HasBarter                  bool      `json:"has_barter"`
	HasFence                   bool      `json:"has_fence"`
	HasFurniture               bool      `json:"has_furniture"`
	HasGarage                  bool      `json:"has_garage"`
	IsActive                   bool      `json:"is_active"`
	IsCompanyLeasingAvailable  bool      `json:"is_company_leasing_available"`
	IsLeasingAvailable         bool      `json:"is_leasing_available"`
	IsPersonalLeasingAvailable bool      `json:"is_personal_leasing_available"`
	IsPublic                   bool      `json:"is_public"`
	Khoroo                     string    `json:"khoroo"`
	Lat                        float32   `json:"lat"`
	Lng                        float32   `json:"lng"`
	Layout                     string    `json:"layout"`
	Name                       string    `json:"name"`
	Price                      float32   `json:"price"`
	PricePerSquare             float32   `json:"price_per_square"`
	Status                     string    `json:"status"`
	TotalBathrooms             int64     `json:"total_bathrooms"`
	TotalBedrooms              int64     `json:"total_bedrooms"`
	TotalFloors                int64     `json:"total_floors"`
	TotalLandSquare            float32   `json:"total_land_square"`
	TotalRooms                 int64     `json:"total_rooms"`
	TotalSquare                float32   `json:"total_square"`
	Type                       string    `json:"type"`
	CreatedAt                  time.Time `json:"created_at"`
	UpdatedAt                  time.Time `json:"updated_at"`
}

func (p *Property) TableName() string {
	return "properties"
}
