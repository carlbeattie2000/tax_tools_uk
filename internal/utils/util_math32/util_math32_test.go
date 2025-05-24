package utilmath32

import (
	"testing"
)

func TestRoundDown(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		value  float32
		places float32
		want   float32
	}{
		{
			value:  1.345,
			places: 2,
			want:   1.34,
		},
		{
			value:  1.345,
			places: 1,
			want:   1.3,
		},
		{
			value:  1.345,
			places: 0,
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RoundDown(tt.value, tt.places)
			if got != tt.want {
				t.Errorf("RoundDown() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoundUp(t *testing.T) {
	tests := []struct {
		name   string
		value  float32
		places float32
		want   float32
	}{
		{
			value:  1.345,
			places: 2,
			want:   1.35,
		},
		{
			value:  1.345,
			places: 1,
			want:   1.4,
		},
		{
			value:  1.345,
			places: 0,
			want:   2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RoundUp(tt.value, tt.places)
			if got != tt.want {
				t.Errorf("RoundUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIncreasePercentage(t *testing.T) {
	tests := []struct {
		name          string
		startingValue float32
		finalValue    float32
		want          float32
	}{
		{
			startingValue: 1000,
			finalValue:    2000,
			want:          100,
		},
		{
			startingValue: 1000,
			finalValue:    4500,
			want:          350,
		},
		{
			startingValue: 3800,
			finalValue:    4500,
			want:          18.42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IncreasePercentage(tt.startingValue, tt.finalValue)
			if got != tt.want {
				t.Errorf("IncreasePercentage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRound(t *testing.T) {
	tests := []struct {
		name   string
		value  float32
		places float32
		want   float32
	}{
		{
			value:  44.44,
			places: 0,
			want:   44,
		},
		{
			value:  44.58,
			places: 2,
			want:   44.58,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Round(tt.value, tt.places)
			if got != tt.want {
				t.Errorf("Round() = %v, want %v", got, tt.want)
			}
		})
	}
}
