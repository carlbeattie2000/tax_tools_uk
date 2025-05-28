package statebenefitsdomain

type StateBenefitsEntity struct {
	bereavementAllowance                                        float32
	employmentSupportAllowance                                  float32
	incapacityBenefit                                           float32
	jobSeekersAllowance                                         float32
	otherStateBenefits                                          float32
	statePension                                                float32
	statePensionLumpSum                                         float32
	totalStateBenefitsIncomeExcludingStatePensionLumpSumBenefit float32
}

func NewStateBenefitsEntity(
	bereavementAllowance float32,
	employmentSupportAllowance float32,
	incapacityBenefit float32,
	jobSeekersAllowance float32,
	otherStateBenefits float32,
	statePension float32,
	statePensionLumpSum float32,
) StateBenefitsEntity {
	return StateBenefitsEntity{
		bereavementAllowance,
		employmentSupportAllowance,
		incapacityBenefit,
		jobSeekersAllowance,
		otherStateBenefits,
		statePension,
		statePensionLumpSum,
		bereavementAllowance + employmentSupportAllowance + incapacityBenefit + jobSeekersAllowance + otherStateBenefits + statePension,
	}
}
