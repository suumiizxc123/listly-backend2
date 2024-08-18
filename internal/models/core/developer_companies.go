package core

import "time"

type DeveloperCompany struct {
	UID       string    `json:"uid"`
	Address   string    `json:"address"`
	Email     string    `json:"email"`
	Logo      string    `json:"logo"`
	Name      string    `json:"name"`
	Note      string    `json:"note"`
	Phone     string    `json:"phone"`
	Website   string    `json:"website"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (d *DeveloperCompany) TableName() string {
	return "developer_companies"
}
