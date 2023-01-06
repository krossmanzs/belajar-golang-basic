package main

import "fmt"

func main() {
	var angka1 int32 = 512
	var angka2 int64 = int64(512)

	fmt.Println(angka1)
	fmt.Println(angka2)
	fmt.Println(int8(angka1))

	nama := "Linux"
	index0 := nama[0]
	println(index0, "=", string(index0))
}
