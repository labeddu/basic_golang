package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "yusran",
		"address": "makassar",
	}

	person["titel"] : "programmer"

	fmt.Println(person)
	// fmt.Println(person.name)
	fmt.Println(person["address"])

}
