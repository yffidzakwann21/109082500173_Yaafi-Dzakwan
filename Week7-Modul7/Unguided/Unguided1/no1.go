package main

import "fmt"

type suhu float64

func CelciusToReamur(celcius suhu) suhu {
	return celcius * 4 / 5
}

func CelciusToFahrenheit(celcius suhu) suhu {
	return (celcius * 9 / 5) + 32
}

func CelciusToKelvin(celcius suhu) suhu {
	return celcius + 273.15
}

func main() {
	var celcius suhu

	fmt.Println("=== KONVERTER CELCIUS ===")
	fmt.Print("Masukkan suhu (celcius) : ")
	fmt.Scan(&celcius)

	reamur := CelciusToReamur(celcius)
	fahrenheit := CelciusToFahrenheit(celcius)
	kelvin := CelciusToKelvin(celcius)

	fmt.Printf("\n%.0f celcius = %.1f reamur\n", celcius, reamur)
	fmt.Printf("%.0f celcius = %.1f fahrenheit\n", celcius, fahrenheit)
	fmt.Printf("%.0f celcius = %.2f kelvin\n", celcius, kelvin)
}