package main

import "fmt"

// Fungsi Selection Sort untuk mengurutkan secara menaik (Ascending)
func selectionSort(arr []int) {
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

func main() {
	var n int
	// Membaca jumlah daerah
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var m int
		// Membaca jumlah rumah di daerah tersebut
		fmt.Scan(&m)

		rumah := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&rumah[j])
		}

		// Urutkan data
		selectionSort(rumah)

		// Cetak hasil sesuai format
		for j := 0; j < m; j++ {
			fmt.Print(rumah[j])
			if j < m-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}