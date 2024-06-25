package admin

import "kcloudb1/internal/config"

type Admin struct {
	ID       int64  `json:"ID"`
	Name     string `json:"name"`
	Password string `json:"password"`
	IsActive int64  `json:"is_active"`
}

func (Admin) TableName() string {
	return "one_admin"
}

func (a *Admin) Get(id any) error {
	return config.DB.First(a, id).Error
}

func (a *Admin) GetByName(name string) error {
	return config.DB.Where("name = ?", name).First(a).Error
}

type AdminOutput struct {
	ID       int64  `json:"ID"`
	Name     string `json:"name"`
	IsActive int64  `json:"is_active"`
	Token    string `json:"token"`
}
