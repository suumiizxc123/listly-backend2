package order

import (
	"kcloudb1/internal/config"
	"kcloudb1/internal/models/client"
	"kcloudb1/internal/models/metal"
	"time"

	"gorm.io/gorm/clause"
)

type WithDraw struct {
	ID          int64     `json:"id"`
	ClientID    int64     `json:"client_id"`
	Amount      float32   `json:"amount"`
	Price       float32   `json:"price"`
	MetalID     int64     `json:"metal_id"`
	Quantity    float32   `json:"quantity"`
	Status      string    `json:"status"`
	AdminStatus string    `json:"admin_status"`
	Type        string    `json:"type"`
	CreatedAt   time.Time `json:"created_at"`
}

func (w *WithDraw) TableName() string {
	return "one_order"
}

func (w *WithDraw) Create() error {
	return config.DB.Create(w).Error
}

func (w *WithDraw) Update() error {
	return config.DB.Updates(w).Error
}

func (w *WithDraw) Delete() error {
	return config.DB.Delete(w).Error
}

func (w *WithDraw) Get() error {
	return config.DB.Where("id = ?", w.ID).First(w).Error
}

func (w *WithDraw) GetAll() ([]WithDraw, error) {
	var wd []WithDraw
	return wd, config.DB.Find(&wd).Error
}

func (w *WithDraw) GetByClientID(clientID int64) ([]WithDraw, error) {
	var wd []WithDraw
	return wd, config.DB.Order("created_at desc").Where("client_id = ?", clientID).Find(&wd).Error
}

type CreateWithDrawInput struct {
	ClientID int64   `json:"client_id"`
	Quantity float32 `json:"quantity"`
	MetalID  int64   `json:"metal_id"`
}

type WithDrawExtend struct {
	ID          int64         `json:"id"`
	ClientID    int64         `json:"user_id"`
	Client      client.Client `json:"client" gorm:"foreignKey:ID; references:ClientID"`
	Amount      float32       `json:"amount"`
	Price       float32       `json:"price"`
	MetalID     int64         `json:"metal_id"`
	Metal       metal.Metal   `json:"metal" gorm:"foreignKey:ID; references:MetalID"`
	Quantity    float32       `json:"quantity"`
	Status      string        `json:"status"`
	AdminStatus string        `json:"admin_status"`
	Type        string        `json:"type"`
	CreatedAt   time.Time     `json:"created_at"`
}

func (w *WithDrawExtend) TableName() string {
	return "one_order"
}

func (w *WithDrawExtend) Get() error {
	return config.DB.Where("id = ?", w.ID).Preload(clause.Associations).First(w).Error
}

func (w *WithDrawExtend) GetListByClientID(clientID int64) ([]WithDrawExtend, error) {
	var wd []WithDrawExtend
	return wd, config.DB.Where("client_id = ?", clientID).Order("created_at desc").Preload(clause.Associations).Find(&wd).Error
}

func (w *WithDrawExtend) GetAll() ([]WithDrawExtend, error) {
	var wd []WithDrawExtend
	return wd, config.DB.Order("created_at desc").Preload(clause.Associations).Find(&wd).Error
}
