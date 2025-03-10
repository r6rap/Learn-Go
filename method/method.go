package method

import(
	"fmt"
)

/*	Method adalah function yang menempel pada suatu tipe data, misalnya custom struct.
	Method bisa diakses lewat variabel objek yang dibuat dari tipe custom struct */

type Person struct{
	name string
}

func (person Person) sayHello(nama string)  {
	fmt.Println("Haloo", person.name, "Nama saya", nama)
}

func Method() {
	p1 := Person{name: "Rapip"}
	p1.sayHello("Eko")
}