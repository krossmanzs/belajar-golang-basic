package main

import "fmt"

func main() {
	var a = 9
	var b = 10
	var jumlah = a + b
	var pengurangan = b - a
	var perkalian = a * b
	var pembagian = b / a
	var modulus = b % a

	fmt.Println("Penjumlahan:", jumlah)
	fmt.Println("Pengurangan:", pengurangan)
	fmt.Println("Perkalian:", perkalian)
	fmt.Println("Pembagian:", pembagian)
	fmt.Println("Modulus:", modulus)
}
