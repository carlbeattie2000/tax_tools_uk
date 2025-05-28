package invoicedomain

type InvoiceRepository interface {
	FindById(id int) (*Invoice, error)
	Create(invoice *Invoice) (*Invoice, error)
	Delete(invoice *Invoice) error
}
