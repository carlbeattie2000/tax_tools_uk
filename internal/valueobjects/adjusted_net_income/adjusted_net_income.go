package adjustednetincome

func CalculateAdjustedNetIncome(
	totalIncomeFromAllSources float32,
	giftOfInvestmentsAndPropertyToCharity float32,
	grossGiftAidPayments float32,
	lossesAppliedToGeneralIncome float32,
	grossAnnuityPayments float32,
	qualifyingLoanInterestFromInvestments float32,
	postCessationTradeReliefs float32,
	totalPensionContributionsAllowance float32,
	totalPensionContributionsRelief float32,
) float32 {
	var totalDeductionsForAdjustedNetIncome float32 = giftOfInvestmentsAndPropertyToCharity + grossGiftAidPayments + lossesAppliedToGeneralIncome + grossAnnuityPayments + qualifyingLoanInterestFromInvestments + postCessationTradeReliefs + totalPensionContributionsAllowance + totalPensionContributionsRelief
	return totalIncomeFromAllSources - totalDeductionsForAdjustedNetIncome
}
