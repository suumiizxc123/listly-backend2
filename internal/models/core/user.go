package core

type Users struct {
	UID          string `json:"uid"`
	AgencyUID    string `json:"agency_uid"`
	AgencyLogo   string `json:"agency_logo"`
	Birthday     string `json:"birthday"`
	Description  string `json:"description"`
	Email        string `json:"email"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	IsSubscribed bool   `json:"is_subscribed"`
	Logo         string `json:"logo"`
	Phone        string `json:"phone"`
	UserImage    string `json:"user_image"`
	UserName     string `json:"user_name"`
}

func (u *Users) TableName() string {
	return "users"
}
