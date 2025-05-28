package financial

import financialvalueobjects "tax_calculator/engine/internal/valueobjects/financial_valueobjects"

type Account struct {
	name        string
	code        string
	accountType financialvalueobjects.AccountType
	normal      financialvalueobjects.AccountNormal
	currency    *Currency
	bankBalance float32
	balance     float32
}
