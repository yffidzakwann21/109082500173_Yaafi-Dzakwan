package main

import "fmt"

func main() {
	var x, y int
	var berat [1000]float64
	var totalWadah [1000]float64

	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}

	jumlahWadah := 0

	for i := 0; i < x; i += y {
		total := 0.0

		for j := i; j < i+y && j < x; j++ {
			total += berat[j]
		}

		totalWadah[jumlahWadah] = total
		jumlahWadah++
	}

	totalSemuaWadah := 0.0

	for i := 0; i < jumlahWadah; i++ {
		fmt.Print(totalWadah[i])

		if i < jumlahWadah-1 {
			fmt.Print(" ")
		}

		totalSemuaWadah += totalWadah[i]
	}

	rataRata := totalSemuaWadah / float64(jumlahWadah)

	fmt.Println()
	fmt.Println(rataRata)
}