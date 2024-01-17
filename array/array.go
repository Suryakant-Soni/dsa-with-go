package main

import (
	"log"
)

func main() {
	var arr1 = []int{4, 8, 13, 16, 20, 25, 28, 33,38}
//	var arr2 = []int{-44, -11, 4, 8, 13, -9, 16, 20, -11, 25, -13, 28, 33, -1}
	var merge1 = []int{2, 6, 8, 13, 16, 19, 22,33,38}
	// var merge2 = []int{3, 4, 6, 8, 12, 14, 15, 24, 28}
	// mergeArrays(merge1, merge2)
	// InsertInSortedArray(&arr1, 17)
	// SeggregateNegative(&arr2)
	//log.Println("is sorted ?", CheckIfSorted(arr1))
	log.Println("UnionArrays", UnionArrays(arr1,merge1))
	// log.Println("new array", arr1)
}

func mergeArrays(m1 []int, m2 []int) {
	m3 := make([]int, (len(m1) + len(m2)))
	i, j, k := 0, 0, 0
	for i < len(m1) && j < len(m2) {
		if m1[i] > m2[j] {
			m3[k] = m2[j]
			j++
			k++
		} else {
			m3[k] = m1[i]
			i++
			k++
		}
	}
	for i < len(m1) {
		m3[k] = m1[i]
		i++
		k++
	}
	for j < len(m2) {
		m3[k] = m2[j]
		j++
		k++
	}
	log.Println("merged array", m3)
}

func SeggregateNegative(a *[]int) {
	i, j := 0, len(*a)-1
	for i < j {
		if (*a)[i] < 0 {
			i++
			if (*a)[j] >= 0 {
				j--
			}
			continue
		}
		if (*a)[i] >= 0 {
			if (*a)[j] >= 0 {
				j--
				continue
			} else {
				swap(i, j, a)
				i++
				j--
				continue
			}
		}
	}
}

func swap(i, j int, a *[]int) {
	temp := (*a)[i]
	(*a)[i] = (*a)[j]
	(*a)[j] = temp
}

func InsertInSortedArray(a *[]int, n int) {
	log.Println("len", len(*a))
	*a = append(*a, (*a)[len(*a)-1])
	log.Println("len", len(*a))
	for i := len(*a) - 3; i > 0; i-- {
		if (*a)[i] < n {
			(*a)[i+1] = n
			break
		}
		(*a)[i+1] = (*a)[i]
		if i == 0 {
			(*a)[0] = n
		}
	}
}

func CheckIfSorted(a []int) bool{
	for i := 1; i < len(a); i++{
		if a[i] < a[i-1]{
			return false
		}
	}
		return true
}

func UnionArrays(a []int,b[]int) (c []int){
	// i, j
	i,j := 0,0
	for( i < len(a) && j < len(b)){
		if a[i] == b[j] {
			c = append(c, a[i])
			i++
			j++
			continue
		}
		if a[i] < b[j]{
			c = append(c, a[i])
			i++
			continue
		}
		if a[i] > b[j]{
			c = append(c, b[j])
			j++
			continue
		}
	}
	return c
}