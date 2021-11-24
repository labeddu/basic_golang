package main

import "fmt"

func hello() (string, string) {
	return "yusran", "tolleng"

}

func namaLengkap() (string, string, string) {
	return "satu", "dua", "tiga"
}

func main() {

	firstName, lastName := hello()
	fmt.Println(firstName, lastName)

	satu, _, tiga := namaLengkap()
	fmt.Println(satu, tiga)

}
