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
	// No. 1 dan 2 sudah dipastikan prime Number.
	if n == 1 || n == 2 || n == 3 {
		fmt.Print(n)
		fmt.Println(" is a Prime Number, try to input value more than 3")
	} else {
		fmt.Println("1 is a Prime Number")
		fmt.Println("2 is a Prime Number")
		fmt.Println("3 is a Prime Number")
		// Jika user menginput n >= 3
		for i := 4; i <= n; i++ {
			// Memastikan Genap tidaklah Prime Number.
			if i%2 == 0 {
				fmt.Print(i)
				fmt.Println(" is Not Prime Number")
			} else if i%2 != 0 && i%3 == 0 {
				// Memastikan bil. ganjil Prime Number atau bukan
				fmt.Print(i)
				fmt.Println(" is Not Prime Number")
			} else if i%2 != 0 && i%5 == 0 {
				// Ditemukan beberapa bil. ganjil kelipatan 5 dianggap Prime Num pdhl tidak Prime Num.
				fmt.Print(i)
				fmt.Println(" is Not Prime Number")
			} else if i%2 != 0 && i%7 == 0 {
				// Ditemukan beberapa bil. ganjil kelipatan 7 dianggap Prime Num pdhl tidak Prime Num.
				fmt.Print(i)
				fmt.Println(" is Not Prime Number")
			} else {
				// jika bil. ganjil adalah Prime Number maka ditampung sini
				fmt.Print(i)
				fmt.Println(" is a Prime Number")
			}
		}
	}
	// PROBLEM: 25 dianggap Prime Number, karena 25 / 5 = 5 tsb

	/**
	Ini untuk menampilkan hasil prompt...
	fmt.Print("result: ")
	fmt.Println(n)
	*/
}
