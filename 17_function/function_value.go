package main

import (
	"fmt"
	"strconv"
)

func getAge(age int) string {
	return "Umur anda adalah " + strconv.Itoa(age)
}

func main() {
	age := getAge
	fmt.Println(age(18))
}
