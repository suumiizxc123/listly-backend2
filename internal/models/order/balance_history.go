package order

import (
	"kcloudb1/internal/config"
	"time"
)

type BalanceHistory struct {
	ID        int64     `json:"ID"`
	ClientID  int64     `json:"user_id"`
	MetalID   int64     `json:"metal_id"`
	Balance   float32   `json:"balance"`
	Quantity  float32   `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
}

func (b *BalanceHistory) TableName() string {
	return "one_balance_history"
}

func (b *BalanceHistory) Create() error {
	return config.DB.Create(b).Error
}

func (b *BalanceHistory) Update() error {
	return config.DB.Updates(b).Error
}

func (b *BalanceHistory) Delete() error {
	return config.DB.Delete(b).Error
}

func (b *BalanceHistory) GetByClientAndMetalID(clientID int64, metalID int64) ([]BalanceHistory, error) {
	var balances []BalanceHistory
	err := config.DB.Where("client_id = ?", clientID).Where("metal_id = ?", metalID).Find(&balances).Error
	return balances, err
}
