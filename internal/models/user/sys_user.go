package user

import (
	"encoding/json"
	"kcloudb1/internal/config"
	"time"
)

type SysUser struct {
	ID        int64     `json:"ID" gorm:"primary_key"`
	UID       string    `json:"uid"`
	RoleID    int64     `json:"role_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	IsActive  int64     `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	Token     string    `gorm:"-" json:"token, omitempty"`
}

// example json body of SysUser
// {
// 	"role_id": 1,
// 	"first_name": "John",
// 	"last_name": "Doe",
// 	"email": "nI8o6@example.com",
// 	"phone": "1234567890",
// 	"password": "password",
// }

func (c *SysUser) TableName() string {
	return "sys_user"
}

func (c *SysUser) Create() error {
	return config.DB.Create(c).Error
}

func (c *SysUser) Update() error {
	return config.DB.Updates(c).Error
}

func (c *SysUser) Delete() error {
	return config.DB.Delete(c).Error
}

func (c *SysUser) Get() error {
	return config.DB.Where("id = ?", c.ID).First(c).Error
}

func (c *SysUser) CheckEmailAndPhoneNotExist() bool {

	var user SysUser

	if err := config.DB.Where("email = ? ", c.Email).First(&user).Error; err == nil {
		return false
	}

	if err := config.DB.Where("phone = ? ", c.Phone).First(&user).Error; err == nil {
		return false
	}

	return true
}

func (c *SysUser) GetList() ([]SysUser, error) {
	var users []SysUser

	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (c *SysUser) Filter() ([]SysUser, error) {

	var users []SysUser
	qry := config.DB

	if c.RoleID != 0 {
		qry = qry.Where("role_id = ?", c.RoleID)
	}

	if c.Email != "" {
		qry = qry.Where("email = ?", c.Email)
	}

	if c.Phone != "" {
		qry = qry.Where("phone = ?", c.Phone)
	}

	if err := qry.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

type SysUserLoginInput struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func (c *SysUser) Login(phone string, password string) (SysUser, error) {

	var user SysUser

	if err := config.DB.Where("phone = ? AND password = ?", phone, password).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (c *SysUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(*c)

}
