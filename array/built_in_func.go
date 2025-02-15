package array

import "fmt"

func Len() {
	//len di golang akan mengembalikan jumlah elemen yang ada pada struktur data
	fruits := []string{"avocado", "apple", "banana", "mango"}

	fmt.Println("Jumlah elemen slice",len(fruits))
}

func Cap() {
	//cap akan mengembalikan jumlah kapasitas dari struktur data
	fruits := []string{"avocao", "apple", "banana", "mango"}

	fmt.Println("cap dari fruits:", cap(fruits))

	aFruits := fruits[0:3]
	bFruits := fruits[1:4]

	fmt.Println("Cap dari aFruits:", cap(aFruits))
	//kenapa bFruits 3? karena kapasitas yang dihitung dimulai dari index ke 1 bukan index ke 0
	fmt.Println("Cap dari bFruits:", cap(bFruits))
}

func Append() {
	//append digunakan untuk menambah elemen pada slice, elemen tersebut akan diposisikan setelah index terakhir

	/* Hal penting saat menggunakan append
		1. Jika jumlah elemen (len) sama dengan kapasitas elemen (cap) maka
			elemen baru hasil append merupakan referensi baru
		2. Jika jumlah elemen (len) kurang dari kapasitas elemen (cap) maka
			elemen baru dari append akan ditempatkan di kapasitas (cap), menjadikan semua slice
			yang memiliki referensinya sama akan berubah nilainya*/

	fruits := []string{"avocado", "apple", "banana", "mango"}

	bFruits := append(fruits, "papaya")
	cFruits := fruits[0:3]
	//mango berubah menjadi papaya karena len < cap
	cFruits = append(cFruits, "papaya")

	fmt.Println(fruits)
	fmt.Println(bFruits)
	fmt.Println(cFruits)
}