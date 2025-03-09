package pointer

import "fmt"

/*	Pointer adalah reference atau alamat memori dari suatu nilai. 
	contoh: ada sebuah variabel bertipe int yang bernilai 4, maka yang dimaksud pointer adalah alamat
	memori di mana nilai 4 disimpan, bukan nilai 4 itu sendiri
	
	variabel pointer ditandai dengan asterisk (*) tepat sebelum tipe data dideklarasikan
	nilai default pointer adalah nil
	
	ampersand (&) = digunakan untuk mengambil pointer dari variabel biasa, referencing
	asterisk (*) = digunakan untuk mengambil nilai asli dari variabel pointer, deferencing */

func Pointer() {
	var numberA int = 4
	var numberB *int = &numberA //numberB berisi alamat dari numberA

	fmt.Println("Number A (value) :", numberA)
	fmt.Println("Number A (address) :", &numberA) //mengambil pointer dari numberA

	fmt.Println("Number B (value) :", *numberB) //mengambil nilai asli dari alamat numberB
	fmt.Println("Number B (address) :", numberB)

	numberA = 5 //satu diubah semua ikut berubah, karena memiliki referensi/alamat yang sama

	fmt.Println("Number A (value) :", numberA)
	fmt.Println("Number A (address) :", &numberA) //mengambil pointer dari numberA

	fmt.Println("Number B (value) :", *numberB) //mengambil nilai asli dari alamat numberB
	fmt.Println("Number B (address) :", numberB)
}

func PointerAsParameter() {
	num := 4
	fmt.Println("Before:", num)

	change(&num, 10)
	fmt.Println("After:", num)
}

func change(referenece *int, value int) {
	*referenece = value
}