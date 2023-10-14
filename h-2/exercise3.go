package main

import "fmt"

// 3. buatlah sebuah program untuk menghitung luas trapesium
func main() {
	// Deklarasi nilai trapesium
	var a, b, t int
	// Input nilai trapesium
	fmt.Print("Input nilai a: ")
	fmt.Scanln(&a)
	fmt.Print("Input nilai b: ")
	fmt.Scanln(&b)
	fmt.Print("Input nilai t: ")
	fmt.Scanln(&t)

	// Hitung trapesium
	L := 0.5 * float64(a*b) * float64(t)

	// Tampilkan hasil Trapesium
	fmt.Print("Luas Trapesium: ")
	fmt.Println(L)

}
