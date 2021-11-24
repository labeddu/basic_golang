package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total

}

func main() {
	total := sumAll(1, 2, 3)
	fmt.Println(total)

	slice := []int{1, 20}

	total2 := sumAll(slice...)
	fmt.Println(total2)

}
