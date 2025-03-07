import { describe, it, expect } from 'bun:test'
import { getFibonacci } from '../src/fibonacci'

describe('Fibonacci function', () => {
    it('should return the n-th Fibonacci number', () => {
        expect(getFibonacci(0)).toBe(0)
        expect(getFibonacci(1)).toBe(1)
        expect(getFibonacci(2)).toBe(1)
        expect(getFibonacci(3)).toBe(2)
        expect(getFibonacci(4)).toBe(3)
        expect(getFibonacci(5)).toBe(5)
        expect(getFibonacci(6)).toBe(8)
        expect(getFibonacci(7)).toBe(13)
        expect(getFibonacci(8)).toBe(21)
        expect(getFibonacci(9)).toBe(34)
        expect(getFibonacci(10)).toBe(55)
    })
})