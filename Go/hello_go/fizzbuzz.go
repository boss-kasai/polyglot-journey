package main

import "strconv"

func fizzBuzz(num int) string {
	if num%15 == 0 {
		return "FizzBuzz"
	} else if num%3 == 0 {
		return "Fizz"
	} else if num%5 == 0 {
		return "Buzz"
	}
	return strconv.Itoa(num)
}
