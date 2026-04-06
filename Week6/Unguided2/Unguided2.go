package main

import "fmt"

func faktor(i int, n int) {
	if i <= n {
		if n%i == 0 {
			fmt.Print(i, " ")
		}
		faktor(i+1, n)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	faktor(1, n)
	fmt.Print("\n")
}