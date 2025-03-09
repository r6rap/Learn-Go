package structs

import(
	"fmt"
)

/*	field yang dibuat dari tipe struct bisa diambil pointernya, dan bisa disimpan
	pada variabel objek yang bertipe struct pointer */

type student struct{
	name string
	grade int
}

func StructPointer() {
	s1 := student{"Rap", 4}
	// s2 adalah variabel pointer hasil dari student
	var s2 *student = &s1 // s2 menampung nilai referensi/alamat dari s1
	// sehingga setiap perubahan pada property variabel s2 juga akan berpengaruh ke variabel s1

	fmt.Println("S1 :", s1.name)
	fmt.Println("S2 :", s2.name)

	s2.name = "Bert"

	fmt.Println("S1 :", s1.name)
	fmt.Println("S2 :", s2.name)
}