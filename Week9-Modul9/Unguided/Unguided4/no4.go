package main

import (
	"bufio"
	"fmt"
	"os"
)

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var input string
	reader := bufio.NewReader(os.Stdin)

	*n = 0
	for *n < NMAX {
		fmt.Fscan(reader, &input)

		for _, ch := range input {
			if ch == '.' {
				return
			}
			t[*n] = ch
			*n++
		}
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(string(t[i]), " ")
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var n int

	fmt.Print("Masukkan karakter, akhiri dengan titik: ")
	isiArray(&tab, &n)

	fmt.Print("Isi array: ")
	cetakArray(tab, n)

	if palindrom(tab, n) {
		fmt.Println("Palindrome: true")
	} else {
		fmt.Println("Palindrome: false")
	}

	balikanArray(&tab, n)

	fmt.Print("Array setelah dibalik: ")
	cetakArray(tab, n)
}