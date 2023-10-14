package main

import "fmt"

// TUJUAN: mencaritahu cara kerjanya MAP

func main() {
	// map
	// banyak data
	// key value
	// dinamis
	// key unique

	// januari => 10 juta
	// februari => 20 juta

	// deklarasi Map, dan string []
	var dataPenjualan = make(map[string]int)
	dataPenjualan["januari"] = 10000
	dataPenjualan["februari"] = 20000
	fmt.Println(dataPenjualan)

	// coba utak atik sendiri
	var coba
}
