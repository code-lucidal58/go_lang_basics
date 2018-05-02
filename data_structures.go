package main

import "fmt"

func main() {

	//Arrays
	var arr [3]string
	arr[0] = "Hello"
	arr[1] = "Aanisha"
	arr[2] = "Mishra"
	fmt.Println(arr)

	//Slices - Method 1
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	//Slices - Method 2
	g := make([]int, 5, 10)
	s := g[2:4] //again slices using high low expression
	s[0] = 10
	s[1] = 11
	g[4] = 14
	fmt.Println(s)
	a := s[:cap(s)]
	//size of a will be 8 as it will take up
	//complete capacity of s and in turn g
	fmt.Println(a)
}
