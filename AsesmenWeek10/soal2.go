package main

import "fmt"

const nMax = 51

type mahasiswa struct {
	NIM   string
	nama  string
	nilai int
}

type arrayMahasiswa [nMax]mahasiswa

func cariNilaiPertama(arr arrayMahasiswa, n int, nimCari string) int {
	for i := 0; i < n; i++ {
		if arr[i].NIM == nimCari {
			return arr[i].nilai 
		}
	}
	return -1 
}

func cariNilaiTerbesar(arr arrayMahasiswa, n int, nimCari string) int {
	var maks int = -1

	for i := 0; i < n; i++ {
		if arr[i].NIM == nimCari {
			if arr[i].nilai > maks {
				maks = arr[i].nilai 
			}
		}
	}
	return maks
}

func main() {
	var arr arrayMahasiswa
	var n int

	fmt.Print("Masukkan jumlah data : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&arr[i].NIM, &arr[i].nama, &arr[i].nilai)
	}

	var searchNIM string
	fmt.Print("Masukkan NIM mahasiswa yang ingin dicari nilai pertama dan nilai terbesarnya : ")
	fmt.Scan(&searchNIM)

	var nilaiPertama int = cariNilaiPertama(arr, n, searchNIM)
	var nilaiTerbesar int = cariNilaiTerbesar(arr, n, searchNIM)

	if nilaiPertama != -1 {
		fmt.Printf("Nilai pertama dari NIM %s adalah %d\n", searchNIM, nilaiPertama)
		fmt.Printf("Nilai terbesar dari NIM %s adalah %d\n", searchNIM, nilaiTerbesar)
	}
}