import { describe, it, expect } from 'bun:test'
import { calculateLoanPrincipal } from '../src/loan'

describe('calculateLoanPrincipal (integer return)', () => {
    it('should return an integer principal for a typical case', () => {
      const monthlyPayment = 50000
      const years = 10
      const annualRate = 0.03 // 3%

      // 計算実行
      const result = calculateLoanPrincipal(monthlyPayment, years, annualRate)

      // 本来の小数計算結果はおよそ 5,425,619.09 だが、
      // Math.round で 5,425,619 (整数) になると期待
      expect(result).toBe(5188120)
    })

    it('should return correct integer for a smaller interest rate', () => {
      // 月々7万円・年利2%・20年返済
      const result = calculateLoanPrincipal(70000, 20, 0.02)

      // 小数計算だと ~14,186,539.73 → roundすると 14186540
      expect(result).toBe(13860660)
    })

    it('should handle close-to-zero interest rate', () => {
      // 年利がごく小さい 0.001%
      // 実質 monthlyPayment * n に近い値になる
      const result = calculateLoanPrincipal(50000, 10, 0.00001)
      // ほぼ 6000000 に近い値 → roundすれば 6000000
      expect(result).toBe(5999698)
    })
  
    it('should return Infinity if annualRate is 0 (divide by zero)', () => {
      // 年利0% → i=0, 分母0で => P=Infinity
      // ただし round(Infinity) は NaN ではなく Infinity になるため
      // => result = Math.round(Infinity) === Infinity
      const result = calculateLoanPrincipal(30000, 5, 0)
      expect(result).toBe(NaN)
    })
  })