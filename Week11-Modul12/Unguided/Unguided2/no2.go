package main

import "fmt"

func main() {
	var x int
	var suara [21]int
	total := 0
	sah := 0

	for {
		fmt.Scan(&x)

		if x == 0 {
			break
		}

		total++

		if x >= 1 && x <= 20 {
			sah++
			suara[x]++
		}
	}

	ketua := 1
	wakil := 1

	for i := 2; i <= 20; i++ {
		if suara[i] > suara[ketua] {
			wakil = ketua
			ketua = i
		} else if i != ketua {
			if suara[i] > suara[wakil] || wakil == ketua {
				wakil = i
			}
		}
	}

	// Jika wakil ternyata lebih besar dari ketua
	if suara[wakil] > suara[ketua] {
		ketua, wakil = wakil, ketua
	}

	fmt.Println("Suara masuk:", total)
	fmt.Println("Suara sah:", sah)
	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)
}