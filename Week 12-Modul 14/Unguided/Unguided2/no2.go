package main

import "fmt"

// Fungsi Selection Sort untuk mengurutkan secara menaik (Ascending)
func selectionSortAscending(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		if minIdx != i {
			arr[i], arr[minIdx] = arr[minIdx], arr[i]
		}
	}
}

// Fungsi Selection Sort untuk mengurutkan secara menurun (Descending)
func selectionSortDescending(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] > arr[maxIdx] {
				maxIdx = j
			}
		}
		if maxIdx != i {
			arr[i], arr[maxIdx] = arr[maxIdx], arr[i]
		}
	}
}

func main() {
	var n int
	// Membaca jumlah daerah
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var m int
		// Membaca jumlah rumah di daerah tersebut
		fmt.Scan(&m)

		var ganjil []int
		var genap []int

		for j := 0; j < m; j++ {
			var nomor int
			fmt.Scan(&nomor)
			// Pisahkan ganjil dan genap
			if nomor%2 != 0 {
				ganjil = append(ganjil, nomor)
			} else {
				genap = append(genap, nomor)
			}
		}

		// Urutkan masing-masing kelompok slice
		selectionSortAscending(ganjil)
		selectionSortDescending(genap)

		// Gabungkan saat pencetakan data
		first := true
		for _, num := range ganjil {
			if !first {
				fmt.Print(" ")
			}
			fmt.Print(num)
			first = false
		}
		for _, num := range genap {
			if !first {
				fmt.Print(" ")
			}
			fmt.Print(num)
			first = false
		}
		fmt.Println()
	}
}