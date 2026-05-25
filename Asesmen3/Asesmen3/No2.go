package main

import "fmt"

type Pemain struct {
	namaDepan    string
	namaBelakang string
	gol          int
	assist       int
}

func main() {
	var n int
	fmt.Scan(&n)

	pemain := make([]Pemain, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&pemain[i].namaDepan, &pemain[i].namaBelakang, &pemain[i].gol, &pemain[i].assist)
	}

	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if pemain[j].gol > pemain[maxIdx].gol || (pemain[j].gol == pemain[maxIdx].gol && pemain[j].assist > pemain[maxIdx].assist) {
				maxIdx = j
			}
		}
		pemain[i], pemain[maxIdx] = pemain[maxIdx], pemain[i]
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%s %s %d %d\n", pemain[i].namaDepan, pemain[i].namaBelakang, pemain[i].gol, pemain[i].assist)
	}
}