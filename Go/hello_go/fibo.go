package main

import "strconv"

// fibo はフィボナッチ数を計算する関数
func fibo(n int) string {
	if n < 0 {
		return "Invalid input"
	}
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return strconv.Itoa(a)
}
