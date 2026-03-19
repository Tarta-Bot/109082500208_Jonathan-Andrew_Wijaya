package main

import "fmt"

func hitungluaspersegipanjang(panjang int, lebar int) int {
	luas := panjang * lebar
	return luas
}

func main() {
	p := 10
	l := 5

	hasil := hitungluaspersegipanjang(p, l)

	fmt.Printf("Luas persegi panjang dengan panjang %d dan lebar %d adalah %d\n", p, l, hasil)
}