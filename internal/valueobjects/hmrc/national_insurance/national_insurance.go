package nationalinsurance

import (
	utilmath32 "tax_calculator/engine/internal/utils/util_math32"

	"github.com/chewxy/math32"
)

// TODO: Refine and make this modular and more readable

type NiBand struct {
	name string
	rate float32
	from float32
	to   float32
}

const (
	smallProfitsThreshold = float32(6725)
	lowerProfitsLimit     = float32(12570)
	upperProfitsLimit     = float32(50270)
)

var (
	niLowerBand  = &NiBand{name: "ZRT", rate: 0.00, from: 0, to: lowerProfitsLimit}
	niBasicBand  = &NiBand{name: "BRT", rate: 0.06, from: lowerProfitsLimit, to: upperProfitsLimit}
	niHigherBand = &NiBand{name: "HRT", rate: 0.02, from: upperProfitsLimit, to: math32.Inf(1)}
)

func CalculateNationalInsurance(totalSelfEmploymentTaxableProfit float32) float32 {
	var totalClass4NIC float32
	var remainingIncome float32

	var niZeroRateNIC float32
	var niBasicRateAllocatedIncome float32
	var niBasicRateNIC float32
	var niHigherRateAllocatedIncome float32
	var niHigherRateNIC float32

	totalSelfEmploymentTaxableProfit = utilmath32.RoundDown(totalSelfEmploymentTaxableProfit, 2)

	if totalSelfEmploymentTaxableProfit < lowerProfitsLimit {
		return 0.0
	} else if totalSelfEmploymentTaxableProfit <= niLowerBand.to {
		niZeroRateNIC = 0
		remainingIncome = 0
	} else if totalSelfEmploymentTaxableProfit > niLowerBand.to && totalSelfEmploymentTaxableProfit <= niBasicBand.to {
		niZeroRateNIC = 0
		remainingIncome = totalSelfEmploymentTaxableProfit - niLowerBand.to
		niBasicRateAllocatedIncome = remainingIncome
		niBasicRateNIC = utilmath32.RoundDown(niBasicRateAllocatedIncome*niBasicBand.rate, 2)
	} else {
		niZeroRateNIC = 0
		niBasicRateAllocatedIncome = niBasicBand.to - niLowerBand.to
		niBasicRateNIC = utilmath32.RoundDown(niBasicRateAllocatedIncome*niBasicBand.rate, 2)
		remainingIncome = totalSelfEmploymentTaxableProfit - niBasicBand.to
		niHigherRateAllocatedIncome = remainingIncome
		niHigherRateNIC = utilmath32.RoundDown(niHigherRateAllocatedIncome*niHigherBand.rate, 2)
	}

	totalClass4NIC = niZeroRateNIC + niBasicRateNIC + niHigherRateNIC
	return utilmath32.RoundDown(totalClass4NIC, 2)
}
