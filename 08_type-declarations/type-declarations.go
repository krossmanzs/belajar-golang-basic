package main

import "fmt"

/*
	Type Declarations adalah kemampuan membuat ulang tipe
		data baru dari tipe data yang sudah ada
	Type Declarations biasanya digunakan untuk membuat alias terhadap
		tipe data yang sudah ada, dengan tujuan agar lebih mudah dimengerti
	Kalau di C seperti define
*/

func main() {
	type NIM string
	type Married bool

	var NIMLinux NIM = "122548975"
	var marriedStatus Married = true

	fmt.Println("Married:", marriedStatus)
	fmt.Println(NIMLinux)
	fmt.Println(NIM("122140079"))
}
