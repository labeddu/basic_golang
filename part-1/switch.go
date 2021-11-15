package main

import "fmt"

func main() {

	names := "gogogogog"

	switch names {
	case "yusran":
		fmt.Println("hello yusran")
	case "eko":
		fmt.Println("hello eko")
	default:
		fmt.Println("gak kenal")
	}

	length := len(names)
	switch {
	case length > 10:
		fmt.Println("terlalu Panjanga")

	case length > 5:
		fmt.Println("cukup Panjang")

	default:
		fmt.Println("cukup")
	}
	fmt.Println("nama : " + names)

}
