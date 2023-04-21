package main

import "fmt"

func endApplication() {
	fmt.Println("Application successfully stopped")
}

func runApplication(number int) {
	defer endApplication()
	divide := 10 / number
	fmt.Println("Divide value:", divide)
	fmt.Println("Running application...")
}

func main() {
	runApplication(0)
}
