package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type json_d struct {
	Name   string   `json:"name"`
	Age    int      `json:"age"`
	Fruits []string `json:"fruits"`
}

func main() {
	byt := []byte(`{"name":"Aanisha", "age":23, "fruits":["apple","guava","grapes"]}`)
	var data map[string]interface{}

	//decoding to primitive data types
	if err := json.Unmarshal(byt, &data); err != nil {
		panic(err)
	}

	fmt.Println(data)
	fmt.Println(data["age"].(float64))

	fruits := data["fruits"].([]interface{})
	fmt.Println(fruits[0].(string))

	//decoding to custom data types
	str1 := `{"name":"Aanisha", "age":23, "fruits":["apple","guava","grapes"]}`
	res1 := json_d{}
	if err := json.Unmarshal([]byte(str1), &res1); err != nil {
		panic(err)
	}
	fmt.Println(res1)
	fmt.Println(res1.Name)

	//Streaming JSON Encoding to streams
	// can to done to HTTPS response also
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"a": 1, "b": 2}
	if err := enc.Encode(d); err != nil {
		panic(err)
	}
}

/*
Output
map[name:Aanisha age:23 fruits:[apple guava grapes]]
23
apple
{Aanisha 23 [apple guava grapes]}
Aanisha
{"a":1,"b":2}
*/
