package main

import "fmt"

const nProv int = 10

type NamaProv [nProv]string
type PopProv [nProv]int
type TumbuhProv [nProv]float64

func inputData(prov *NamaProv, pop *PopProv, tumbuh *TumbuhProv) {
	for i := 0; i < nProv; i++ {
		fmt.Print("Masukkan data ke-", i+1, " : ")
		fmt.Scan(&prov[i], &pop[i], &tumbuh[i])
	}
}

func provinsiTercepat(tumbuh TumbuhProv) int {
	idx := 0

	for i := 1; i < nProv; i++ {
		if tumbuh[i] > tumbuh[idx] {
			idx = i
		}
	}

	return idx
}

func prediksi(prov NamaProv, pop PopProv, tumbuh TumbuhProv) {
	fmt.Println("=== Prediksi Jumlah Penduduk Tahun Depan Pada Provinsi Dengan Pertumbuhan Diatas 2% ===")

	for i := 0; i < nProv; i++ {
		if tumbuh[i] > 0.02 {
			hasil := float64(pop[i]) * (1 + tumbuh[i])
			fmt.Printf("%s %.0f\n", prov[i], hasil)
		}
	}
}

func indeksProvinsi(prov NamaProv, nama string) int {
	for i := 0; i < nProv; i++ {
		if prov[i] == nama {
			return i + 1
		}
	}

	return -1
}

func main() {
	var prov NamaProv
	var pop PopProv
	var tumbuh TumbuhProv
	var nama string
	var idxCepat, idxCari int

	fmt.Println("=== Masukkan Nama Provinsi, Populasi Provinsi, Angka Pertumbuhan Provinsi ===")

	inputData(&prov, &pop, &tumbuh)

	fmt.Print("Masukkan data ke-11 : ")
	fmt.Scan(&nama)

	idxCepat = provinsiTercepat(tumbuh)
	idxCari = indeksProvinsi(prov, nama)

	fmt.Println()
	fmt.Println("Provinsi dengan angka pertumbuhan tercepat :", prov[idxCepat])
	fmt.Println()
	fmt.Println("Data provinsi yang dicari :", nama, idxCari)
	fmt.Println()

	prediksi(prov, pop, tumbuh)
}