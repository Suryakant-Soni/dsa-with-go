package main

import "fmt"

func main() {
	duplicate := "ccfindingrrkk"
	findduplicatesinString(duplicate)
}

func findduplicatesinString(dup string) {
	var h int
	var x int
	for i := 0; i < len(dup); i++ {
		// reset x
		x = 1
		// shift x by the bits of difference from reference of 97 = small a ascii value
		x = x << (dup[i] - 97)
		// if and return true that means the bit is already present
		if x & h > 0 {
       fmt.Printf("duplicate character %c \n", dup[i])
		}else{
			//merge the bit if the char is not found
			h = x | h
		}
	}
}
