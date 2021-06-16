package main

import "fmt"

func main() {
	name := "bill"
	namePointer := &name
	fmt.Println(&namePointer)
	printPOinter(namePointer)
}

func printPOinter(namePOinter *string) {
	fmt.Println(&namePOinter)
}
