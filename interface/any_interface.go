package interfaces

import "fmt"

/*	Interface kosong atau empty interface yang dinotasikan dengan interface{}
atau any, merupakan tipe data yang sangat spesial karena variabel bertipe ini
bisa menampung segala jenis data, baik itu numerik, string, bahkan array, pointer,
apapun */

func Kantin() {
	// slice data berisi elemen berupa map dengan key string dan value interface
	// ini memungkinkan setiap nilai di dalam map memiliki tipe data yang berbeda 
	var data = []map[string]any{
		{"Nama" : "Rafif", "Umur" : 19,
		 "Minat" : []string{"Coding", "crypto"}},
		{"Nama" : "Someone", "Umur" : 17,
		 "Minat" : []string{"Coding", "web3"}},
	}

	// iterasi melalui setiap elemen pada slice data
	for _, key := range data {
		fmt.Println("Nama:",key["Nama"], "Umur:", key["Umur"].(int), "Minat:", key["Minat"])
		// type assertion / casting untuk memastikan bahwa value dari key umur bertipe int
	}

	var snack any

	snack = "Roti"
	fmt.Println("Bisa menyimpan makanan:", snack)

	snack = "Jus mangga"
	fmt.Println("Bisa menyimpan minuman:", snack)

	snack = 3000.00
	fmt.Printf("Bisa menyimpan harga %.2f\n", snack.(float64))
}