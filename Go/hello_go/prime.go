package main

import "math"

// isPrime 関数: 指定された数が素数か判定
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// getPrimes 関数: n個の素数を求める
func getPrimes(count int) []int {
	primes := []int{}
	num := 2 // 最初の素数
	for len(primes) < count {
		if isPrime(num) {
			primes = append(primes, num)
		}
		num++
	}
	return primes
}
