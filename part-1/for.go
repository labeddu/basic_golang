package main

import "fmt"

func main() {

	// //model  normal
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("perulangan ke ", counter)
	// 	counter++
	// }

	//dengan menggunakan init statement

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("init statement ke: ", counter)
	}

	slice := []string{"makassar", "pare", "bojo"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	fmt.Println("cara lain :")

	//tanda (_) digunakan untuk variabel yg dideklarasikan tapi tidak digunakan
	for _, value := range slice {
		// fmt.Println("Index", i, "= ", value)
		fmt.Println(value)

	}

}
