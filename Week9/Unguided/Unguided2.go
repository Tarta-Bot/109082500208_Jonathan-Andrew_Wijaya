package main

import (
	"fmt"
	"math"
)

const NMAX int = 100

func main() {
	var arr [NMAX]int
	var n int

	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&n)

	fmt.Printf("Masukkan %d bilangan bulat:\n", n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	// a. Menampilkan keseluruhan isi array
	fmt.Print("Isi array keseluruhan: ")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	// b. Indeks ganjil saja
	fmt.Print("Elemen indeks ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	// c. Indeks genap saja (0 adalah genap)
	fmt.Print("Elemen indeks genap: ")
	for i := 0; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	// d. Indeks kelipatan x
	var x int
	fmt.Print("Masukkan kelipatan indeks (x): ")
	fmt.Scan(&x)
	fmt.Printf("Elemen indeks kelipatan %d: ", x)
	for i := 0; i < n; i++ {
		if x != 0 && i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	// e. Menghapus elemen pada indeks tertentu
	var hapusIdx int
	fmt.Print("Masukkan indeks elemen yang akan dihapus: ")
	fmt.Scan(&hapusIdx)
	if hapusIdx >= 0 && hapusIdx < n {
		for i := hapusIdx; i < n-1; i++ {
			arr[i] = arr[i+1]
		}
		n--
		fmt.Print("Isi array setelah penghapusan: ")
		for i := 0; i < n; i++ {
			fmt.Print(arr[i], " ")
		}
		fmt.Println()
	}

	// f. Rata-rata
	var sum float64
	for i := 0; i < n; i++ {
		sum += float64(arr[i])
	}
	avg := sum / float64(n)
	fmt.Printf("Rata-rata: %.2f\n", avg)

	// g. Standar deviasi
	var sumSq float64
	for i := 0; i < n; i++ {
		sumSq += math.Pow(float64(arr[i])-avg, 2)
	}
	stdDev := math.Sqrt(sumSq / float64(n))
	fmt.Printf("Standar Deviasi: %.2f\n", stdDev)

	// h. Frekuensi
	var cari, freq int
	fmt.Print("Masukkan bilangan untuk dicari frekuensinya: ")
	fmt.Scan(&cari)
	for i := 0; i < n; i++ {
		if arr[i] == cari {
			freq++
		}
	}
	fmt.Printf("Frekuensi kemunculan %d adalah %d kali\n", cari, freq)
}