package main

import "fmt"

func endApp() {
	fmt.Println("Aplikasi selesai")
}

func startApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR!")
	}

	fmt.Println("Aplikasi berjalan normal")
}

func main() {
	startApp(true)
}
