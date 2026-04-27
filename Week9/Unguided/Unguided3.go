package main

import "fmt"

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var hasil [100]string
	var jumlahPertandingan int

	fmt.Print("Klub A: ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scan(&klubB)

	pertandinganKe := 1
	for {
		fmt.Printf("Pertandingan %d: ", pertandinganKe)
		fmt.Scan(&skorA, &skorB)

		// Berhenti jika skor salah satu atau keduanya negatif
		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			hasil[jumlahPertandingan] = klubA
		} else if skorB > skorA {
			hasil[jumlahPertandingan] = klubB
		} else {
			hasil[jumlahPertandingan] = "Draw"
		}

		jumlahPertandingan++
		pertandinganKe++
	}

	for i := 0; i < jumlahPertandingan; i++ {
		fmt.Printf("Hasil %d: %s\n", i+1, hasil[i])
	}
	fmt.Println("Pertandingan selesai")
}