package main

import "fmt"

func main() {
	/* 这是我的第一个简单的程序 */
	fmt.Print("what?")
	fmt.Println("Hello, World!")
	fmt.Println(add(1,2))
}

func add(value1, value2  int) int {
	newValue := value1 + value2
	return newValue
}