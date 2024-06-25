package order

import "kcloudb1/internal/config"

type OrderPayment struct {
	ID                  int64   `json:"id"`
	OrderID             int64   `json:"order_id"`
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
}

func (q *OrderPayment) TableName() string {
	return "one_order_payment"
}

func (q *OrderPayment) Create() error {
	return config.DB.Create(q).Error
}

func (q *OrderPayment) Update() error {
	return config.DB.Updates(q).Error
}

func (q *OrderPayment) Delete() error {
	return config.DB.Delete(q).Error
}

func (q *OrderPayment) Get() error {
	return config.DB.Where("id = ?", q.ID).First(q).Error
}

func (q *OrderPayment) GetAll() ([]OrderPayment, error) {
	var orderPayments []OrderPayment
	return orderPayments, config.DB.Find(&orderPayments).Error
}
