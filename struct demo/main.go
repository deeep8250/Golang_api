package main

import "fmt"

func main() {

	person := []map[string]interface{}{
		{"name": "deep", "age": 12},
		{"name": "er", "age": 2},
	}

	per := []map[string]interface{}{}

	for _, p := range person {
		if p["name"] != "deep" {
			per = append(per, p)
		}
	}

	fmt.Print(per)

}
