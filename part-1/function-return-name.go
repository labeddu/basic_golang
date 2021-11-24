package main

import "fmt"

func geFullName() (firstname string, middleName string, age int) {
	firstname = "yusran"
	middleName = "tolleng"
	age = 40

	return

}

func main() {
	a, b, c := geFullName()

	fmt.Println(a, b, c)

}
