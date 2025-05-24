package taxliability

import (
	"testing"
)

func TestCalculateTaxLiability(t *testing.T) {
	tests := []struct {
		name                             string
		totalProfitFromPayPensionsProfit float32
		want                             float32
	}{
		{
			totalProfitFromPayPensionsProfit: 22_400,
			want:                             1_966,
		},
		{
			totalProfitFromPayPensionsProfit: 38_900,
			want:                             5_266,
		},
		{
			totalProfitFromPayPensionsProfit: 68_800,
			want:                             14_952,
		},
		{
			totalProfitFromPayPensionsProfit: 117_000,
			want:                             37_632,
		},
		{
			totalProfitFromPayPensionsProfit: 130_000,
			want:                             44_703,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateTaxLiability(tt.totalProfitFromPayPensionsProfit)
			if got != tt.want {
				t.Errorf(
					"%f > CalculateTaxLiability() = %v, want %v",
					tt.totalProfitFromPayPensionsProfit,
					got,
					tt.want,
				)
			}
		})
	}
}
