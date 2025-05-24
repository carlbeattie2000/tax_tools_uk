package taxliability

import (
	personalallowancedomain "tax_calculator/engine/internal/domains/personal_allowance_domain"
	utilmath32 "tax_calculator/engine/internal/utils/util_math32"
)

const (
	pppBasicRateName      = "BRT"
	pppBasicRate          = .2
	pppBasicRateThreshold = 37700
)

const (
	pppHigherRateName  = "HRT"
	pppHigherRate      = .4
	pppHigherRateLimit = 125140
)

const (
	pppAdditionalRateName  = "ART"
	pppAdditionalRate      = .45
	pppAdditionalRateLimit = -1 // No Limit
)

func CalculateTaxLiability(totalProfitFromPayPensionsProfit float32) float32 {
	var totalPayPensionsProfitTaxableIncome float32
	var personalAllowance float32

	var pppBasicRateLimit float32
	var pppBasicRateAllocatedIncome float32
	var pppBasicRateTax float32

	var pppHigherRateAllocatedIncome float32
	var pppHigherRateTax float32

	var pppAdditonalRateAllocatedIncome float32
	var pppAdditionalRateTax float32

	if totalProfitFromPayPensionsProfit <= personalallowancedomain.PersonalAllowance {
		totalPayPensionsProfitTaxableIncome = 0
	} else {
		personalAllowanceService := personalallowancedomain.NewPersonalAllowanceService()
		personalAllowance = personalAllowanceService.CalculatePersonalAllowance(totalProfitFromPayPensionsProfit)
		totalPayPensionsProfitTaxableIncome = totalProfitFromPayPensionsProfit - personalAllowance
		pppBasicRateLimit = pppBasicRateThreshold + personalAllowance
	}

	if totalPayPensionsProfitTaxableIncome <= pppBasicRateLimit {
		pppBasicRateAllocatedIncome = totalPayPensionsProfitTaxableIncome
		pppBasicRateTax = utilmath32.RoundDown(pppBasicRateAllocatedIncome*pppBasicRate, 2)

		return utilmath32.RoundDown(pppBasicRateTax, 2)
	}

	if totalPayPensionsProfitTaxableIncome <= pppHigherRateLimit {
		pppBasicRateAllocatedIncome = pppBasicRateThreshold
		pppBasicRateTax = utilmath32.RoundDown(pppBasicRateAllocatedIncome*pppBasicRate, 2)
		pppHigherRateAllocatedIncome = totalPayPensionsProfitTaxableIncome - pppBasicRateThreshold
		pppHigherRateTax = utilmath32.RoundDown(pppHigherRateAllocatedIncome*pppHigherRate, 2)
		return utilmath32.RoundDown(pppBasicRateTax+pppHigherRateTax, 2)
	}

	pppBasicRateAllocatedIncome = pppBasicRateThreshold
	pppBasicRateTax = utilmath32.RoundDown(pppBasicRateAllocatedIncome*pppBasicRate, 2)
	pppHigherRateAllocatedIncome = pppHigherRateLimit - pppBasicRateThreshold
	pppHigherRateTax = utilmath32.RoundDown(pppHigherRateAllocatedIncome*pppHigherRate, 2)
	pppAdditonalRateAllocatedIncome = totalPayPensionsProfitTaxableIncome - pppHigherRateLimit
	pppAdditionalRateTax = utilmath32.RoundDown(
		pppAdditonalRateAllocatedIncome*pppAdditionalRate,
		2,
	)

	return utilmath32.RoundDown(pppBasicRateTax+pppHigherRateTax+pppAdditionalRateTax, 2)
}
