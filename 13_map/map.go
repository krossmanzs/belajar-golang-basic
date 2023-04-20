package main

import (
	"fmt"
)

func main() {
	person := map[string]string{
		"name": "Cornelius Linux",
		"nim":  "122140079",
	}

	person["major"] = "Informatics Engineering"
	person["age"] = "18"

	delete(person, "age")

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["nim"])
	fmt.Println(len(person))

	person2 := make(map[string]string)
	person2["name"] = "Awwikwok"
	fmt.Println(person2)
}
