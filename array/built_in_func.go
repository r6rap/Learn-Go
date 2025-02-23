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
	fmt.Println("Sebelum di append dengan cFruits:",fruits)

	cFruits := fruits[0:3] //avocado, apple, banana
	/*mango berubah menjadi papaya karena len(cFruits) < cap(fruits)
		maka papaya akan ditempatkan di kapasitas */
	cFruits = append(cFruits, "papaya")

	fmt.Println("Setelah di append cFruits:", fruits)
	fmt.Println("bFruits:",bFruits)
	fmt.Println("cFruits:",cFruits)
}

func Copy() {
	/* copy(dst, src) digunakan untuk menyalin element slice
		parameter pertama dst (destination) : slice tujuan tempat element akan disalin
		parameter kedua src (source) : sumber slice yang elementnya akan disalin
		n := copy(dst, src) akan mengembalikan jumlah element yang berhasil disalin
		function copy() hanya mengcopy element sebenyak len() dari destination / dst */

	src := []string{"watermelon", "pinnaple", "apple", "orange"}
	dst := make([]string, 3)
	n := copy(dst, src)

	//orange tidak disalin len(dst) hanya 3 karena function copy() hanya mengcopy element sebanyak len(dst)
	fmt.Println("dst atau tujuan element slice disalin:",dst)
	fmt.Println("src atau sumber element yang disalin",src)
	fmt.Println("element yang berhasil disalin:",n)

	fmt.Println("===========================")

	src2 := []string{"watermelon", "avocado"}
	//2 element pertama dari dst2 diganti dengan src2 karena src2 hanya berisi 2 element itu
	dst2 := []string{"potato", "potato", "potato"}
	n2 := copy(dst2, src2)

	fmt.Println(dst2)
	fmt.Println(src2)
	fmt.Println(n2)
}