package moscap

import (
	"kcloudb1/internal/config"
	"time"
)

type MosCapUser struct {
	ID        int64     `json:"ID" gorm:"primary_key"`
	UID       string    `json:"uid"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	IsActive  int64     `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	Token     string    `gorm:"-" json:"token, omitempty"`
}

func (c *MosCapUser) TableName() string {
	return "moscap_user"
}

func (c *MosCapUser) Create() error {
	return config.DB.Create(c).Error
}

func (c *MosCapUser) Update() error {
	return config.DB.Updates(c).Error
}

func (c *MosCapUser) Delete() error {
	return config.DB.Delete(c).Error
}

func (c *MosCapUser) Get() error {
	return config.DB.Where("id = ?", c.ID).First(c).Error
}

func (c *MosCapUser) GetList() ([]MosCapUser, error) {
	var orgs []MosCapUser

	if err := config.DB.Find(&orgs).Error; err != nil {
		return nil, err
	}

	return orgs, nil
}
