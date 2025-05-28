package selfemploymentdomain

type SelfEmploymentService struct{}

func NewSelfEmploymentService() SelfEmploymentService {
	return SelfEmploymentService{}
}

func (selfEmploymentService *SelfEmploymentService) CalculateTotalTaxableSelfEmploymentProfit(
	selfEmploymentIncome *SelfEmploymentIncomeEntity,
	selfEmploymentExpenses *SelfEmploymentExpensesEntity,
	selfEmploymentAdditions *SelfEmploymentAdditionsEntity,
	selfEmploymentDeductions *SelfEmploymentDeductionsEntity,
	selfEmploymentAccountingAdjustments *SelfEmploymentAccountingAdjustmentsEntity,
) TotalTaxableSelfEmploymentProfit {
	totalSelfEmploymentIncome := selfEmploymentIncome.totalSelfEmploymentIncome
	totalSelfEmploymentExpenses := selfEmploymentExpenses.totalSelfEmploymentExpenses
	totalSelfEmploymentAdditions := selfEmploymentAdditions.totalSelfEmploymentAdditons
	totalSelfEmploymentDeductions := selfEmploymentDeductions.totalSelfEmploymentDeductions
	totalSelfEmploymentAccountingAdjustments := selfEmploymentAccountingAdjustments.totalSelfEmploymentAccountingAdjustments

	var netProfitFromSelfEmployment float32
	var netLossFromSelfEmployment float32
	var adjustedProfitOrLossFromSelfEmployment float32
	var taxableProfitFromSelfEmployment float32
	var taxableLossFromSelfEmployment float32

	if totalSelfEmploymentIncome >= totalSelfEmploymentExpenses {
		netProfitFromSelfEmployment = totalSelfEmploymentIncome - totalSelfEmploymentExpenses
	} else {
		netLossFromSelfEmployment = totalSelfEmploymentIncome - totalSelfEmploymentExpenses
	}

	adjustedProfitOrLossFromSelfEmployment = netProfitFromSelfEmployment + netLossFromSelfEmployment + totalSelfEmploymentAdditions + totalSelfEmploymentDeductions + totalSelfEmploymentAccountingAdjustments

	if adjustedProfitOrLossFromSelfEmployment >= 0 {
		taxableProfitFromSelfEmployment = adjustedProfitOrLossFromSelfEmployment
	} else {
		taxableLossFromSelfEmployment = adjustedProfitOrLossFromSelfEmployment
	}

	return NewSelfEmploymentTotalTaxableSelfEmploymentProfit(
		taxableProfitFromSelfEmployment,
		taxableLossFromSelfEmployment,
	)
}
