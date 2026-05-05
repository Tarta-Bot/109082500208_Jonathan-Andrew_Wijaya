package main

import "fmt"

const nProv = 10

type NamaProv [nProv]string
type PopProv [nProv]int
type TumbuhProv [nProv]float64

func InputData(prov *NamaProv, pop *PopProv, tumbuh *TumbuhProv) {
	fmt.Println("--- Masukkan Nama Provinsi, Populasi Provinsi, Angka Pertumbuhan Provinsi ---")
	for i := 0; i < nProv; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&(*prov)[i], &(*pop)[i], &(*tumbuh)[i])
	}
}

func ProvinsiTercepat(tumbuh TumbuhProv) int {
	var maxIdx int = 0
	var maxVal float64 = tumbuh[0]

	for i := 1; i < nProv; i++ {
		if tumbuh[i] > maxVal {
			maxVal = tumbuh[i]
			maxIdx = i
		}
	}
	return maxIdx
}

func IndeksProvinsi(prov NamaProv, nama string) int {
	for i := 0; i < nProv; i++ {
		if prov[i] == nama {
			return i
		}
	}
	return -1
}

func Prediksi(prov NamaProv, pop PopProv, tumbuh TumbuhProv) {
	fmt.Println("\n--- Prediksi Jumlah Penduduk Tahun Depan Pada Provinsi Dengan Pertumbuhan Diatas 2% ---")
	for i := 0; i < nProv; i++ {
		if tumbuh[i] > 0.02 {
			var tambahan float64 = float64(pop[i]) * tumbuh[i]
			var prediksi float64 = float64(pop[i]) + tambahan
			
			fmt.Printf("%s %.0f\n", prov[i], prediksi)
		}
	}
}

func main() {
	var prov NamaProv
	var pop PopProv
	var tumbuh TumbuhProv

	InputData(&prov, &pop, &tumbuh)

	var searchProv string
	fmt.Scan(&searchProv)

	var idxTercepat int = ProvinsiTercepat(tumbuh)
	fmt.Printf("\nProvinsi dengan angka pertumbuhan tercepat : %s\n", prov[idxTercepat])
	
	var idxCari int = IndeksProvinsi(prov, searchProv)
	if idxCari != -1 {
		fmt.Printf("\nData provinsi yang dicari : %s\n", prov[idxCari])
	}

	Prediksi(prov, pop, tumbuh)
}