package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var char rune
	*n = 0
	for {
		fmt.Scanf("%c", &char)
		// Proses berhenti jika mendeteksi titik atau newline/batas maksimal
		if char == '.' || char == '\n' || *n >= NMAX {
			break
		}
		(*t)[*n] = char
		*n++
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		// Menukar elemen depan dengan elemen belakang
		(*t)[i], (*t)[n-1-i] = (*t)[n-1-i], (*t)[i]
	}
}

func palindrom(t tabel, n int) bool {
	tBalikan := t
	balikanArray(&tBalikan, n)
	
	for i := 0; i < n; i++ {
		if t[i] != tBalikan[i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	fmt.Print("Masukkan teks (akhiri dengan . atau enter): ")
	isiArray(&tab, &m)

	fmt.Print("Teks asli: ")
	cetakArray(tab, m)

	tabBalikan := tab
	balikanArray(&tabBalikan, m)
	fmt.Print("Reverse teks: ")
	cetakArray(tabBalikan, m)

	isPalindrom := palindrom(tab, m)
	fmt.Printf("Apakah palindrom? %t\n", isPalindrom)
}