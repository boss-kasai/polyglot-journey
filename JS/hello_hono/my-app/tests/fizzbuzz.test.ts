import { describe, it, expect } from 'bun:test'
import { fizzBuzz } from '../src/fizzbuzz'

describe('fizzBuzz function', () => {
  it('should return "Fizz" for multiples of 3', () => {
    expect(fizzBuzz(3)).toBe('Fizz')
    expect(fizzBuzz(6)).toBe('Fizz')
  })

  it('should return "Buzz" for multiples of 5', () => {
    expect(fizzBuzz(5)).toBe('Buzz')
    expect(fizzBuzz(10)).toBe('Buzz')
  })

  it('should return "FizzBuzz" for multiples of 15', () => {
    expect(fizzBuzz(15)).toBe('FizzBuzz')
    expect(fizzBuzz(30)).toBe('FizzBuzz')
  })

  it('should return the number for non-multiples of 3 or 5', () => {
    expect(fizzBuzz(1)).toBe('1')
    expect(fizzBuzz(2)).toBe('2')
  })
})
