package main

import "fmt"

/*
	Tipe data Slice merupakan potongan dari sebuah Array

	Ukuran Slice bisa berubah

	Slice dan Array saling berhubungan, karena slice didapat dari sebuah Array

	Terdapat 3 komponen dalam slice:
	- Pointer
		Penunjuk data pertama di Array pada slice
	- Length
		Panjang dari slice
	- Capacity
		Kapasitas dari slice, sehingga length tidak boleh melebihi capacity
*/

/*
	Cara membuat Slice
	- arr[l:h]
		Membuat slice dari arr dimulai index l sampai index sebelum h
	- arr[l:]
		Membuat slice dari arr dimulai index l sampai akhir arr
	- arr[:h]
		Membuat slice dari arr dimulai index 0 arr sampai index sebelum h
	- arr[:]
		Membuat slice dari arr dimulai index 0 arr sampai index terakhir arr
*/

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[3:6]
	fmt.Println(slice1)
	fmt.Println("Panjang slice1: ", len(slice1))
	fmt.Println("Kapasitas slice1: ", cap(slice1))
	fmt.Println("slice1[2]: ", slice1[2])
	slice1[2] = "Leenogs"
	fmt.Println("slice1[2]: ", slice1[2])
	fmt.Println("months[4]: ", months[5])

	var slice2 = slice1[:]
	slice2[2] = "Berubah lagi"
	fmt.Println("slice1[2]: ", slice1[2])
	fmt.Println("slice2[2]: ", slice2[2])
	fmt.Println("months[4]: ", months[5])
	fmt.Println("Months: ", months)

	var slice3 = append(slice1, "test", "anjay", "baru", "woi", "Cornelius", "Linux", "Eksdi", "Bambang")
	fmt.Println(slice3)
	slice3[0] = "Leenoggs_Lagi"
	fmt.Println("slice1[2]: ", slice1[2])
	fmt.Println("slice2[2]: ", slice2[2])
	fmt.Println("months[4]: ", months[5])
	fmt.Println("Slice 3 size: ", cap(slice3))
	fmt.Println("Slice3: ", slice3)
	fmt.Println("Months: ", months)

	// membuat slice baru menggunakan function make(tipe_data, len, cap)
	newSlice := make([]string, 3, 5)
	newSlice[0] = "test"
	newSlice[1] = "Cornelius"
	newSlice[2] = "Linux"
	// newSlice[3] = "Haha" ERROR
	newSlice2 := append(newSlice, "Haha")
	fmt.Println("isi NewSlice:", newSlice)
	fmt.Println("Panjang newSlice:", len(newSlice))
	fmt.Println("Kapasitas newSlice:", cap(newSlice))

	fmt.Println("isi NewSlice2:", newSlice2)
	fmt.Println("Panjang newSlice2:", len(newSlice2))
	fmt.Println("Kapasitas newSlice2:", cap(newSlice2))

}
