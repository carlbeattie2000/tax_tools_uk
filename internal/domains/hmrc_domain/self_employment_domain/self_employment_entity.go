package selfemploymentdomain

import (
	utilmath32 "tax_calculator/engine/internal/utils/util_math32"
)

type SelfEmploymentIncomeEntity struct {
	turnover                  float32
	other                     float32
	totalSelfEmploymentIncome float32
}

func NewSelfEmploymentEntity(turnover float32, other float32) SelfEmploymentIncomeEntity {
	selfEmploymentEntity := &SelfEmploymentIncomeEntity{
		turnover:                  turnover,
		other:                     other,
		totalSelfEmploymentIncome: utilmath32.RoundDown(turnover+other, 2),
	}
	return *selfEmploymentEntity
}

type SelfEmploymentExpensesEntity struct {
	consolidatedExpenses                float32
	costOfGoodsAllowable                float32
	paymentToSubcontractorsAllowable    float32
	wagesAndStaffCostsAllowable         float32
	carVanTravelExpensesAllowable       float32
	premisesRunningCostsAllowable       float32
	maintenanceCostsAllowable           float32
	adminCostsAllowable                 float32
	intrestOnBankOtherLoansAllowable    float32
	financeChargesAllowable             float32
	irrecoverableDebtsAllowable         float32
	profressionalFeesAllowable          float32
	depreciationAllowable               float32
	otherExpensesAllowable              float32
	advertisingCostsAllowable           float32
	businessEntertainmentCostsAllowable float32
	totalSelfEmploymentExpenses         float32
}

func NewSelfEmploymentExpensesEntity(
	consolidatedExpenses float32,
	costOfGoodsAllowable float32,
	paymentToSubcontractorsAllowable float32,
	wagesAndStaffCostsAllowable float32,
	carVanTravelExpensesAllowable float32,
	premisesRunningCostsAllowable float32,
	maintenanceCostsAllowable float32,
	adminCostsAllowable float32,
	intrestOnBankOtherLoansAllowable float32,
	financeChargesAllowable float32,
	irrecoverableDebtsAllowable float32,
	profressionalFeesAllowable float32,
	depreciationAllowable float32,
	otherExpensesAllowable float32,
	advertisingCostsAllowable float32,
	businessEntertainmentCostsAllowable float32,
) SelfEmploymentExpensesEntity {
	selfEmplymentExpensesEntity := &SelfEmploymentExpensesEntity{
		consolidatedExpenses,
		costOfGoodsAllowable,
		paymentToSubcontractorsAllowable,
		wagesAndStaffCostsAllowable,
		carVanTravelExpensesAllowable,
		premisesRunningCostsAllowable,
		maintenanceCostsAllowable,
		adminCostsAllowable,
		intrestOnBankOtherLoansAllowable,
		financeChargesAllowable,
		irrecoverableDebtsAllowable,
		profressionalFeesAllowable,
		depreciationAllowable,
		otherExpensesAllowable,
		advertisingCostsAllowable,
		businessEntertainmentCostsAllowable,
		utilmath32.RoundUp(
			consolidatedExpenses+costOfGoodsAllowable+paymentToSubcontractorsAllowable+wagesAndStaffCostsAllowable+carVanTravelExpensesAllowable+premisesRunningCostsAllowable+maintenanceCostsAllowable+adminCostsAllowable+intrestOnBankOtherLoansAllowable+financeChargesAllowable+irrecoverableDebtsAllowable+profressionalFeesAllowable+depreciationAllowable+otherExpensesAllowable+advertisingCostsAllowable+businessEntertainmentCostsAllowable,
			2,
		),
	}

	return *selfEmplymentExpensesEntity
}

type SelfEmploymentAdditionsEntity struct {
	costOfGoodsDisallowable                float32
	paymentToSubcontractorsDisallowable    float32
	wagesAndStaffCostsDisallowable         float32
	carVanTravelExpensesDisallowable       float32
	premisesRunningCostsDisallowable       float32
	maintenanceCostsDisallowable           float32
	adminCostsDisallowable                 float32
	intrestOnBankOtherLoansDisallowable    float32
	financeChargesDisallowable             float32
	irrecoverableDebtsDisallowable         float32
	profressionalFeesDisallowable          float32
	depreciationDisallowable               float32
	otherExpensesDisallowable              float32
	advertisingCostsDisallowable           float32
	businessEntertainmentCostsDisallowable float32
	outstandingBusinessIncome              float32
	balancingChargeOther                   float32
	balancingChargeBpra                    float32
	goodAndServiceOwnUse                   float32
	totalSelfEmploymentAdditons            float32
}

func NewSelfEmploymentAdditionsEntity(
	costOfGoodsDisallowable float32,
	paymentToSubcontractorsDisallowable float32,
	wagesAndStaffCostsDisallowable float32,
	carVanTravelExpensesDisallowable float32,
	premisesRunningCostsDisallowable float32,
	maintenanceCostsDisallowable float32,
	adminCostsDisallowable float32,
	intrestOnBankOtherLoansDisallowable float32,
	financeChargesDisallowable float32,
	irrecoverableDebtsDisallowable float32,
	profressionalFeesDisallowable float32,
	depreciationDisallowable float32,
	otherExpensesDisallowable float32,
	advertisingCostsDisallowable float32,
	businessEntertainmentCostsDisallowable float32,

	outstandingBusinessIncome float32,
	balancingChargeOther float32,
	balancingChargeBpra float32,
	goodAndServiceOwnUse float32,
) SelfEmploymentAdditionsEntity {
	selfEmploymentAdditionsEntity := &SelfEmploymentAdditionsEntity{
		costOfGoodsDisallowable,
		paymentToSubcontractorsDisallowable,
		wagesAndStaffCostsDisallowable,
		carVanTravelExpensesDisallowable,
		premisesRunningCostsDisallowable,
		maintenanceCostsDisallowable,
		adminCostsDisallowable,
		intrestOnBankOtherLoansDisallowable,
		financeChargesDisallowable,
		irrecoverableDebtsDisallowable,
		profressionalFeesDisallowable,
		depreciationDisallowable,
		otherExpensesDisallowable,
		advertisingCostsDisallowable,
		businessEntertainmentCostsDisallowable,

		outstandingBusinessIncome,
		balancingChargeOther,
		balancingChargeBpra,
		goodAndServiceOwnUse,
		utilmath32.RoundDown(
			costOfGoodsDisallowable+paymentToSubcontractorsDisallowable+wagesAndStaffCostsDisallowable+
				carVanTravelExpensesDisallowable+premisesRunningCostsDisallowable+maintenanceCostsDisallowable+
				adminCostsDisallowable+intrestOnBankOtherLoansDisallowable+financeChargesDisallowable+
				irrecoverableDebtsDisallowable+profressionalFeesDisallowable+depreciationDisallowable+
				otherExpensesDisallowable+advertisingCostsDisallowable+businessEntertainmentCostsDisallowable+
				outstandingBusinessIncome+balancingChargeOther+balancingChargeBpra+goodAndServiceOwnUse,
			2,
		),
	}

	return *selfEmploymentAdditionsEntity
}

type SelfEmploymentDeductionsEntity struct {
	tradingAllowance                    float32
	annualInvestmentAllowance           float32
	capitalAllowanceMainPool            float32
	capitalAllowanceSpecialRatePool     float32
	zeroEmissionGoods                   float32
	businessPremisesRenvoationAllowance float32
	enhancedCapitalAllowance            float32
	allowanceOnSales                    float32
	capitalAllowanceSingleAssetPool     float32
	electricChargePointAllowance        float32
	zeroEmissionsCarAllowance           float32
	structuredBuildingAllowance         float32
	enhancedStructureBuilingAllowance   float32
	includedNoTaxableProfits            float32
	totalSelfEmploymentDeductions       float32
}

func NewSelfEmploymentDeductionsEntity(
	tradingAllowance float32,
	annualInvestmentAllowance float32,
	capitalAllowanceMainPool float32,
	capitalAllowanceSpecialRatePool float32,
	zeroEmissionGoods float32,
	businessPremisesRenvoationAllowance float32,
	enhancedCapitalAllowance float32,
	allowanceOnSales float32,
	capitalAllowanceSingleAssetPool float32,
	electricChargePointAllowance float32,
	zeroEmissionsCarAllowance float32,
	structuredBuildingAllowance float32,
	enhancedStructureBuilingAllowance float32,
	includedNoTaxableProfits float32,
) SelfEmploymentDeductionsEntity {
	selfEmploymentDeductionsEntity := &SelfEmploymentDeductionsEntity{
		tradingAllowance,
		annualInvestmentAllowance,
		capitalAllowanceMainPool,
		capitalAllowanceSpecialRatePool,
		zeroEmissionGoods,
		businessPremisesRenvoationAllowance,
		enhancedCapitalAllowance,
		allowanceOnSales,
		capitalAllowanceSingleAssetPool,
		electricChargePointAllowance,
		zeroEmissionsCarAllowance,
		structuredBuildingAllowance,
		enhancedStructureBuilingAllowance,
		includedNoTaxableProfits,
		utilmath32.RoundUp(
			tradingAllowance+annualInvestmentAllowance+capitalAllowanceMainPool+capitalAllowanceSpecialRatePool+
				zeroEmissionGoods+businessPremisesRenvoationAllowance+enhancedCapitalAllowance+allowanceOnSales+
				capitalAllowanceSingleAssetPool+includedNoTaxableProfits+electricChargePointAllowance+
				zeroEmissionsCarAllowance+structuredBuildingAllowance+enhancedStructureBuilingAllowance,
			2,
		),
	}

	return *selfEmploymentDeductionsEntity
}

type SelfEmploymentAccountingAdjustmentsEntity struct {
	basisAmjustment                          float32
	accountingAdjustment                     float32
	averagingAdjustment                      float32
	totalSelfEmploymentAccountingAdjustments float32
}

func NewSelfEmploymentAccountingAdjustmentsEntity(
	basisAdjustment float32,
	accountingAdjustment float32,
	averagingAdjustment float32,
	totalSelfEmploymentAccountingAdjustments float32,
) SelfEmploymentAccountingAdjustmentsEntity {
	selfEmploymentAccountingAdjustmentsEntity := &SelfEmploymentAccountingAdjustmentsEntity{
		basisAdjustment,
		accountingAdjustment,
		averagingAdjustment,
		utilmath32.RoundDown(basisAdjustment+accountingAdjustment+averagingAdjustment, 2),
	}

	return *selfEmploymentAccountingAdjustmentsEntity
}

type TotalTaxableSelfEmploymentProfit struct {
	TaxableProfitFromSelfEmployment float32
	TaxableLossFromSelfEmployment   float32
}

func NewSelfEmploymentTotalTaxableSelfEmploymentProfit(
	taxableProfitFromSelfEmployment float32,
	taxableLossFromSelfEmployment float32,
) TotalTaxableSelfEmploymentProfit {
	return TotalTaxableSelfEmploymentProfit{
		taxableProfitFromSelfEmployment,
		taxableLossFromSelfEmployment,
	}
}
