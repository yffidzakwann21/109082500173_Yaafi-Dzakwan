package main

import "fmt"

// Fungsi Insertion Sort untuk mengurutkan secara menaik (Ascending)
func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		// Geser elemen-elemen yang lebih besar dari key ke kanan
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func main() {
	var num int
	var data []int

	// Loop untuk membaca input sampai ditemukan bilangan negatif
	for {
		fmt.Scan(&num)
		if num < 0 {
			break
		}
		data = append(data, num)
	}

	// Jika input kosong, program langsung selesai
	if len(data) == 0 {
		return
	}

	// 1. Urutkan data menggunakan Insertion Sort
	insertionSort(data)

	// 2. Cetak baris pertama: isi array yang sudah terurut
	for i := 0; i < len(data); i++ {
		fmt.Print(data[i])
		if i < len(data)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()

	// 3. Periksa jarak antar elemen
	// Jika elemen kurang dari atau sama dengan 1, otomatis dianggap berjarak tetap 0 atau tidak punya selisih
	if len(data) <= 1 {
		fmt.Println("Data berjarak 0")
		return
	}

	// Ambil selisih pertama sebagai acuan awal
	selisihAwal := data[1] - data[0]
	isTetap := true

	// Periksa seluruh sisa elemen bersebelahan
	for i := 1; i < len(data)-1; i++ {
		if data[i+1]-data[i] != selisihAwal {
			isTetap = false
			break
		}
	}

	// 4. Cetak baris kedua: status jarak
	if isTetap {
		fmt.Printf("Data berjarak %d\n", selisihAwal)
	} else {
		fmt.Println("data berjarak tidak tetap")
	}
}