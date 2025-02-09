package chapter1

import (
	"fmt"
)

func Condition() {
	nilai := 49

	if nilai >= 50 {
		fmt.Println("C")
	} else if nilai >= 80 {
		fmt.Println("B")
	} else if nilai >= 90 {
		fmt.Println("A")
	} else {
		fmt.Printf("Anda tidak lulus, nilai anda %d \n", nilai)
	}
}

func Temporary()  {
	/* variabel temporary variabel yang hanya bisa digunakan 
		pada blok kondisi di mana dia ditempatkan*/

	var point = 8840.0

	/* contoh variabel percent hanya bisa digunakan pada blok kondisi if else,
	   variabel temporary hanya bisa dibuat menggunakan type inference (:=)*/

	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect! \n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good \n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad \n", percent, "%")
	}

}

func Switch() {
	/* 	switch di go ketika casenya terpenuhi dia tidak akan melanjutkan ke case selanjutnya,
		berbeda dengan bahasa pemrograman lainnya ketika case terpenuhi maka akan tetap lanjut
		mengecek case selanjutnya kecuali ada keyword break*/

	var point = 1

	switch point {
	case 6, 8: //solusi gunakan 2 kondisi di dalam case
		fmt.Println("good")
	case 5:
		fmt.Println("not bad")
	default: { // {} berfungsi untuk memberikan lebih banyak statement di dalam case
		fmt.Println("bad")
		fmt.Println("you can be better!")
		}	
	}
}

func SwitchCondition() {
	/*	di go switch bisa ditulis dengan gaya if else nilai yang dibandingkan tidak ditulis setelah switch,
		melainkan ditulis langsung dalam bentuk perbandingan dalam keyword case */

	var point = 5

	switch {
	case point == 8:
		fmt.Println("good")
	case (point < 8) && (point > 3):
		fmt.Println("not bad")
	default:{
		fmt.Println("bad")
		fmt.Println("you can be better!")
		}
	}
}

func Fallthrough() {
	/* 	keyword fallthrough digunkan untuk memaksa proses pengecekan ke case selanjutnya
		bahkan jika kondisi case berikutnya tidak terpenuhi */

	var point = 6

	switch {
	case point == 8:
		fmt.Println("good")
	case (point < 8) && (point > 4):
		fmt.Println("not bad")
		fallthrough
	case point < 5:
		fmt.Println("bad") // bad tetap diprint padahal seharusnya kondisinya false
	default: {
		fmt.Println("don't worry")
		fmt.Println("you can be better")
		}
	}
}

func NestingCondition () {
	/* ini adalah kondisi bercabang yang dimana seleksi kondisi berada di dalam seleksi kondisi */
	point := 8 // deklarasi variabel point yaitu berisi nilai 8

	//kita cek apakah point lebih dari 7
	if point > 7 {
		switch point { // jika true maka akan lanjut ke switch condition
		case 10: // disini kita cek apakah point == case
			fmt.Println("perfect") // jika true maka akan mengeksekusi statement ini
		default: // jika tidak ada case yang terpenuhi maka akan mengeksekusi statement ini
			fmt.Println("nice")
		}
	// kalau point < 7 maka akan melakukan pengecekan ke kondisi ini
	} else {
		if point == 5 { // cek apakah point == 5
			fmt.Println("not bad")
		} else if point == 3 { // jika point tidak sama dengan 5 makan akan cek apakah point == 3
			fmt.Println("u can do it!")
		} else{ // jika point tidak sama dengan 5 ataupun 3
			fmt.Println("keep trying!") // sebelum statement ini dieksekusi
			if point == 0 { // kita cek apakah point == 0, jika tidak maka akan mengeksekusi statement sebelumnya
				fmt.Println("try harder!") // jika iya maka akan mengeksekusi statement di bawah
			}
		}
	}
}