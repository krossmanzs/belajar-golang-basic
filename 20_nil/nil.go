package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			name: name,
		}
	}
}

func main() {
	var person map[string]string = nil
	fmt.Println(person)

	name := NewMap("anjay")
	if (name == nil) {
		fmt.Println("Belum ada datanya")
	} else {
		fmt.Println(name)
	}
}
