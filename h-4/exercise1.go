package main

import "fmt"

type Car struct {
	tipe   string
	fuelIn int
}

// Counting Fuel Cost:
// https://www.transpoco.com/knowledge/how-is-fuel-consumption-calculated

/* Cara menghitung konsumsi BBM adalah membagi jarak tempuh dengan jumlah konsumsi BBM. contohnya
adalah jarak tempuh 120 km menghabiskan 2 liter BBM, maka konsumsi per liter: 120 km/2 liter = 60 km/liter.
Konsumsi per kilometer : 2 liter/120 km = 0,017 liter/km. */
// https://otoklix.com/blog/cara-menghitung-konsumsi-bbm/#:~:text=Cara%20menghitung%20konsumsi%20BBM%20adalah,km%20%3D%200%2C017%20liter%2Fkm.

// 1km sama dengan 0.6214 mil. Mil
// https://www.metric-conversions.org/id/panjang/kilometer-ke-mil.htm

func main() {
	carSedan := Car{"sedan", 1500}
	var konsumsi_liter double
	//car.fuelIn = 1500
	//car.tipe = "Sedan" //contoh mobil sedan fuelnya 1,5L
	// Input Jarak
	var jarak int
	fmt.Print("Input Jarak per Mil: ")
	fmt.Scanf("%d", &jarak)
	fmt.Println("Jarak: ", jarak)

	// Hitung Jarak Liter/Mil (100km * 0.6214 mil = 62.14 per mil)
	konsumsi_liter := jarak / 62.14(float64)
	fmt.Println("Konsumsi Liter/Mil: ", konsumsi_liter)
	// Output Jarak
	//fmt.Printf("%d km\n", jarak/3.6)

}
