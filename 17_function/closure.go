package main

import "fmt"

func main() {
	/*
		Closure adalah kemampuan sebuah function berinteraksi dengan data-data
		disekitarnya dalam satu scope yang sama
	*/

	counter := 0

	increment := func() {
		fmt.Println("Melakukan increment")
		counter++
	}

	increment()
	increment()
	fmt.Println("Nilai counter:", counter)
}
