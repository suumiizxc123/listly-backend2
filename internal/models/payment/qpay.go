package payment

type QPayBankURL struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
	Link        string `json:"link"`
}
type QPayInvoiceResponse struct {
	InvoiceID    string        `json:"invoice_id"`
	QRText       string        `json:"qr_text"`
	QRImage      string        `json:"qr_image"`
	QPayShortUrl string        `json:"qPay_shortUrl"`
	Urls         []QPayBankURL `json:"urls"`
}

type QPayInvoiceInput struct {
	InvoiceCode         string  `json:"invoice_code"`
	SenderInvoiceNo     string  `json:"sender_invoice_no"`
	InvoiceReceiverCode string  `json:"invoice_receiver_code"`
	InvoiceDescription  string  `json:"invoice_description"`
	SenderBranchCode    string  `json:"sender_branch_code"`
	Amount              float32 `json:"amount"`
	CallbackURL         string  `json:"callback_url"`
}
