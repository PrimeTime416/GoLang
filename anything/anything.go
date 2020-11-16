package main

import "fmt"

// This function is just to play around
func main() {
	fmt.Println("Go practice code this is really kool :)")

	foo()

	fmt.Println("Let me count to 100: ")
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("I an in foo")
	for i := 0; i <= 50; i++ {
		fmt.Println(i)
	}
}

func bar() {
	fmt.Println("The program is now over :)")
}
