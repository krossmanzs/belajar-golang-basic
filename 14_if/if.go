package main

import "fmt"

func main() {
	nama := "Linux"

	if nama == "Linux" {
		fmt.Println("Namanya adalah Linux")
	} else if nama == "linux" {
		fmt.Println("Namanya adalah linux")
	} else {
		fmt.Println("Namanya bukan Linux atau linux")
	}

	// if short statement
	if length := len(nama); length > 3 {
		fmt.Println("Panjang", nama, "lebih dari 3")
	} else {
		fmt.Println("Panjang", nama, "tidak lebih dari 3")
	}

}
