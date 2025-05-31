package invoicedomain

type PaymentReceived struct {
	id        int
	invoiceId int
	amount    float32
}
