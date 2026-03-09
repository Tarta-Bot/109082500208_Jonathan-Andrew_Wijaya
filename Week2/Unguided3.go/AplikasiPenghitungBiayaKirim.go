package main

import (
	"fmt"
)

func main() {
	var totalBeratGram int

	fmt.Print("Masukkan total berat (gram): ")
	fmt.Scan(&totalBeratGram)

	beratKg := totalBeratGram / 1000
	sisaGram := totalBeratGram % 1000

	biayaDasar := beratKg * 10000

	var biayaTambahan int

	if beratKg > 10 {
		biayaTambahan = 0
	} else if sisaGram >= 500 {
		biayaTambahan = sisaGram * 5
	} else {
		biayaTambahan = sisaGram * 15
	}

	totalBiaya := biayaDasar + biayaTambahan

	fmt.Println("\n===== Detail Perhitungan =====")
	fmt.Printf("Detail berat : %d kg + %d gram\n", beratKg, sisaGram)
	fmt.Printf("Detail biaya : Rp. %d + Rp. %d\n", biayaDasar, biayaTambahan)
	fmt.Printf("Total biaya: Rp %d\n", totalBiaya)
}	