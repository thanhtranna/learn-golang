package main

import "fmt"

func f(from string) {
	for index := 0; index < 3; index++ {
		fmt.Println(from, " : ", index)
	}
}

func main() {
	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	fmt.Scanln()
	fmt.Println("done")
}
