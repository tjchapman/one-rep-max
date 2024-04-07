package main

import (
	"math"
	"testing"
)

func AlmostEqual(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestCalculateOneRepMax(t *testing.T) {
	tests := []struct {
		weight    float64
		reps      float64
		want      float64
		tolerance float64
	}{
		{100, 1, 100, 0},
		{100, 10, 133, 0.5},
		{200, 5, 225, 0.5},
	}

	for _, tt := range tests {
		got := CalculateOneRepMax(tt.weight, tt.reps)
		if !AlmostEqual(got, tt.want, tt.tolerance) {
			t.Errorf("CalculateOneRepMax(%v, %v) = %v, want %v (within tolerance %v)", tt.weight, tt.reps, got, tt.want, tt.tolerance)
		}
	}
}
