package main

import "fmt"

func ganjil(awal int, n int) {
	if awal <= n {
		fmt.Print(awal, " ")
		ganjil(awal+2, n)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	ganjil(1, n)
	fmt.Print("\n")
}