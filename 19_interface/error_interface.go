package main

import (
	"errors"
	"fmt"
)

func pembagian(angka int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		result := angka / pembagi
		return result, nil
	}
}

func main() {
	angka1 := 6
	angka2 := 3

	hasil, err := pembagian(angka1, angka2)

	if err == nil {
		fmt.Println("Hasil:", hasil)
	} else {
		fmt.Println("Error:", err.Error())
	}

}
