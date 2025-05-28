package financialdomain

type CurrencyRepository interface {
	FindById(id int) (*Currency, error)
	Create(currency *Currency) (*Currency, error)
	Delete(currency *Currency) error
}
