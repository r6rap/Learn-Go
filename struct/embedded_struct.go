package structs

import (
	"fmt"
)

/*	Embedded struct adalah mekanisme untuk menanamkan satu struct ke struct lain tanpa harus
	mendefinisikan fieldnya */

type people struct{
	name string
	age int
}

type work struct{
	people //struct people di-embed di struct work
	age int
	salary int
}

//	Embedded struct dengan nama field yang sama

func Embedded() {
	var peope = work{}
	peope.name = "Rafif"
	peope.age = 20 // age from struct work
	peope.salary = 5000000
	peope.people.age = 19 // age from struct people

	fmt.Println("Nama:", peope.name)
	fmt.Println("umur work:", peope.age) // 0 karena field age dari struct work di shadowing oleh field age dari struct people
	fmt.Println("Salary:", peope.salary)
	fmt.Println("Umur people:", peope.people.age) // untuk mengaksesnya dengan cara menuliskan nama structnya lalu dilanjut fieldnya
}

// atau bisa lebih simpel seperti ini

func Embedded2() {
	pep := people{name: "Rap", age: 20}
	peope := work{people: pep, salary: 5000000} //field person pada struct work diisi dengan variabel pep

	fmt.Println("Nama:", peope.name)
	fmt.Println("umur work:", peope.age) // 0 karena field age dari struct work di shadowing oleh field age dari struct people
	fmt.Println("Salary:", peope.salary)
	fmt.Println("Umur people:", peope.people.age) // untuk mengaksesnya dengan cara menuliskan nama structnya lalu dilanjut fieldnya
}