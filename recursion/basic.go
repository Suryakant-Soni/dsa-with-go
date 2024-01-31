package main

import "fmt"

func SumOfFirstNNaturalNumbers(n int64) int64 {
	if n > 0 {
		return n + SumOfFirstNNaturalNumbers(n-1)
	} else {
		return 0
	}
}

func FactorialN(n int64) (p int64) {
	if n > 1 {
		p = n * FactorialN(n-1)
		fmt.Println(p)
		return p
	} else {
		p = 1
		return p
	}
}

func Power(m int64, n int64) int64 {
	if n > 0 {
		return m * Power(m, n-1)
	} else {
		return 1
	}
}
