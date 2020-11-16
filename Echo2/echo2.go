package main

import (
	"fmt"
	"os"
)

//Fake echo
func main() {
	fmt.Println("Hello Echo2")
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

}
