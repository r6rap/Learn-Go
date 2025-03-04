package function

import "fmt"

/*	Anonymous function adalah function tanpa nama,
	ini berguna jika kita membutuhkan function sederhana. */

type Blacklist func(string) bool 

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You're blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func Register() {
	//1. menulis anonymous function di variabel blackList
	blackList := func(nama string) bool {
		return nama == "anjing"
	}
	registerUser("Rafif", blackList)

	//2. menulis langsung di parameternya
	registerUser("anjing", func(nama string) bool {
		return nama == "anjing"
	})
}