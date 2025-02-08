package chapter1

import "fmt"

func ForLoop1() {
	for i := 1; i <= 5; i++{
		fmt.Println("angka ke", i)
	}
}

func ForLoop2() {
	i := 0 // deklarasi variabel i dengan nilai 0

	for i < 5{ // jika i kurang dari 5 maka akan mengeksekusi kode ini
		fmt.Println("angka ke", i)
		i++ // i+1
	}
}

func ForLoop3() {
	// for loop ini ditulis tanpa kondisi tetapi menggunakan keyword break agar perulangan berhenti
	i := 0

	for {
		fmt.Println("angka ke", i)
		i++

		if i == 5 { // jika i == 5 maka for loop akan berhenti
			break // keyword break digunakan untuk menghentikan paksa sebuah perulangan
		}
	}
}

func ForLoop4() { /* menampilkan bilangan genap yang lebih besar dari 0 dan kurang dari atau sama dengan 8 */
	for i := 1; i <= 10; i++ { //melakukan perulangan sampai 10
		if i % 2 == 1 { //jika i / 2 sisa 1 maka akan melanjutkan ke perulangan selanjutnya
			continue // keyword continue digunakan untuk memaksa ke perulangan selanjutnya
		}

		if i > 8 { //cek apakah i lebih dari 8
			break //jika iya maka perulangan akan berhenti
		}

		fmt.Println("angka ke", i)
	}
}

func NestedLoop() {
	for i := 0; i < 5; i++{
		for j := i; j < 5; j++{ //menghabiskan j terlebih dahulu
			fmt.Print(j, " ")
		}
		fmt.Println() //lalu print baris baru
	}
}

func LabelInGo() {
	outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++{
			if i == 3 { //ketika i == 3
				break outerLoop //maka akan mengehentikan paksa perulangan yang ada di dalam label outerLoop
			}

			fmt.Println("matriks [", i, "][", j, "]")
		}
	}
}