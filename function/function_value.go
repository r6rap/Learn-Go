package function

import "fmt"

/* di go function juga bisa digunakan sebagai value */

func goodBye(name string) string {
	return "Goodbye " + name
}

func Hi() {
	hi := goodBye
	fmt.Println(hi("Rafif"))
}