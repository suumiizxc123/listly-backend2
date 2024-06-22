package order

import (
	"kcloudb1/internal/config"
	"time"
)

type Order struct {
	ID        int64     `json:"ID"`
	ClientID  int64     `json:"user_id"`
	Amount    float32   `json:"amount"`
	Price     float32   `json:"price"`
	MetalID   int64     `json:"metal_id"`
	Quantity  float32   `json:"quantity"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

func (o *Order) TableName() string {
	return "one_order"
}

func (o *Order) Create() error {
	return config.DB.Create(o).Error
}

func (o *Order) Update() error {
	return config.DB.Updates(o).Error
}

func (o *Order) Delete() error {
	return config.DB.Delete(o).Error
}

func (o *Order) Get() error {
	return config.DB.First(o).Error
}

func (o *Order) GetAll() ([]Order, error) {
	var orders []Order
	err := config.DB.Find(&orders).Error
	return orders, err
}

func (o *Order) GetByClientID(clientID int64) ([]Order, error) {
	var orders []Order

	err := config.DB.Order("created_at desc").Where("client_id = ?", clientID).Find(&orders).Error

	return orders, err
}

func (o *Order) GetSenderInvoiceNo(senderInvoiceNo string) error {
	return config.DB.Where("sender_invoice_no = ?", senderInvoiceNo).First(o).Error
}

type CreateOrderInput struct {
	ClientID int64   `json:"user_id"`
	Amount   float32 `json:"amount"`
	MetalID  int64   `json:"metal_id"`
}
