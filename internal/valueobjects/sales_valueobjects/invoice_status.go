package salesvalueobjects

type InvoiceStatus string

const (
	DRAFT          InvoiceStatus = "Draft"
	DELIVERED      InvoiceStatus = "Delivered"
	UNPAID         InvoiceStatus = "Unpaid"
	PARTIALLY_PAID InvoiceStatus = "Partially Paid"
	PAID           InvoiceStatus = "Paid"
)
