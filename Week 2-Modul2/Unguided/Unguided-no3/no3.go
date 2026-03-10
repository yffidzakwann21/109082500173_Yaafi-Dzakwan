package main
import "fmt"
func main() {
	var berat int

	fmt.Print("Masukkan total berat (gram): ")
	fmt.Scan(&berat)

	kg := berat / 1000
	sisa := berat % 1000

	biayaKg := kg * 10000
	biayaTambahan := 0

	if sisa > 500 {
		biayaTambahan = 5 * sisa
	} else {
		if kg < 10 {
			biayaTambahan = 15 * sisa
		}
	}

	total := biayaKg + biayaTambahan

	fmt.Println("===== Detail Perhitungan =====")
	fmt.Println("Detail berat :", kg, "kg +", sisa, "gram")
	fmt.Println("Detail biaya : Rp.", biayaKg, "+ Rp.", biayaTambahan)
	fmt.Println("Total biaya: Rp", total)
}