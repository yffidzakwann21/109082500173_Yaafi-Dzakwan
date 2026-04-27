package main

import (
	"fmt"
	"math"
)

const kapasitas = 100

func tampilArray(arr [kapasitas]int, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

func main() {
	var arr [kapasitas]int
	var n int

	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Print("Elemen ke-", i, ": ")
		fmt.Scan(&arr[i])
	}

	fmt.Println("\nSeluruh isi array:")
	tampilArray(arr, n)

	fmt.Println("Elemen dengan indeks ganjil:")
	for i := 1; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Println("Elemen dengan indeks genap:")
	for i := 0; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	var x int
	fmt.Print("Masukkan kelipatan indeks x: ")
	fmt.Scan(&x)

	fmt.Println("Elemen dengan indeks kelipatan", x, ":")
	if x != 0 {
		for i := 0; i < n; i++ {
			if i%x == 0 {
				fmt.Print(arr[i], " ")
			}
		}
	}
	fmt.Println()

	var hapus int
	fmt.Print("Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&hapus)

	for i := hapus; i < n-1; i++ {
		arr[i] = arr[i+1]
	}
	n--

	fmt.Println("Array setelah penghapusan:")
	tampilArray(arr, n)

	var total int
	for i := 0; i < n; i++ {
		total += arr[i]
	}

	rata := float64(total) / float64(n)
	fmt.Println("Rata-rata:", rata)

	var jumlahKuadrat float64
	for i := 0; i < n; i++ {
		selisih := float64(arr[i]) - rata
		jumlahKuadrat += selisih * selisih
	}

	standarDeviasi := math.Sqrt(jumlahKuadrat / float64(n))
	fmt.Println("Standar deviasi:", standarDeviasi)

	var cari int
	fmt.Print("Masukkan bilangan yang ingin dicari frekuensinya: ")
	fmt.Scan(&cari)

	frekuensi := 0
	for i := 0; i < n; i++ {
		if arr[i] == cari {
			frekuensi++
		}
	}

	fmt.Println("Frekuensi", cari, "adalah", frekuensi)
}