package main

import "fmt"

func main() {
	var nama [3]string
	nama[0] = "Linux"
	nama[1] = "Niggajiz"
	nama[2] = "Rizlerr"

	fmt.Println("Nama[0]", nama[0])
	fmt.Println("Nama[1]", nama[1])
	fmt.Println("Nama[2]", nama[2])

	var angka = [3]int{5, 2, 6}
	fmt.Println("Panjang array angka:", len(angka))
	fmt.Println(angka)

	var tinggiSiswa = [...]int{61, 56, 77, 45, 23}
	fmt.Println("Jumlah siswa: ", len(tinggiSiswa))
	fmt.Println(tinggiSiswa)

}
