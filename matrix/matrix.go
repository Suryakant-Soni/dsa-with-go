package main

import "fmt"

func main() {

	// Diagonal Matrix operations
	// arr := make([]int, 10)
	// SetDiangonalMatrix(arr, 1, 1, 12)
	// SetDiangonalMatrix(arr, 2, 2, 2)
	// SetDiangonalMatrix(arr, 3, 3, 6)
	// SetDiangonalMatrix(arr, 4, 4, 8)
	// SetDiangonalMatrix(arr, 5, 5, 9)
	// SetDiangonalMatrix(arr, 6, 6, 18)

	// fmt.Println(GetDiangonalMatrix(arr, 2, 2))
	// fmt.Println(GetDiangonalMatrix(arr, 5, 5))
	// fmt.Println(GetDiangonalMatrix(arr, 4, 4))
	// fmt.Println(GetDiangonalMatrix(arr, 7, 7))
	// fmt.Println(GetDiangonalMatrix(arr, 2, 9))
	// fmt.Println(GetDiangonalMatrix(arr, 2, 0))

	// get and set oper for lower triangular matrix
	//    number of elements in a triangular matrix effectively removing all the zeroes will be equal to n(n+1)/2
	// so for n = 5 , that is 15
	arr := make([]int, 15)
	SetLowerTriangularMatrixElement(1, 1, 8, arr)
	SetLowerTriangularMatrixElement(2, 1, 9, arr)
	SetLowerTriangularMatrixElement(3, 1, 10, arr)
	SetLowerTriangularMatrixElement(4, 1, 16, arr)
	SetLowerTriangularMatrixElement(5, 1, 71, arr)
	SetLowerTriangularMatrixElement(2, 2, 89, arr)
	SetLowerTriangularMatrixElement(3, 2, 7, arr)
	SetLowerTriangularMatrixElement(4, 2, 11, arr)
	SetLowerTriangularMatrixElement(5, 2, 78, arr)
	SetLowerTriangularMatrixElement(3, 3, 99, arr)
	SetLowerTriangularMatrixElement(4, 3, 28, arr)
	SetLowerTriangularMatrixElement(5, 3, 17, arr)
	SetLowerTriangularMatrixElement(4, 4, 77, arr)
	SetLowerTriangularMatrixElement(5, 4, 56, arr)
	SetLowerTriangularMatrixElement(5, 5, 43, arr)

	fmt.Println(GetLowerTriangularMatrixElement(5, 3, arr))
	fmt.Println(GetLowerTriangularMatrixElement(5, 5, arr))
	fmt.Println(GetLowerTriangularMatrixElement(5, 1, arr))
	fmt.Println(GetLowerTriangularMatrixElement(3, 3, arr))
	fmt.Println(GetLowerTriangularMatrixElement(2, 1, arr))
	fmt.Println(GetLowerTriangularMatrixElement(1, 1, arr))
	fmt.Println(GetLowerTriangularMatrixElement(4, 4, arr))
}
