package main

import "fmt"

type Titik struct {
	x int
	y int
}

type Lingkaran struct {
	pusat Titik
	r     int
}

func diDalamLingkaran(p Titik, l Lingkaran) bool {
	dx := p.x - l.pusat.x
	dy := p.y - l.pusat.y

	return dx*dx+dy*dy <= l.r*l.r
}

func main() {
	var l1, l2 Lingkaran
	var p Titik

	fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.r)
	fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.r)
	fmt.Scan(&p.x, &p.y)

	dalam1 := diDalamLingkaran(p, l1)
	dalam2 := diDalamLingkaran(p, l2)

	if dalam1 && dalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}