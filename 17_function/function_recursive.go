package main

import "fmt"

func faktorialRekursif(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * faktorialRekursif(n-1)
	}
}

func main() {
	fmt.Println(faktorialRekursif(4))
}
