package structs

import (
	"fmt"
)

/*	Anonymous struct adalah struct yang tidak dideklarasikan di awal sebagai tipe data baru, melainkan
	langsung ketika pembuatan object.
	Teknik ini cukup efisien jika membutuhkan struct yang hanya dipakai sekali */

type child struct{
	name string
	age int
}


/*	saat membuat anonymous struct setelah dideklarasikan anonymous struct harus langsung diinisialisasi 
	dengan cara menambahkan kurung kurawal setelah struct dideklarasikan */

func AnonymousStruct() {
 // variabel c1 langsung diisi objek anonymous struct
	c1 := struct {
		child
		gender string
	}{} // mirip seperti IIFE
	c1.child = child{"Rafif", 17}
	c1.gender = "Male"

	fmt.Println(c1.name)
	fmt.Println(c1.age)
	fmt.Println(c1.gender)
}

