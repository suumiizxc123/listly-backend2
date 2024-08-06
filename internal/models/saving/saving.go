package saving

import (
	"kcloudb1/internal/config"
	"time"
)

type SavingOrder struct {
	ID          int64     `json:"id"`
	ClientID    int64     `json:"client_id"`
	MetalID     int64     `json:"metal_id"`
	Quantity    float32   `json:"quantity"`
	Amount      float32   `json:"amount"`
	Price       float32   `json:"price"`
	Type        string    `json:"type"`
	Status      string    `json:"status"`
	AdminStatus string    `json:"admin_status"`
	CreatedAt   time.Time `json:"created_at"`
}

func (o *SavingOrder) TableName() string {
	return "one_saving_order"
}

func (o *SavingOrder) Create() error {
	return config.DB.Create(o).Error
}

func (o *SavingOrder) Update() error {
	return config.DB.Updates(o).Error
}

func (o *SavingOrder) Delete() error {
	return config.DB.Delete(o).Error
}

func (o *SavingOrder) Get() error {
	return config.DB.First(o).Error
}

func (o *SavingOrder) GetAll() ([]SavingOrder, error) {
	var orders []SavingOrder
	err := config.DB.Find(&orders).Error
	return orders, err
}

func (o *SavingOrder) GetByClientID(clientID int64) ([]SavingOrder, error) {
	var orders []SavingOrder
	err := config.DB.Where("client_id = ?", clientID).Find(&orders).Error
	return orders, err
}

func (o *SavingOrder) GetByMetalID(metalID int64) ([]SavingOrder, error) {
	var orders []SavingOrder
	err := config.DB.Where("metal_id = ?", metalID).Find(&orders).Error
	return orders, err
}

func (o *SavingOrder) GetByClientIDAndMetalID(clientID, metalID int64) ([]SavingOrder, error) {
	var orders []SavingOrder
	err := config.DB.Where("client_id = ?", clientID).Where("metal_id = ?", metalID).Find(&orders).Error
	return orders, err
}

type SavingOrderPayment struct {
	ID                  int64   `json:"id"`
	SavingOrderID       int64   `json:"saving_order_id"`
	InvoiceCode         string  `json:"invoice_code"`
	SenderInvoiceNo     string  `json:"sender_invoice_no"`
	InvoiceReceiverCode string  `json:"invoice_receiver_code"`
	InvoiceDescription  string  `json:"invoice_description"`
	SenderBranchCode    string  `json:"sender_branch_code"`
	CallbackURL         string  `json:"callback_url"`
	Amount              float32 `json:"amount"`
	InvoiceID           string  `json:"invoice_id"`
	QRText              string  `json:"qr_text"`
	QRImage             string  `json:"qr_image"`
	QPayShortUrl        string  `json:"qPay_shortUrl"`
	Urls                string  `json:"urls"`
}

func (o *SavingOrderPayment) TableName() string {
	return "one_saving_order_payment"
}

func (o *SavingOrderPayment) Create() error {
	return config.DB.Create(o).Error
}

func (o *SavingOrderPayment) Update() error {
	return config.DB.Updates(o).Error
}

func (o *SavingOrderPayment) Delete() error {
	return config.DB.Delete(o).Error
}

func (o *SavingOrderPayment) Get() error {
	return config.DB.First(o).Error
}

func (o *SavingOrderPayment) GetAll() ([]SavingOrderPayment, error) {
	var orders []SavingOrderPayment
	err := config.DB.Find(&orders).Error
	return orders, err
}

type CreateSavingOrderInput struct {
	ClientID int64   `json:"client_id"`
	MetalID  int64   `json:"metal_id"`
	Quantity float32 `json:"quantity"`
}
