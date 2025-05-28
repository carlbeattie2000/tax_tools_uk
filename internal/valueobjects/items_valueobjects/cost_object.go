package itemsvalueobjects

import "tax_calculator/engine/internal/domains/financial"

type CostObject struct {
	price       float32
	account     int
	taxRate     *financial.TaxRate
	description string
}
