package main

import (
	"testing"
)

func TestCalculateLoanPrincipal(t *testing.T) {
	tests := []struct {
		A          float64
		years      int
		annualRate float64
		expected   int
	}{
		{50000, 10, 0.03, 5188120},
		{70000, 20, 0.02, 13860660},
		{50000, 10, 0.00001, 5999698},
	}

	for _, tt := range tests {
		result := calculateLoanPrincipal(tt.A, tt.years, tt.annualRate)
		if result != tt.expected {
			t.Errorf("calculateLoanPrincipal(%f, %d, %f) = %d; want %d", tt.A, tt.years, tt.annualRate, result, tt.expected)
		}
	}
}
