package order

import (
	"kcloudb1/internal/config"
	"kcloudb1/internal/models/client"
	"kcloudb1/internal/models/metal"
	"time"
)

type Order struct {
	ID          int64     `json:"id"`
	ClientID    int64     `json:"client_id"`
	Amount      float32   `json:"amount"`
	Price       float32   `json:"price"`
	MetalID     int64     `json:"metal_id"`
	Quantity    float32   `json:"quantity"`
	Status      string    `json:"status"`
	AdminStatus string    `json:"admin_status"`
	CreatedAt   time.Time `json:"created_at"`
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
	Quantity float32 `json:"quantity"`
	MetalID  int64   `json:"metal_id"`
}

type OrderExtend struct {
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
	CreatedAt   time.Time     `json:"created_at"`
}

func (o *OrderExtend) TableName() string {
	return "one_order"
}
