package main

import "fmt"

func getGoodBye(name string) string {
	return "Goodbye " + name

}

func main() {
	sayGoodBye := getGoodBye

	result := sayGoodBye("yus")

	fmt.Println(result)
	fmt.Println(getGoodBye("eko"))

}
