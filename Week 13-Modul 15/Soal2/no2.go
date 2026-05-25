package main

import "fmt"

const NMAX = 1001

type Pemain struct {
	namaDepan    string
	namaBelakang string
	jumlahGol    int
	jumlahAssist int
}

type DaftarPemain [NMAX]Pemain

// ====== 1. FUNGSI SORTING (Insertion Sort - Menurun) ======
func InsertionSort(T *DaftarPemain, n int) {
	for i := 1; i < n; i++ {
		key := T[i]
		j := i - 1

		// Urutkan menurun berdasarkan Gol, jika Gol sama urutkan berdasarkan Assist
		for j >= 0 && (T[j].jumlahGol < key.jumlahGol || (T[j].jumlahGol == key.jumlahGol && T[j].jumlahAssist < key.jumlahAssist)) {
			T[j+1] = T[j]
			j--
		}
		T[j+1] = key
	}
}

// ====== 2. FUNGSI SEARCHING (Linear Search) ======
// Perbaikan pada baris 32: Tipe data ditulis dengan benar (depan string, belakang string)
func CariPemain(T DaftarPemain, n int, depan string, belakang string) int {
	for i := 0; i < n; i++ {
		// Mencocokkan nama depan dan nama belakang
		if T[i].namaDepan == depan && T[i].namaBelakang == belakang {
			return i // Mengembalikan indeks jika ditemukan
		}
	}
	return -1 // Mengembalikan -1 jika tidak ditemukan
}

func main() {
	var T DaftarPemain
	var n int

	fmt.Println("Masukkan Data Input :")
	fmt.Scanln(&n)

	if n > NMAX {
		n = NMAX
	}

	// Membaca input data pemain langsung menggunakan fmt.Scan
	for i := 0; i < n; i++ {
		fmt.Scan(&T[i].namaDepan, &T[i].namaBelakang, &T[i].jumlahGol, &T[i].jumlahAssist)
	}

	// Proses Pengurutan (Sorting)
	InsertionSort(&T, n)

	// Menampilkan Hasil Pengurutan
	fmt.Println("\nHasil Sorting :")
	for i := 0; i < n; i++ {
		fmt.Printf("%s %s %d %d\n", T[i].namaDepan, T[i].namaBelakang, T[i].jumlahGol, T[i].jumlahAssist)
	}

	// Proses Pencarian (Searching)
	var cariDepan, cariBelakang string
	fmt.Println("\n=================================")
	fmt.Print("Masukkan nama pemain yang dicari (NamaDepan NamaBelakang): ")
	fmt.Scan(&cariDepan, &cariBelakang)

	indeks := CariPemain(T, n, cariDepan, cariBelakang)
	if indeks != -1 {
		fmt.Printf("\nData Ditemukan! Pemain berada di Peringkat: %d\n", indeks+1)
	} else {
		fmt.Println("\nPemain tidak ditemukan.")
	}
}