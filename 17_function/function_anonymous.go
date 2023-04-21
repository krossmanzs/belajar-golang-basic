package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println(name, "are blocked!")
	} else {
		fmt.Println(name, "are welcomed!")
	}
}

// tidak wajib
// func blackListAdmin(name string) bool {
// 	return name == "Admin"
// }

// func blackListRoot(name string) bool {
// 	return name == "Root"
// }

func main() {
	// gunakan seperti ini
	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blacklist)
	registerUser("user", blacklist)

	registerUser("linux", func(name string) bool {
		return name == "root"
	})
}
