package main

import "fmt"

func cekGenap(angka int) bool {
	if angka%2 == 0 {
		return true
	}

	return false
}

func main() {
	angkates := 7

	hasilgenap := cekGenap(angkates)

	fmt.Printf("Apakah %d itu angka genap? jawabannya: %t\n", angkates, hasilgenap)
}