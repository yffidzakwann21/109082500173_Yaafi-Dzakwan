package main

import "fmt"

// Batas maksimum array sesuai spesifikasi soal
const nMax = 7919

// Definisi Struct Buku
type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

// Definisi tipe data tipe array DaftarBuku
type DaftarBuku [nMax]Buku

// 1. Prosedur DaftarkanBuku (Membaca data dari input)
func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&pustaka[i].id, &pustaka[i].judul, &pustaka[i].penulis, &pustaka[i].penerbit)
		fmt.Scan(&pustaka[i].eksemplar, &pustaka[i].tahun, &pustaka[i].rating)
	}
}

// 2. Prosedur UrutBuku (Insertion Sort - Descending berdasarkan rating)
func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		key := pustaka[i]
		j := i - 1

		// Geser elemen yang ratingnya lebih kecil dari key.rating ke kanan
		for j >= 0 && pustaka[j].rating < key.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}
		pustaka[j+1] = key
	}
}

// 3. Prosedur CetakTerfavorit (Mencari rating tertinggi secara manual dari array acak)
func CetakTerfavorit(pustaka DaftarBuku, n int) {
	if n <= 0 {
		return
	}

	maxRating := pustaka[0].rating
	idxMax := 0
	for i := 1; i < n; i++ {
		if pustaka[i].rating > maxRating {
			maxRating = pustaka[i].rating
			idxMax = i
		}
	}

	b := pustaka[idxMax]
	fmt.Printf("%s %s %s %d\n", b.judul, b.penulis, b.penerbit, b.tahun)
}

// 4. Prosedur Cetak5Terbaru (Mencetak hingga 5 judul buku dengan rating tertinggi)
func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	limit := 5
	if n < 5 {
		limit = n
	}

	for i := 0; i < limit; i++ {
		fmt.Printf("%s\n", pustaka[i].judul)
	}
}

// 5. Prosedur CariBuku (Binary Search pada array yang terurut menurun)
func CariBuku(pustaka DaftarBuku, n int, r int) {
	left := 0
	right := n - 1
	foundIdx := -1

	for left <= right {
		mid := (left + right) / 2
		if pustaka[mid].rating == r {
			foundIdx = mid
			break
		} else if pustaka[mid].rating < r {
			right = mid - 1 // Cari ke kiri karena array terurut menurun
		} else {
			left = mid + 1  // Cari ke kanan
		}
	}

	if foundIdx != -1 {
		b := pustaka[foundIdx]
		fmt.Printf("%s %s %s %d %d %d\n", b.judul, b.penulis, b.penerbit, b.tahun, b.eksemplar, b.rating)
	} else {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

// Fungsi Utama: Mengatur alur eksekusi sesuai instruksi soal
func main() {
	var n int
	fmt.Scan(&n)

	var pustaka DaftarBuku

	// Membaca seluruh data buku
	DaftarkanBuku(&pustaka, n)

	// Membaca rating yang akan dicari pada baris terakhir input
	var ratingDicari int
	fmt.Scan(&ratingDicari)

	// --- PROSES OUTPUT ---
	
	// Baris 1 Output: Cetak buku terfavorit saat kondisi array belum terurut
	CetakTerfavorit(pustaka, n)

	// Proses pengurutan internal menggunakan Insertion Sort secara Descending
	UrutBuku(&pustaka, n)

	// Baris 2 Output: Cetak 5 besar judul buku rating tertinggi
	Cetak5Terbaru(pustaka, n)

	// Baris 3 Output: Pencarian biner berdasarkan rating yang diinput di akhir
	CariBuku(pustaka, n, ratingDicari)
}