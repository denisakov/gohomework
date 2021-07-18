package main

import (
	"fmt"

	"example.com/Lesson_02/fibonacciOne"
)

func main() {
	elementCount := 9
	defer fmt.Println("There you have it")
	fmt.Print("Here comes the fibonacci sequence up to ")
	fmt.Print(elementCount)
	fmt.Print(" element")
	fmt.Println("")
	for i := 0; i <= elementCount; i++ {
		fmt.Print(fibonacciOne.Fib(i))
		fmt.Print(" ")
	}
	fmt.Println("")

}
