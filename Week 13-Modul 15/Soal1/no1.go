package main

import "fmt"

const NMAX = 1000000 // jumlah maks data masukan

type arrInt [NMAX]int // tipe data alias array integer

func SelectionSort(T *arrInt, n int) {
	/* I.S. terdefinisi array T yang berisi sejumlah n bilangan bulat
	   F.S. array T terurut secara membesar berdasarkan algoritma selection sort */
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if T[j] < T[minIdx] {
				minIdx = j
			}
		}
		// Pertukaran nilai (swap)
		temp := T[i]
		T[i] = T[minIdx]
		T[minIdx] = temp
	}
}

func median(T arrInt, n int) float64 {
	/* mengembalikan median dari array T yang berisi sejumlah n bilangan bulat
	   yang telah terurut membesar */
	if n%2 == 1 {
		// Jika jumlah data ganjil, ambil nilai tengahnya
		return float64(T[n/2])
	} else {
		// Jika jumlah data genap, ambil rata-rata dari dua nilai tengah
		mid1 := T[(n/2)-1]
		mid2 := T[n/2]
		return float64(mid1+mid2) / 2.0
	}
}

func main() {
	var A arrInt    // array integer
	var x int       // variabel masukan
	var n int = 0   // jumlah data saat ini
	
	fmt.Println("Input data masukan :")
	fmt.Scan(&x)
	
	for x != -5313541 && n < NMAX {
		if x == 0 {
			if n > 0 {
				SelectionSort(&A, n)
				fmt.Println("Median :")
				fmt.Println(median(A, n))
			}
		} else {
			A[n] = x
			n++
		}
		fmt.Scan(&x)
	}
}