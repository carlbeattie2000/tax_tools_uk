package incomesummarydomain

import (
	propertydomain "tax_calculator/engine/internal/domains/property_domain"
	selfemploymentdomain "tax_calculator/engine/internal/domains/self_employment_domain"
)

type IncomeSummaryEntity struct {
	totalUntaxedInterest                                        float32
	totalGrossUkInterest                                        float32
	foreignSavingsInterest                                      float32
	totalGrossSecurities                                        float32
	untaxedUKGainsIncome                                        float32
	untaxedForeignGainsIncome                                   float32
	taxableProfitFromSelfEmployment                             []selfemploymentdomain.TotalTaxableSelfEmploymentProfit
	taxableProfitFromUkPropertyOther                            propertydomain.TotalTaxablePropertyProfit
	taxableProfitFromUkPropertyFhl                              float32
	taxableProfitFromForeignPropertyOther                       float32
	taxableProfitFromEeaPropertyFhl                             float32
	totalOccupationalPensionIncome                              float32
	totalEmploymentIncomePlusBenefitsInKindMinusExpenses        float32
	taxableProfitFromShareSchemes                               float32
	totalStateBenefitsIncomeExcludingStatePensionLumpSumBenefit float32
	totalOverseasIncomeAndGains                                 float32
	otherIncomesWhileAbroad                                     float32
	foreignPensionIncome                                        float32
	chargeableForeignBenefitsAndGifts                           float32
	postCessationTradeReceipts                                  float32
	totalDividendIncomeForUkOtherAndForeign                     float32
	totalProfitFromTaxedUkGains                                 float32
	totalProfitFromTaxedForeignGains                            float32
	totalEmploymentLumpSumsNotLiableForPPP                      float32

	totalProfitFromSelfEmployment    float32
	totalSavingsIncome               float32
	totalProfitFromProperty          float32
	totalEmploymentIncome            float32
	totalProfitFromPayPensionsProfit float32
	totalIncomeFromAllSources        float32
}

func (incomeSummaryEntity *IncomeSummaryEntity) TotalIncomeFromAllSources() float32 {
	return incomeSummaryEntity.totalIncomeFromAllSources
}

func (incomeSummaryEntity *IncomeSummaryEntity) TotalProfitFromPayPensionsProfit() float32 {
	return incomeSummaryEntity.totalProfitFromPayPensionsProfit
}

func NewIncomeSummaryEntity(
	totalUntaxedInterest float32,
	totalGrossUkInterest float32,
	foreignSavingsInterest float32,
	totalGrossSecurities float32,
	untaxedUKGainsIncome float32,
	untaxedForeignGainsIncome float32,
	taxableProfitFromSelfEmployment []selfemploymentdomain.TotalTaxableSelfEmploymentProfit,
	taxableProfitFromUkPropertyOther propertydomain.TotalTaxablePropertyProfit,
	taxableProfitFromUkPropertyFhl float32,
	taxableProfitFromForeignPropertyOther float32,
	taxableProfitFromEeaPropertyFhl float32,
	totalOccupationalPensionIncome float32,
	totalEmploymentIncomePlusBenefitsInKindMinusExpenses float32,
	taxableProfitFromShareSchemes float32,
	totalStateBenefitsIncomeExcludingStatePensionLumpSumBenefit float32,
	totalOverseasIncomeAndGains float32,
	otherIncomesWhileAbroad float32,
	foreignPensionIncome float32,
	chargeableForeignBenefitsAndGifts float32,
	postCessationTradeReceipts float32,
	totalDividendIncomeForUkOtherAndForeign float32,
	totalProfitFromTaxedUkGains float32,
	totalProfitFromTaxedForeignGains float32,
	totalEmploymentLumpSumsNotLiableForPPP float32,
) IncomeSummaryEntity {
	var totalProfitFromSelfEmployment float32
	var totalSavingsIncome float32
	var totalProfitFromProperty float32
	var totalEmploymentIncome float32
	var totalProfitFromPayPensionsProfit float32
	var totalIncomeFromAllSources float32

	for _, taxableprotaxableProfitFromSelfEmploymentBlock := range taxableProfitFromSelfEmployment {
		totalProfitFromSelfEmployment += taxableprotaxableProfitFromSelfEmploymentBlock.TaxableProfitFromSelfEmployment
	}

	totalSavingsIncome = totalUntaxedInterest + totalGrossUkInterest + foreignSavingsInterest + totalGrossSecurities + untaxedUKGainsIncome + untaxedForeignGainsIncome

	totalProfitFromProperty = taxableProfitFromUkPropertyOther.TaxableProfitFromUkPropertyOther + taxableProfitFromUkPropertyFhl + taxableProfitFromForeignPropertyOther + taxableProfitFromEeaPropertyFhl

	totalEmploymentIncome = totalOccupationalPensionIncome + totalEmploymentIncomePlusBenefitsInKindMinusExpenses

	totalProfitFromPayPensionsProfit = totalProfitFromProperty + totalProfitFromSelfEmployment + taxableProfitFromShareSchemes + totalStateBenefitsIncomeExcludingStatePensionLumpSumBenefit + totalEmploymentIncome + totalOverseasIncomeAndGains + otherIncomesWhileAbroad + foreignPensionIncome + chargeableForeignBenefitsAndGifts + postCessationTradeReceipts

	totalIncomeFromAllSources = totalDividendIncomeForUkOtherAndForeign + totalSavingsIncome + totalProfitFromPayPensionsProfit + totalProfitFromTaxedUkGains + totalProfitFromTaxedForeignGains + totalEmploymentLumpSumsNotLiableForPPP

	return IncomeSummaryEntity{
		totalUntaxedInterest,
		totalGrossUkInterest,
		foreignSavingsInterest,
		totalGrossSecurities,
		untaxedUKGainsIncome,
		untaxedForeignGainsIncome,
		taxableProfitFromSelfEmployment,
		taxableProfitFromUkPropertyOther,
		taxableProfitFromUkPropertyFhl,
		taxableProfitFromForeignPropertyOther,
		taxableProfitFromEeaPropertyFhl,
		totalOccupationalPensionIncome,
		totalEmploymentIncomePlusBenefitsInKindMinusExpenses,
		taxableProfitFromShareSchemes,
		totalStateBenefitsIncomeExcludingStatePensionLumpSumBenefit,
		totalOverseasIncomeAndGains,
		otherIncomesWhileAbroad,
		foreignPensionIncome,
		chargeableForeignBenefitsAndGifts,
		postCessationTradeReceipts,
		totalDividendIncomeForUkOtherAndForeign,
		totalProfitFromTaxedUkGains,
		totalProfitFromTaxedForeignGains,
		totalEmploymentLumpSumsNotLiableForPPP,
		totalProfitFromSelfEmployment,
		totalSavingsIncome,
		totalProfitFromProperty,
		totalEmploymentIncome,
		totalProfitFromPayPensionsProfit,
		totalIncomeFromAllSources,
	}
}
