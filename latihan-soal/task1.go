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

}