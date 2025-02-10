package array

import "fmt"

func Array() {
	//cara penulisan array part 1
	var names[3] string
	names[0] = "Monkey"
	names[1] = "D"
	names[2] = "Luffy"

	fmt.Println(names[0], names[1], names[2])
}

func ArrayTwo() {
	//cara penulisan array part 2
	var fruits [4]string

	fruits = [4]string{"Mango", "Avocado", "Pawpaw", "Watermelon"}

	fmt.Println("Total element array \t \t", len(fruits)) //fungsi len untuk menghitung jumlah element sebuah array
	fmt.Println("Isi semua element \t", fruits)
}

func ArrayThree() {
	//jika lebar array-nya tidak diset dari awal bisa diganti menggunakan tanda ...
	var numbers = [...]int{2, 3, 5, 6, 7}

	fmt.Println("Total element array \t", len(numbers))
	fmt.Println("Isi element \t", numbers)
}

func ArrayMultidimention() { //ada 2 cara untuk menulis array multidimensi
	//cara 1
	var numbers1 = [2][3]int{
		[3]int{3,2,3},
		[3]int{4,2,5},
	}

	//cara 2
	var numbers2 = [2][3]int{
		{3,2,4},
		{4,2,6},
	}

	fmt.Println(numbers1)
	fmt.Println(numbers2)
}

func LoopInArray() {
	fruits := [4]string{
		"Mango",
		"Avocado",
		"Pawpaw",
		"Banana",
	}

	for i := 0; i < len(fruits); i++{
		fmt.Printf("Element ke %d : %s\n", i, fruits[i])
	}
}

func ForRange() {
	//for range di golang
	var fruits = [4]string{
		"Avocado",
		"Banana",
		"Mango",
		"Pawpaw",
	}

	//i : digunakan untuk menampung index dari fruits
	//fruit : digunakan untuk menampung value / isi dari fruits
	for i, fruit := range fruits{
		fmt.Printf("Element %d : %s\n", i, fruit)
	}
}

func ForRange2() {
	var fruits = [4]string{"avocado", "mango", "banana", "pawpaw"}

	//jika hanya ingin menampilkan valuenya, tampung saja nilai yang tidak ingin digunakan ke underscore (_)
	for _, fruit := range fruits{
		fmt.Printf("Nama buah: %s\n", fruit)
	}

	//sama seperti diatas disini hanya menampilkan index dari array
	for i, _ := range fruits{
		fmt.Printf("Index ke %d\n", i)
	}
}

func MakeInGo() {
	//deklarasi sekaligus alokasi kapasitas array menggunakan keyword make

	//parameter pertama diisi dengan element array dengan tipe data string
	//parameter kedua diisi dengan jumlah elementnya
	var fruits = make([]string, 2)

	fruits[0] = "Avocado"
	fruits[1] = "Mango"

	fmt.Println(fruits)
}