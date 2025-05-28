package itemsvalueobjects

import "tax_calculator/engine/internal/domains/book_keeping_domain/financial_domain"

type CostObject struct {
	price       float32
	account     int
	taxRate     *financialdomain.TaxRate
	description string
}
