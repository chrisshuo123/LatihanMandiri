package main

import (
	"fmt"
)

/* Buatlah sebuah program menggabungkan 2 array yang diberikan, dan jangan
sampai terdapat nama yang sama di data yang sudah tergabung tadi! */

func ArrayMerge(arrayA, arrayB []string) []string {
	// your code here...
	arrayA = append[]string(arrayA, arrayB)
	
}

func main() {
	// Test Cases
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	// ["king", "devil jin", "akuma", "eddie", "steve", "geese]

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	// ["sergei", "jin", "steve", "bryan"]

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	// ["alisa", "yoshimitsu", "devil jin", "law"]

	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	// ["hwoarang"]

	fmt.Println(ArrayMerge([]string{}, []string{}))
	// []
}
