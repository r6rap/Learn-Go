package latihansoal

import "fmt"

func FirstTask() {
	for i := 1; i <= 20; i++ {
		if i % 3 == 0 {
			continue
		}

		fmt.Println("angka yang tidak habis dibagi 3", i)
	}
}

func SecTask() {
	for i := 1; i < 20; i++ {
		fmt.Println("loop ke", i)
		if i == 15 {
			fmt.Println("iterasi dihentikan pada loop ke", i)
			break
		}
	}
}

func ThirdTask() {
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 10; j++ {
			if j % i == 0 {
				continue
			}

			if j > 8 {
				break
			}

			fmt.Println(" nilai i =", i , "nilai j =", j)
		}
	}
}

func FourthTask() {
	numbers := [...]int{
		12,-5,34,0,150,7,
	}

	for i := 0; i < len(numbers); i++{
		if numbers[i] < 0{
			fmt.Println("Angka negatif diabaikan")
			continue
		}
		if numbers[i] > 100{
			fmt.Println("Angka terlalu besar, perulangan dihentikan")
			break
		}

		fmt.Printf("Element %d : %d\n", i, numbers[i])
	}
}