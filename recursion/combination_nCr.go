package main

func Combination(n int, r int) int {
	if r == 0 || n == r {
		return 1
	} else {
		return Combination(n-1, r-1) + Combination(n-1, r)
	}
}
