package latihansoal

import "fmt"

/* Saya meminta chatgpt untuk memberikan latihan soal tetapi tidak dengan kodenya agar saya bisa mengerjakan sendiri.
link command: https://chatgpt.com/share/67b1ea7f-2664-8006-a0ef-f1581713d89c*/

func FirstTask() {
	/*Buatlah program yang melakukan perulangan untuk mencetak angka dari 1 sampai 20.
	Namun, jika angka tersebut habis dibagi 3, maka lewati (skip) iterasi itu menggunakan perintah
	continue sehingga angka yang dicetak hanyalah angka yang tidak habis dibagi 3.*/
	for i := 1; i <= 20; i++ {
		if i % 3 == 0 {
			continue
		}

		fmt.Println("angka yang tidak habis dibagi 3", i)
	}
}

func SecTask() {
	/*Buatlah program yang melakukan perulangan untuk mencetak angka dari 1 sampai 20.
	Namun, ketika angka mencapai 15, hentikan perulangan dengan menggunakan break dan cetak
	pesan "Loop dihentikan pada angka 15".*/
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