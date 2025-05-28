package propertydomain

import (
	utilmath32 "tax_calculator/engine/internal/utils/util_math32"
)

type PropertyIncomeEntity struct {
	totalRentsReceived             float32
	premiumsOfLeaseGrant           float32
	reversePremiums                float32
	otherPropertyIncome            float32
	rarRentReceived                float32
	totalIncomeFromUkPropertyOther float32
}

func NewPropertyIncomeEntity(
	totalRentsReceived float32,
	premiumsOfLeaseGrant float32,
	reversePremiums float32,
	otherPropertyIncome float32,
	rarRentReceived float32,
) PropertyIncomeEntity {
	propertyIncomeEntity := &PropertyIncomeEntity{
		totalRentsReceived,
		premiumsOfLeaseGrant,
		reversePremiums,
		otherPropertyIncome,
		rarRentReceived,
		utilmath32.RoundDown(
			totalRentsReceived+premiumsOfLeaseGrant+reversePremiums+otherPropertyIncome+rarRentReceived,
			2,
		),
	}
	return *propertyIncomeEntity
}

type PropertyExpensesEntity struct {
	consolidatedExpenses             float32
	premisesRunningCosts             float32
	repairsAndMainteance             float32
	financialCosts                   float32
	professionalFees                 float32
	costOfServices                   float32
	other                            float32
	travelCosts                      float32
	totalExpensesFromUkPropertyOther float32
}

func NewPropertyExpensesEntity(
	consolidatedExpenses float32,
	premisesRunningCosts float32,
	repairsAndMainteance float32,
	financialCosts float32,
	professionalFees float32,
	costOfServices float32,
	other float32,
	travelCosts float32,
) PropertyExpensesEntity {
	propertyExpensesEntity := &PropertyExpensesEntity{
		consolidatedExpenses,
		premisesRunningCosts,
		repairsAndMainteance,
		financialCosts,
		professionalFees,
		costOfServices,
		other,
		travelCosts,
		utilmath32.RoundUp(
			consolidatedExpenses+premisesRunningCosts+repairsAndMainteance+financialCosts+professionalFees+costOfServices+other+travelCosts,
			2,
		),
	}

	return *propertyExpensesEntity
}

type PropertyAdditionsEntity struct {
	privateUseAdjustment float32
	balancingCharge      float32
	bpraBalancingCharge  float32

	totalAdditionsFromUkPropertyOther float32
}

func NewPropertyAdditionsEntity(
	privateUseAdjustment float32,
	balancingCharge float32,
	bpraBalancingCharge float32,
) PropertyAdditionsEntity {
	return PropertyAdditionsEntity{
		privateUseAdjustment,
		balancingCharge,
		bpraBalancingCharge,
		utilmath32.RoundDown(privateUseAdjustment+balancingCharge+bpraBalancingCharge, 2),
	}
}

type PropertyDeductionsEntity struct {
	zeroEmissionsGoodsVehicleAllowance  float32
	annualInvestmentAllowance           float32
	costOfReplacingDomesticItems        float32
	businessPremisesRenovationAllowance float32
	propertyAllowance                   float32
	otherCaptialAllowance               float32
	electricChargePointAllowance        float32
	zeroEmissionCarAllowance            float32
	structuredBuildingAllowance         float32
	enhancedStructuredBuildingAllowance float32
	rarReliefClaimed                    float32
	totalDeductionsFromUkPropertyOther  float32
}

func NewPropertyDeductionsEntity(
	zeroEmissionsGoodsVehicleAllowance float32,
	annualInvestmentAllowance float32,
	costOfReplacingDomesticItems float32,
	businessPremisesRenovationAllowance float32,
	propertyAllowance float32,
	otherCaptialAllowance float32,
	electricChargePointAllowance float32,
	zeroEmissionCarAllowance float32,
	structuredBuildingAllowance float32,
	enhancedStructuredBuildingAllowance float32,
	rarReliefClaimed float32,
) PropertyDeductionsEntity {
	var totalDeductionsFromUkPropertyOther float32
	if propertyAllowance > 0 {
		totalDeductionsFromUkPropertyOther = propertyAllowance
	} else {
		totalDeductionsFromUkPropertyOther = zeroEmissionsGoodsVehicleAllowance + annualInvestmentAllowance + costOfReplacingDomesticItems + businessPremisesRenovationAllowance + otherCaptialAllowance + electricChargePointAllowance + zeroEmissionCarAllowance + structuredBuildingAllowance + enhancedStructuredBuildingAllowance + rarReliefClaimed
	}
	totalDeductionsFromUkPropertyOther = utilmath32.RoundUp(totalDeductionsFromUkPropertyOther, 2)

	return PropertyDeductionsEntity{
		zeroEmissionsGoodsVehicleAllowance,
		annualInvestmentAllowance,
		costOfReplacingDomesticItems,
		businessPremisesRenovationAllowance,
		propertyAllowance,
		otherCaptialAllowance,
		electricChargePointAllowance,
		zeroEmissionCarAllowance,
		structuredBuildingAllowance,
		enhancedStructuredBuildingAllowance,
		rarReliefClaimed,
		totalDeductionsFromUkPropertyOther,
	}
}

type TotalTaxablePropertyProfit struct {
	TaxableProfitFromUkPropertyOther float32
	TaxableLossFromUkPropertyOther   float32
}

func NewSelfEmploymentTotalTaxableSelfEmploymentProfit(
	taxableProfitFromUkPropertyOther float32,
	taxableLossFromUkPropertyOther float32,
) TotalTaxablePropertyProfit {
	return TotalTaxablePropertyProfit{
		taxableProfitFromUkPropertyOther,
		taxableLossFromUkPropertyOther,
	}
}
