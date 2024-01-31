package main

import "fmt"

func SetDiangonalMatrix(arr []int, i int, j int, e int) {

	if i == j || i == 0 {
		arr[i-1] = e
	} else {
		fmt.Println("this cannot be a diagonal element")
	}
}
func GetDiangonalMatrix(arr []int, i int, j int) int {

	if (i == j) && (i > 0) && (j > 0) && i < len(arr) && j < len(arr) {
		return arr[i-1]
	} else {
		fmt.Println("this cannot be a diagonal element")
		return 0
	}
}
