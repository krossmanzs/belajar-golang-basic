package main

import "fmt"

func endApp() {
	fmt.Println("Aplikasi selesai dijalankan")
	message := recover()
	if message != nil {
		fmt.Println("Error message:", message)
	}
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("TERJADI ERROR")
	}

	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(true)
}
