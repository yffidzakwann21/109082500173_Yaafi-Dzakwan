package main

import "fmt"

const nMax int = 51

type mahasiswa struct {
	NIM   string
	nama  string
	nilai int
}

type arrayMahasiswa [nMax]mahasiswa

func inputData(T *arrayMahasiswa, N *int) {
	fmt.Print("Masukkan jumlah data : ")
	fmt.Scan(N)

	for i := 0; i < *N; i++ {
		fmt.Print("Masukkan data ke-", i+1, " : ")
		fmt.Scan(&T[i].NIM, &T[i].nama, &T[i].nilai)
	}
}

func nilaiPertama(T arrayMahasiswa, N int, nim string) int {
	for i := 0; i < N; i++ {
		if T[i].NIM == nim {
			return T[i].nilai
		}
	}
	return -1
}

func nilaiTerbesar(T arrayMahasiswa, N int, nim string) int {
	var max int
	var ketemu bool

	ketemu = false

	for i := 0; i < N; i++ {
		if T[i].NIM == nim {
			if ketemu == false {
				max = T[i].nilai
				ketemu = true
			} else if T[i].nilai > max {
				max = T[i].nilai
			}
		}
	}

	if ketemu == false {
		return -1
	}
	return max
}

func main() {
	var data arrayMahasiswa
	var N int
	var cari string
	var pertama, terbesar int

	inputData(&data, &N)

	fmt.Print("Masukkan NIM mahasiswa yang ingin dicari nilai pertama dan nilai terbesarnya : ")
	fmt.Scan(&cari)

	pertama = nilaiPertama(data, N, cari)
	terbesar = nilaiTerbesar(data, N, cari)

	fmt.Println("Nilai pertama dari NIM", cari, "adalah", pertama)
	fmt.Println("Nilai terbesar dari NIM", cari, "adalah", terbesar)
}