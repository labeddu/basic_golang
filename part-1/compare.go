package main

import "fmt"

func main() {

	var ujian = 95
	var absensi = 91

	var lulusUjian = ujian > 90
	var lulusAbsensi = absensi > 80

	var lulus = lulusAbsensi && lulusUjian

	fmt.Println(lulus)
}
