package chapter1

import "fmt"

func Kerucut() {
	r, tinggi := 14, 10

	const(
		pi float32 = 3.14
		def float32 = 0.33
	)

	hasil := def * pi * float32(r*r) * float32(tinggi) //tipe data harus sama agar bisa melakukan perhitungan

	fmt.Println(hasil)
}

func Operator() {
	value := (((2+6) % 3) *4-2) /3
	isEqual := value == 2

	fmt.Printf("apakah %d = 2 ? %t \n", value, isEqual) // %t menampilkan nilai bool
}