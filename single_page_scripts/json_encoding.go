package main

import (
	"encoding/json"
	"fmt"
)

type json_e struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	// Basic data types encoding
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(2)
	fmt.Println(string(intB))

	floatB, _ := json.Marshal(1.34)
	fmt.Println(string(floatB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// encoding to json array and objects
	slcB, _ := json.Marshal([]string{"a", "b", "c"})
	fmt.Println(string(slcB))

	mapB, _ := json.Marshal(map[string]int{"x": 1, "y": 2, "z": 3})
	fmt.Println(string(mapB))

	res1B, _ := json.Marshal(&json_e{Page: 23, Fruits: []string{"apple", "guava", "grapes"}})
	fmt.Println(string(res1B))

	res2B, _ := json.Marshal(&response2{Page: 23, Fruits: []string{"apple", "guava", "grapes"}})
	fmt.Println(string(res2B))
}

/*
Ouput
true
2
1.34
"gopher"
["a","b","c"]
{"x":1,"y":2,"z":3}
{"Page":23,"Fruits":["apple","guava","grapes"]}
{"page":23,"fruits":["apple","guava","grapes"]}
*/
