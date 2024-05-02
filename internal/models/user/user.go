package user

import (
	"kcloudb1/internal/config"
	"time"
)

type User struct {
	ID        int64     `json:"ID" gorm:"primary_key"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	KaraokeID int64     `json:"karaoke_id"`
	RoleID    int64     `json:"role_id"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	IsActive  int64     `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}

func (c *User) TableName() string {
	return "user"
}

func (c *User) Create() error {
	return config.DB.Create(c).Error
}

func (c *User) Update() error {
	return config.DB.Updates(c).Error
}

func (c *User) Delete() error {
	return config.DB.Delete(c).Error
}

func (c *User) Get() error {
	return config.DB.Where("id = ?", c.ID).First(c).Error
}

func (c *User) GetList() ([]User, error) {
	var users []User

	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
