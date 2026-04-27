package main

import "fmt"

const kapasitas = 100

func main() {
	var klubA, klubB string
	var hasil [kapasitas]string
	var skorA, skorB int
	var jumlah int

	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)

	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	for {
		fmt.Print("Pertandingan ", jumlah+1, " : ")
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			hasil[jumlah] = klubA
		} else if skorB > skorA {
			hasil[jumlah] = klubB
		} else {
			hasil[jumlah] = "Draw"
		}

		jumlah++
	}

	for i := 0; i < jumlah; i++ {
		fmt.Println("Hasil", i+1, ":", hasil[i])
	}

	fmt.Println("Pertandingan selesai")
}