package array

import "fmt"

/*Map adalah tipe data asosiatif yang berbentuk key dan value
	contoh di bawah, saya membuat variabel day dengan tipe data map
	dengan tipe data key string dan tipe data value int */

//nilai default tipe data map adalah nill
//pengaksesan key yang belum diinisialisasi nilainya, maka nilai defaultnya mengikuti tipe data valuenya

func Map() {
	day := make(map[string]int)
	_ = *new(map[string]int)//khusus new yang dihasilkan adalah data pointer

	day["Sunday"] = 1
	day["Monday"] = 2
	day["Monday"] = 4

	fmt.Println(day)
	fmt.Println(day["Monday"])
	fmt.Println(day["Tuesday"])
}

func IterationMap() {
	chicken := make(map[string]int)

	chicken["januari"] = 20
	chicken["februari"] = 30
	chicken["maret"] = 10

	for key, v := range chicken{
		fmt.Println(key, " :", v)
	}
}

func DeleteMap() {
	/*delete(m, key)
		parameter pertama adalah map yang ingin anda hapus elemennya
		parameter kedua adalah key yang ingin dihapus dari map */
	chicken := map[string]int{
		"senin" : 5,
		"selasa" : 10,
		"rabu" : 8,
	}

	delete(chicken, "selasa")
	fmt.Println(len(chicken))
	fmt.Println(chicken)
}

func IsExist() {
	/* ada cara unik untuk mengecek di variabel map ada key tertentu atau tidak, yaitu dengan memanfaatkan
		2 variabel sebagai penampung pengembalian nilai saat mengakses item  */
	banana := make(map[string]int)
	banana["selasa"] = 10
	banana["senin"] = 30

	//value berisi nilai dari key "selasa"
	//isExist bool akan bernilai true jika key tersebut ada di map dan false jika tidak
	var value, isExist = banana["selasa"]

	if isExist{
		fmt.Println(value)
	}else{
		fmt.Println("Item is not exists")
	}
}