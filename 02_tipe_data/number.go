package main

import "fmt"

/*
	Tipe Data Integer
				Nilai Min				Nilai Max
	int8		-128					127
	int16		-32768					32767
	int32		-2147483648				2147483647
	int64		-9223372036854775808	9223372036854775807

	Tipe Data Unsigned Integer
				Nilai Min				Nilai Max
	uint8		0						255
	uint16		0						65535
	uint32		0						4294967295
	uint64		0						18446744073709551615

	Tipe Data Floating Point
				Nilai Min				Nilai Max
	float32		1.18*10^-38				3.4*10^38
	float64		2.23*10^-308			1.80*10^308
	complex64	complex numbers with float32 real and imaginary parts.
	complex128	complex numbers with float64 real and imaginary parts.

	Alias (seperti define di c / nama lain)
	byte --> uint8
	rune --> int32
	int  --> min. int32 (tergantung bit OS)
	uint --> min. int32 (tergantung bit OS)
*/

func main() {
	fmt.Println("Satu =", 1)
	fmt.Println("Dua =", 2)
	fmt.Println("tiga koma lima =", 3.5)
	fmt.Println("lima kali tiga =", 5*3)
}
