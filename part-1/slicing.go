package main

import "fmt"

func main() {

	var nama = "yusran"

	fmt.Println(nama)
	// fmt.Println(type(nama))

	fmt.Println("hello")

	data1 := "dataku"
	fmt.Println(data1)
	fmt.Println(type(data2))

	var month = [...]string{
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

	var slice1 = month[4:7]
	var slice2 = month[4:]
	// var slice3 = month[:12]

	fmt.Println(slice1, slice2)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	fmt.Println(month)

}
