package itemsdomain

import (
	financial "tax_calculator/engine/internal/domains/book_keeping_domain/financial_domain"
	"tax_calculator/engine/internal/valueobjects/items_valueobjects"
)

type Item struct {
	id                 int
	itemType           itemsvalueobjects.ItemType
	name               string
	code               string
	category           *Category
	sellingPrice       float32
	taxRate            *financial.TaxRate
	sellingPriceObject *itemsvalueobjects.CostObject
	buyingPriceObject  *itemsvalueobjects.CostObject
}
