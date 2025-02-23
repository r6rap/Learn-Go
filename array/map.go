package array

import "fmt"

/*Map adalah tipe data asosiatif yang berbentuk key dan value
	contoh di bawah, saya membuat variabel day dengan tipe data map
	dengan tipe data key string dan tipe data value int */

//nilai default tipe data map adalah nill
//pengaksesan key yang belum diinisialisasi nilainya, maka nilai defaultnya mengikuti tipe data valuenya

func Map() {
	day := make(map[string]int)

	day["Sunday"] = 1
	day["Monday"] = 2

	fmt.Println(day)
	fmt.Println(day["Monday"])
	fmt.Println(day["Tuesday"])
}