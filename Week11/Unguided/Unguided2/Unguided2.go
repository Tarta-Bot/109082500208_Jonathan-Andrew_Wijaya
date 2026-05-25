package main

import "fmt"

func main() {
	var suara int
	var masuk, sah int
	var hitung [21]int

	for {
		fmt.Scan(&suara)
		if suara == 0 {
			break
		}

		masuk++

		if suara >= 1 && suara <= 20 {
			sah++
			hitung[suara]++
		}
	}

	ketua := 0
	wakil := 0

	for i := 1; i <= 20; i++ {
		if hitung[i] > hitung[ketua] {
			wakil = ketua
			ketua = i
		} else if hitung[i] > hitung[wakil] && i != ketua {
			wakil = i
		}
	}

	fmt.Println("Suara masuk:", masuk)
	fmt.Println("Suara sah:", sah)
	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)
}