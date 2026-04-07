package main

import "fmt"

func tanggunganHari(jumlahHari int, tujuan string) int {
	if tujuan == "domestik" {
		if jumlahHari > 3 {
			return 3
		}
		return jumlahHari
	} else if tujuan == "mancanegara" {
		if jumlahHari > 8 {
			return 8
		}
		return jumlahHari
	}
	return 0
}

func biayaPerHari(jumlahMhs int) int {
	biayaSatuMhs := 70000 + 250000 + 300000
	return jumlahMhs * biayaSatuMhs
}

func perhitunganBiaya(jumlahMhs, lamaPerjalanan int, tujuan string, totalBiaya *float64) {
	hariDitanggung := tanggunganHari(lamaPerjalanan, tujuan)
	biayaDasarPerHari := biayaPerHari(jumlahMhs)

	if tujuan == "domestik" {
		*totalBiaya = float64(hariDitanggung * biayaDasarPerHari)
	} else if tujuan == "mancanegara" {
		*totalBiaya = float64(hariDitanggung) * float64(biayaDasarPerHari) * 1.5
	}
}

func main() {
	var jumlah, lama int
	var tujuan string
	var biaya float64

	fmt.Print("Masukkan jumlah mahasiswa : ")
	fmt.Scan(&jumlah)
	fmt.Print("Masukkan lama hari study tour : ")
	fmt.Scan(&lama)
	fmt.Print("Masukkan tujuan study tour (domestik/mancanegara) : ")
	fmt.Scan(&tujuan)

	perhitunganBiaya(jumlah, lama, tujuan, &biaya)

	fmt.Printf("\nBiaya perjalanan yang harus dikeluarkan Tel-U : Rp. %.0f\n", biaya)
}	