package main

import "fmt"

func main() {
	var name string = "Reyhan Kunci Inggris"
	var height = 167.2
	var age byte

	fmt.Println(name)
	fmt.Println(height)

	name = "Sucipto Kebab"
	age = 5

	fmt.Println(name)
	fmt.Println(age)

	name = "Cornelius Linux"

	fmt.Println(name)
	fmt.Println(age)

	name = "Cornelius Linux"
	fmt.Println(name)

	// dapat menggunakan := untuk deklarasi tanpa menggunakan var

	// country := "string"
	country := string("Indonesia")
	fmt.Println(country)

	// membuat variabel sekaligus banyak (supaya readable)
	var (
		firstName = string("Cornelius")
		lastName  = "Linux"
	)

	fmt.Println(firstName, lastName)
}
