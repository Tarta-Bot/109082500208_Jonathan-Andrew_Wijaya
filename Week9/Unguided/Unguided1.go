package main

import (
	"fmt"
	"math"
)

type titik struct {
	x, y float64
}

type lingkaran struct {
	pusat  titik
	radius float64
}

func jarak(p, q titik) float64 {
	return math.Sqrt(math.Pow(p.x-q.x, 2) + math.Pow(p.y-q.y, 2))
}

func didalam(c lingkaran, p titik) bool {
	return jarak(c.pusat, p) <= c.radius
}

func main() {
	var l1, l2 lingkaran
	var p titik

	// Masukan untuk lingkaran 1 (cx, cy, r)
	fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.radius)
	// Masukan untuk lingkaran 2 (cx, cy, r)
	fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.radius)
	// Masukan untuk titik sembarang (x, y)
	fmt.Scan(&p.x, &p.y)

	inL1 := didalam(l1, p)
	inL2 := didalam(l2, p)

	if inL1 && inL2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if inL1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if inL2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}