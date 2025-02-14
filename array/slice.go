package array

import "fmt"

/*Perbedaan tipe nilai (value types) dan tipe referensi (reference type) 
value types: variabel dengan type value akan menyimpan nilai datanya langsung di dalam variabel,
			 artinya ketika kamu menyalin variabel value type ke variabel lain, seluruh nilainya disalin
reference types: variabel reference tidak menyimpan data secara langsung, melainkan menyimpan alamat ke data sebenarnya,
				 jadi ketika kamu menyalin variabel reference type ke variabel lain hanya alamat data (pointer)
				 yang disalin sehingga jika data diubah melalui salah satunya, perubahan itu akan terlihat di variabel
				 lainnya.
				 

contoh ada di bawah
*/
func ValuTypes() {
	var arr1 = [3]string{"Mango", "Avocado", "Banana"}

	arr2 := arr1 		//arr2 adalah salinan dari arr1
	arr2[1] = "Pawpaw"	//ubah index ke 1 di arr2

	fmt.Println("Arr1 :", arr1)
	fmt.Println("Arr2 :", arr2)
	//perubahan dari arr2 tidak mempengaruhi arr1 karena array adalah value types
}

func ReferenceTypes() {
	var slice1 = []string{"Avocado", "Banana", "Mango"}

	slice2 := slice1	 //slice2 salinan reference dari slice1 (hanya alamat yang disalin)
	slice2[1] = "Pawpaw" //ubah index ke 1 di slice2

	fmt.Println("Slice1 :", slice1)
	fmt.Println("Slice2 :", slice2)
	//perubahan dari slice2 mempengaruhi slice1 karena keduanya merujuk ke data yang sama
}

func Slice() {
	/*Slice adalah data reference elemen array.
	  Slice bisa dibuat atau juga dihasilkan dari manipulasi sebuah array ataupun slice lainnya.
	  Jika ada slice baru yang terbentuk dari slice lama, maka data elemen slice baru juga memiliki alamat memori
	  yang sama dengan slice lama, setiap perubahan yang terjadi di elemen slice baru, akan berdampak juga
	  ke elemen slice lama yang memiliki referensi yang sama*/
	var fruits = []string{"avocado", "mango", "banana", "pawpaw"}

	//newFruits adalah variabel slice baru yang tercetak dari slice fruits
	newFruits := fruits[0:2] //0:2 digunakan untuk mengakses slice dari fruits dimulai dari index ke 0 hingga elemen sebelum index ke 2
	// index awal (inklusif):index terakhir tetapi tidak disertakan(eksklusif)
	fmt.Println(newFruits)
	fmt.Println(fruits[0:2])
	fmt.Println(fruits[1:3])
}

func SliceTwo() {
	fruits := []string{"melon", "banana", "avocado", "apple"}

	aFruits := fruits[0:3]
	bFruits := fruits[1:4]

	aaFruits := aFruits[1:2]
	bbFruits := bFruits[0:1]

	fmt.Println(fruits)
	fmt.Println(aFruits)
	fmt.Println(aaFruits)
	fmt.Println(bFruits)
	fmt.Println(bbFruits)
}

func Pointer() {
	x := 24
	p := &x //& digunakan untuk mencatat alamat memori dari x

	fmt.Printf("Alamat x: %d\n", x)
	fmt.Printf("Alamat p: %d\n", p)

	*p = 40 //* digunakan untuk mengakses atau mengubah alamat memori tersebut
	fmt.Printf("Nilai x setelah diubah melalui pointer: %d\n", x)
}

