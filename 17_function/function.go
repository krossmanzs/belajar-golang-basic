package main

import "fmt"

// variadic/varargs function
func printNumbers(numbers ...int) {
	for _, number := range numbers {
		fmt.Println(number)
	}
}

// return named values
func getStudentData() (name, nim string, age int) {
	name = "Cornelius Linux"
	nim = "122140079"
	age = 18

	// secara eksplisit
	// return name, nim, age

	// atau bisa begini saja
	return
}

// returning multiple values
func getFullName() (string, string, int) {
	return "Cornelius", "Linux", 18
}

func penjumlahan(angka1, angka2 int) int {
	return angka1 + angka2
}

func sayHelloTo(name string) {
	fmt.Println("Halo", name)
}

func helloWorld() {
	fmt.Println("Hello World")
}

func main() {
	helloWorld()
	sayHelloTo("Linux")
	fmt.Println("Hasil penjumlahan 5 + 2:", penjumlahan(5, 2))

	firstName, lastName, _ := getFullName()

	fmt.Println(firstName)
	fmt.Println(lastName)

	a, _, c := getStudentData()

	fmt.Println(a)
	fmt.Println(c)
	fmt.Println()

	printNumbers(1, 2, 3, 4, 12, 343, 535, 32, 123, 42)

	fmt.Println()

	numbers := []int{2, 2, 1, 3, 4, 1, 5, 2, 5, 6}
	printNumbers(numbers...)
}
