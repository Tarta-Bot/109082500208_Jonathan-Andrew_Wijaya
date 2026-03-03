package main
import "fmt"

func main() {
	var tiket,  harga int
//kategori tiket
	fmt.Print("Masukkan kategori tiket (1, 2, atau 3): ")
	fmt.Scan(&tiket)
	switch tiket {
	case 1:
		harga = 100000
	case 2:
		harga = 75000
	case 3:
		harga = 50000
	default:
		fmt.Println("Kategori tiket tidak valid.")
		return
	}
	fmt.Printf("Harga tiket untuk kategori %d adalah: Rp%d\n", tiket, harga)
}
