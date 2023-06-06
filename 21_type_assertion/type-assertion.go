package main

/*
	Type Assertions

	merupakan kemampuan merubah tipe data menjadi tipe data yang diinginkan
	Biasanya sering ditemukan pada tipe data interface kosong
*/

import "fmt"

func helloWorld() interface{} {
	return "Hello World"
}

func random() interface{} {
	return 1.23
}

func cek(apaAja interface{}) {
	switch value := apaAja.(type) {
	case int:
		fmt.Println("Integer", value)
	case string:
		fmt.Println("String", value)
	default:
		fmt.Println("Tipe data tidak diketahui")
	}
}

func main() {
	haloDunia := helloWorld()
	fmt.Println(haloDunia)

	cek(haloDunia)

	random := random()
	fmt.Println(random)

	cek(random)

	// haloJuga := haloDunia.(int) // panic
	// fmt.Println(haloJuga)

	
}