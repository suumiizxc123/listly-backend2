package user

import "time"

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

type SysUser struct {
	ID        int64     `json:"ID" gorm:"primary_key"`
	RoleID    int64     `json:"role_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	IsActive  int64     `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}

func (c *SysUser) TableName() string {
	return "sys_user"
}

type ServiceLog struct {
	ID        int64     `json:"ID" gorm:"primary_key"`
	KaraokeID int64     `json:"karaoke_id"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (c *ServiceLog) TableName() string {
	return "service_log"
}
