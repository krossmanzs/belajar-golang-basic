package main

import "fmt"

func main() {
	// const tidak bisa menggunakan :=
	const firstName string = "Cornelius"
	const lastName string = "Linux"
	const age byte = 17

	fmt.Println(firstName, lastName)

	// multiple const variable declaration
	const (
		anak1 string = "Reyhan"
		anak2 string = "Kebab"
	)

	fmt.Println(anak1, anak2)
}
