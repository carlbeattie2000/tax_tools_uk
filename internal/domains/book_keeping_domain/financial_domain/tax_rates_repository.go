package financialdomain

type TaxRateRepository interface {
	FindById(id int) (*TaxRate, error)
	Create(taxRate *TaxRate) (*TaxRate, error)
	Delete(taxRate *TaxRate) error
}
