package main

import "fmt"

type Filter func(string) string

func sayHello(name string, filter Filter) {
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHello("Linux", spamFilter)
	sayHello("Anjing", spamFilter)
}
