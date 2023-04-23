package main

import "fmt"

type Mahasiswa struct {
	Nama, Nim string
	Umur      int
}

func (mahasiswa Mahasiswa) sayHello() {
	fmt.Println("Halo", mahasiswa.Nama)
}

func printMahasiswa(data Mahasiswa) {
	fmt.Println("Nama:", data.Nama)
	fmt.Println("NIM:", data.Nim)
	fmt.Println("Umur:", data.Umur)
	data.sayHello()
	fmt.Println()
}

func main() {
	var mhs1 Mahasiswa
	mhs1.Nama = "Cornelius Linux"
	mhs1.Nim = "122140079"
	mhs1.Umur = 18

	mhs2 := Mahasiswa{
		Nama: "Awikwok Bujirman",
		Nim:  "11221400888",
		Umur: 16,
	}

	printMahasiswa(mhs1)
	printMahasiswa(mhs2)
}
