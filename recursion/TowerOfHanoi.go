package main

import "fmt"

func TowerOfHanoi(n int, a int, b int, c int) {
	if n > 0 {
		TowerOfHanoi(n-1, a, c, b)
		fmt.Printf("\n ( %v , %v)", a, c)
		TowerOfHanoi(n-1, b, a, c)
	}
}