package main
import "fmt"
func main() {
	var target, tabungan, total, Bulan int

	fmt.Print("Masukkan target tabungan: ")
	fmt.Scan(&target)

	total = 0
	Bulan = 0

	for total < target {
		Bulan++
		fmt.Printf("Masukkan jumlah tabungan untuk bulan ke-%d: ", Bulan)
		fmt.Scan(&tabungan)
		
		total = total + tabungan
	}

	fmt.Printf("Selamat! Anda telah mencapai target tabungan dalam %d bulan.\n", Bulan)
	fmt.Printf("Total tabungan Anda: Rp%d\n", total)
}