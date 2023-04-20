package main

import "fmt"

func main() {
	kode := 6

	switch kode {
	case 1:
		fmt.Println("Halo juga")
	case 2:
		fmt.Println("Selamat datang")
	case 3:
		fmt.Println("Ini kamar anda")
	default:
		fmt.Println("Ini default")
	}

	// juga mendukung short statement
	nama := "Linux"

	switch length := len(nama); length > 3 {
	case true:
		fmt.Println("Panjang", nama, "lebih dari 3 huruf")
	case false:
		fmt.Println("Panjang", nama, "tidak lebih dari 3 huruf")
	}

	length := len(nama)

	// switch tanpa kondisi
	switch {
	case length > 3:
		fmt.Println("Panjang", nama, "lebih dari 3 huruf")
	default:
		fmt.Println("Panjang", nama, "tidak lebih dari 3 huruf")
	}
}
