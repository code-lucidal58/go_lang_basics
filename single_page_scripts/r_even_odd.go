package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, item := range numbers {
		if item%2 == 0 {
			fmt.Printf("%d is even\n", item)
		} else {
			fmt.Printf("%d is odd\n", item)
		}
	}
}
