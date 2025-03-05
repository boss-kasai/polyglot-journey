import { describe, it, expect } from 'bun:test'
import { getPrime } from '../src/prime'

describe('prime function', () => {
  it('should return prime numbers up to the given number', () => {
    expect(getPrime(2)).toEqual([2, 3])
    expect(getPrime(10)).toEqual([2, 3, 5, 7, 11, 13, 17, 19, 23, 29])
  })
})