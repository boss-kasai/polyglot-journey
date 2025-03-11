package main

import (
	"math"
)

// calculateLoanPrincipal は月の返済額 A, 返済年数 years, 年利 annualRate から
// 借入可能額 (元本) を計算する関数
func calculateLoanPrincipal(A float64, years int, annualRate float64) int {
	// 返済回数 n (年数 * 12)
	n := float64(years * 12)

	// 月利 i (年利を12で割り、月複利で計算)
	i := math.Pow(1+annualRate, 1.0/12.0) - 1

	// A = P * ( i / (1 - (1 + i)^(-n)) )
	// → P = A * (1 - (1 + i)^(-n)) / i
	numerator := 1 - math.Pow(1+i, -n)
	P := A * (numerator / i)

	return int(math.Round(P))
}
