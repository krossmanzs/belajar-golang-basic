package main

import "fmt"

func main() {
	// for biasa
	var counter int = 0
	for counter <= 10 {
		fmt.Println("Ini adalah perulangan ke", counter)
		counter++
	}

	fmt.Println()
	// for dengan statement
	for i := 1; i <= 10; i++ {
		if i == 3 {
			continue
		} else if i == 9 {
			break
		} else {
			fmt.Println("Perulangan ke", i)
		}
	}

	fmt.Println()
	// for range
	// misal ada sebuah slice
	names := []string{"Cornelius", "Linux", "Selebew", "Uwu"}

	// jika manual
	for index := 0; index < len(names); index++ {
		fmt.Println(names[index])
	}

	fmt.Println()
	// jika menggunakan for range
	for index, name := range names { // penulisan index tidak wajib atau bisa diganti _
		fmt.Println("Index:", index, "Name:", name)
	}

	person := make(map[string]string)
	person["name"] = "Cornelius Linux"
	person["age"] = "18"

	for key, value := range person {
		fmt.Println("Key:", key, "Value:", value)
	}
}
