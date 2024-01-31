package main

func SetLowerTriangularMatrixElement(i int, j int, elem int, arr []int) {
	if i < j {
		panic("lower triangular matrix cannot have j greater than i")
	}
	index := i*(i-1)/2 + j - 1
	if index >= len(arr) {
		panic("index is getting calculated more than the calculated capacity of array")
	}
	arr[index] = elem
}

func GetLowerTriangularMatrixElement(i int, j int, arr []int) int {
	if j > i {
		panic("lower triangular matrix cannot have j greater than i")
	}
	index := i*(i-1)/2 + j - 1
	if index >= len(arr) {
		panic("index is getting calculated more than the calculated capacity of array")
	}
	return arr[index]
}
