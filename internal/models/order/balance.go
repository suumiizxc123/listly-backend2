package order

import (
	"kcloudb1/internal/config"
	"time"
)

type Balance struct {
	ID        int64     `json:"id"`
	ClientID  int64     `json:"client_id"`
	MetalID   int64     `json:"metal_id"`
	Quantity  float32   `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
}

func (b *Balance) TableName() string {
	return "one_balance"
}

func (b *Balance) Create() error {
	return config.DB.Create(b).Error
}

func (b *Balance) Update() error {
	return config.DB.Updates(b).Error
}

func (b *Balance) Delete() error {
	return config.DB.Delete(b).Error
}

func (b *Balance) Get() error {
	return config.DB.Where("client_id = ?", b.ClientID).First(b).Error
}

func (b *Balance) GetAll() ([]Balance, error) {
	var balances []Balance
	err := config.DB.Find(&balances).Error
	return balances, err
}

func (b *Balance) GetByClientAndMetalID(clientID any, metalID int64) error {
	err := config.DB.Where("client_id = ?", clientID).Where("metal_id = ?", metalID).Find(&b).Error
	return err
}

type BalanceResponse struct {
	ID        int64     `json:"id"`
	ClientID  int64     `json:"user_id"`
	MetalID   int64     `json:"metal_id"`
	Quantity  float32   `json:"quantity"`
	Balance   float32   `json:"balance"`
	Changes   float32   `json:"changes"`
	CreatedAt time.Time `json:"created_at"`
}
