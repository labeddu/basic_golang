package main

import "fmt"

func sayHelloTo(name string) string {
	return "hello " + name
}

func main() {
	result := sayHelloTo("yusran")
	fmt.Println(result)
}
