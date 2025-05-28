package items

import (
	"tax_calculator/engine/internal/domains/financial"
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
