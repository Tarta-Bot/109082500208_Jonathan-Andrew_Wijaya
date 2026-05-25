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

	fmt.Println("Suara masuk:", masuk)
	fmt.Println("Suara sah:", sah)

	for i := 1; i <= 20; i++ {
		if hitung[i] > 0 {
			fmt.Printf("%d: %d\n", i, hitung[i])
		}
	}
}