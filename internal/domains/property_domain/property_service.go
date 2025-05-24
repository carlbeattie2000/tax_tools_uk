package propertydomain

type PropertyTaxService struct{}

func NewPropertyTaxService() PropertyTaxService {
	return PropertyTaxService{}
}

func (propertyTaxService *PropertyTaxService) CalculateTotalTaxablePropertyProfit(
	propertyIncome *PropertyIncomeEntity,
	propertyExpenses *PropertyExpensesEntity,
	propertyAdditions *PropertyAdditionsEntity,
	propertyDeductions *PropertyDeductionsEntity,
) TotalTaxablePropertyProfit {
	var totalIncomeFromUkPropertyOther float32 = propertyIncome.totalIncomeFromUkPropertyOther
	var totalExpensesFromUkPropertyOther float32 = propertyExpenses.totalExpensesFromUkPropertyOther
	var totalAdditionsFromUkPropertyOther float32 = propertyAdditions.totalAdditionsFromUkPropertyOther
	var totalDeductionsFromUkPropertyOther float32 = propertyDeductions.totalDeductionsFromUkPropertyOther

	var netProfitFromUkPropertyOther float32
	var netLossFromUkPropertyOther float32
	var adjustedProfitOrLossFromUkPropertyOther float32
	var taxableProfitFromUkPropertyOther float32
	var taxableLossFromUkPropertyOther float32

	if totalIncomeFromUkPropertyOther >= totalExpensesFromUkPropertyOther {
		netProfitFromUkPropertyOther = totalIncomeFromUkPropertyOther - totalExpensesFromUkPropertyOther
	} else {
		netLossFromUkPropertyOther = totalIncomeFromUkPropertyOther - totalExpensesFromUkPropertyOther
	}

	adjustedProfitOrLossFromUkPropertyOther = netProfitFromUkPropertyOther + netLossFromUkPropertyOther + totalAdditionsFromUkPropertyOther - totalDeductionsFromUkPropertyOther
	if adjustedProfitOrLossFromUkPropertyOther >= 0 {
		taxableProfitFromUkPropertyOther = adjustedProfitOrLossFromUkPropertyOther
	} else {
		taxableLossFromUkPropertyOther = adjustedProfitOrLossFromUkPropertyOther
	}

	return TotalTaxablePropertyProfit{
		taxableProfitFromUkPropertyOther,
		taxableLossFromUkPropertyOther,
	}
}
