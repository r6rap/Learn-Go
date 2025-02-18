package latihansoal

import (
	"fmt"
)

func TaskOne() {
	//menampilkan isi slice
	fruits := []string{"apel", "jeruk", "pisang", "mangga"}

	for i, buah := range fruits {
		fmt.Printf("Index %d : %s\n", i, buah)
	}
}

func TaskTwo() {
	//menjumlahkan semua isi slice
	numbers := []int{2, 4, 6, 8, 10}
	sum := 0
	sum2 := 0

	for _, v := range numbers{
		sum += v
	}

	for i := 0; i<len(numbers); i++{
		sum2 += numbers[i]
	}
	fmt.Println("dari range",sum)
	fmt.Println("dari for", sum2)
}

func TaskThree() {
	//membuat slice baru yang isinya nilai kuadrat dari slice lama
	numbers := []int{2, 4, 6, 8, 10, 12}
	square := make([]int, len(numbers)) //membuat slice dengan panjang dan kapasitas mengikuti panjang dari numbers
	squares := make([]int, len(numbers))

	for i := 0; i < len(numbers); i++ {
		square = append(square, numbers[i] * numbers[i])
	}

	for _, v := range numbers{
		squares = append(squares, v*v)
	}

	fmt.Println("Hasil dari for:", square)
	fmt.Println("Hasil dari range:", squares)
}

func TaskFour() {
	//membuat slice baru yang hanya berisi angka angka yang lebih besar dari 10
	numbers := []int{7, 12, 5, 20, 9, 3, 11}
	place := make([]int, 0, len(numbers))//membuat slice dengan panjang 0 dan kapasitas sesuai panjang numbers
	place2 := make([]int, 0, len(numbers))

	for i := 0; i < len(numbers); i++{
		if numbers[i] > 10 {
			place = append(place, numbers[i])
		}
	}
	for _, v := range numbers{
		if v > 10 {
			place2 = append(place2, v)
		}
	}
	fmt.Println("Hasil dari for",place)
	fmt.Println("Hasil dari range",place2)
}