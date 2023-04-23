package main

import "fmt"

func Ups(pilihan int) interface{} {
	switch pilihan {
	case 1:
		return 2
	case 2:
		return 'a'
	case 3:
		return "Ups"
	default:
		return "Eksdi"
	}
}

func main() {
	kosong1 := Ups(1)
	var kosong2 interface{} = Ups(3)

	fmt.Println(kosong1)
	fmt.Println(kosong2)
}
