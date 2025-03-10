package structs

import (
	"fmt"
)

/*	Struct adalah tipe data struktur yang memungkinkan kamu untuk mengelompokkan beberapa variabel
	di dalamnya dengan tipe data yang berbeda.
	Struct biasanya digunakan untuk merepresentasikan entitas atau objek dalam program */

//deklarasi struct menggunakn keyword type dan struct
type person struct { //struct person dideklarasikan memiliki 2 field yaitu name dan age
	name string
	age int
}

//cara mengakses struct

func Struct() {
	// 1. inisialisasi struct langsung
	p := person{"Rafif", 19}
	fmt.Println("Person:", p)

	// 2. inisialisasi struct dengan deklarasi variabel
	var ps person //deklarasi variabel ps bertipe person
	ps.name = "Mong" //set nilai filed name menjadi mong
	ps.age = 14
	fmt.Println("Name:", ps.name) //cetak nilai field name
	fmt.Println("Age:", ps.age)
}