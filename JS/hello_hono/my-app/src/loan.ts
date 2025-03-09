/**
 * 月の返済額 A, 返済年数 years, 年利 annualRate から
 * 借入可能額 (元本) を計算する関数
 */
export function calculateLoanPrincipal(A: number, years: number, annualRate: number): number {
    // 返済回数 n (年数 * 12)
    const n = years * 12

    // 月利 i (年利を12で割り、月複利で計算)
    const i = Math.pow(1 + annualRate, 1 / 12) - 1

    // A = P * ( i / (1 - (1 + i)^(-n)) )
    // → P = A * (1 - (1 + i)^(-n)) / i
    const numerator = 1 - Math.pow(1 + i, -n)
    const P = A * (numerator / i)

    return Math.round(P)
  }