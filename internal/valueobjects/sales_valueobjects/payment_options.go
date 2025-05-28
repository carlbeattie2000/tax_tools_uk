package salesvalueobjects

type PaymentOptions int

const (
	CASH PaymentOptions = iota
	CARD
	CHECK
	VOUCHER
)
