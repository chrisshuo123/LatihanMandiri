package main

import "fmt"

//2. buatlah sebuah program untuk menentukan bilangan kelipatan 7
func main() {
	var n int
	fmt.Print("Input nilai n: ")
	fmt.Scanln(&n)

	// Menganalisis bahwa Nilai n adalah kelipatan 7 atau tidak
	if n%7 == 0 {
		fmt.Print(n)
		fmt.Println(" adalah kelipatan 7")
	} else {
		fmt.Print(n)
		fmt.Println(" bukanlah kelipatan 7")
	}

}
