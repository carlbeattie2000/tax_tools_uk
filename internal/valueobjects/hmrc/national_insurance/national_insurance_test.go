package nationalinsurance_test

import (
	"tax_calculator/engine/internal/valueobjects/hmrc/national_insurance"
	"testing"
)

func TestCalculateNationalInsurance(t *testing.T) {
	tests := []struct {
		name                             string
		totalSelfEmploymentTaxableProfit float32
		want                             float32
	}{
		{
			totalSelfEmploymentTaxableProfit: 0,
			want:                             0,
		},
		{
			totalSelfEmploymentTaxableProfit: 12000,
			want:                             0,
		},
		{
			totalSelfEmploymentTaxableProfit: 12570,
			want:                             0,
		},
		{
			totalSelfEmploymentTaxableProfit: 45_000,
			want:                             1_945.80,
		},
		{
			totalSelfEmploymentTaxableProfit: 70_000,
			want:                             2_656.60,
		},
		{
			totalSelfEmploymentTaxableProfit: 100_000,
			want:                             3_256.60,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := nationalinsurance.CalculateNationalInsurance(tt.totalSelfEmploymentTaxableProfit)
			if got != tt.want {
				t.Errorf(
					"CalculateNationalInsurance() = %v, want %v, income %v",
					got,
					tt.want,
					tt.totalSelfEmploymentTaxableProfit,
				)
			}
		})
	}
}
