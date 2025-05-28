package personalallowancedomain

import utilmath32 "tax_calculator/engine/internal/utils/util_math32"

type PersonalAllowanceService struct{}

const (
	PersonalAllowance            = 12_570
	reducedAllowanceLimit        = 100_000
	personalAllowanceCutoffLimit = 125_140
)

func (personalAllowanceService *PersonalAllowanceService) CalculatePersonalAllowance(
	adjustedNetIncome float32,
) float32 {
	if adjustedNetIncome <= reducedAllowanceLimit {
		return PersonalAllowance
	}

	if adjustedNetIncome > reducedAllowanceLimit &&
		adjustedNetIncome < personalAllowanceCutoffLimit {
		allowanceReductionAmount := utilmath32.RoundDown(
			(adjustedNetIncome-reducedAllowanceLimit)/2,
			0,
		)
		return utilmath32.RoundUp(PersonalAllowance-allowanceReductionAmount, 0)
	}

	return 0
}

func NewPersonalAllowanceService() PersonalAllowanceService {
	return PersonalAllowanceService{}
}
