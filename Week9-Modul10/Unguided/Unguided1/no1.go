package main

import "fmt"

func main() {
	var N int
	var berat [1000]float64

	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Scan(&berat[i])
	}

	terkecil := berat[0]
	terbesar := berat[0]

	for i := 1; i < N; i++ {
		if berat[i] < terkecil {
			terkecil = berat[i]
		}

		if berat[i] > terbesar {
			terbesar = berat[i]
		}
	}

	fmt.Println(terkecil, terbesar)
}