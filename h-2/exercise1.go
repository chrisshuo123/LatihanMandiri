package main

import (
	"fmt"
)

// 1. buatlah sebuah program untuk menentukan bilang prima
func main() {
	// declare & input variable
	var n int
	fmt.Println("EXERCISE-1: MENENTUKAN BIL. PRIMA")
	fmt.Print("Enter your Number n: ")
	fmt.Scanln(&n)

	// Processing Prime Number of n
	/*fmt.Print("Alright, You input ")
	fmt.Print(n)
	fmt.Println(", and here's the result:")*/

	// PROCESSING & OUTPUT HASIL PRIME NUMBER:

	// Jika user menginput n di 1 atau 2 karena,
	// No. 1, 2, dan 3 sudah dipastikan prime Number.
	if n == 1 || n == 2 || n == 3 {
		fmt.Print(n)
		fmt.Println(" is a Prime Number, try to input value more than 3")
	} else {
		// Jika user menginput diatas 3
		if n%2 == 0 {
			// Bil. genap sudah pasti bukan prime number kecuali 2
			fmt.Print(n)
			fmt.Println(" is Not Prime Number")
		} else if n%2 != 0 && n%3 == 0 {
			// Memfilter Bil. ganjil yang tidak prime number
			fmt.Print(n)
			fmt.Println(" is Not Prime Number")
		} else if n%2 != 0 && n%5 == 0 {
			// Terdapat bil. Ganjil kelipatan 5 Bukan Prime Num. dianggap Prime Num.
			fmt.Print(n)
			fmt.Println(" is Not Prime Number")
		} else if n%2 != 0 && n%7 == 0 {
			// Terdapat bil. Ganjil kelipatan 7 Bukan Prime Num. dianggap Prime Num.
			fmt.Print(n)
			fmt.Println(" is Not Prime Number")
		} else {
			// jika bil. ganjil adalah Prime Number maka ditampung sini
			fmt.Print(n)
			fmt.Println(" is a Prime Number")
		}
	}
	// PROBLEM: 25 dianggap Prime Number, karena 25 / 5 = 5 tsb

	/**
	Ini untuk menampilkan hasil prompt...
	fmt.Print("result: ")
	fmt.Println(n)
	*/
}
