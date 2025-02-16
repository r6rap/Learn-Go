package latihansoal

import "fmt"

func TaskOne() {
	fruits := []string{"apel", "jeruk", "pisang", "mangga"}

	for i, buah := range fruits {
		fmt.Printf("Index %d : %s\n", i, buah)
	}
}

func TaskTwoV1() {
	numbers := []int{2, 4, 6, 8, 10}
	sum := 0

	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	fmt.Println(sum)
}