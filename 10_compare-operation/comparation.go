package main

import "fmt"

func main() {
	var a int = 5
	var b = 6
	var kurangDari bool = a < b
	var lebihDari bool = a > b
	var kurangDariSamaDengan = a <= b
	var lebihDariSamaDengan = a >= b
	var samaDengan = a == b
	var tidakSamaDengan = a != b

	fmt.Println("Kurang dari:", kurangDari)
	fmt.Println("Lebih dari:", lebihDari)
	fmt.Println("Kurang dari sama dengan:", kurangDariSamaDengan)
	fmt.Println("Lebih dari sama dengan:", lebihDariSamaDengan)
	fmt.Println("Sama dengan: ", samaDengan)
	fmt.Println("Tidak sama dengan: ", tidakSamaDengan)
}
