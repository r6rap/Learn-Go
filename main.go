package main

import "fmt"

func tambah(a, b uint)uint {
	return a + b
}

func main() {
	hasil := tambah(10, 20)

	fmt.Println("Hasil Penjumlahan",hasil)
	fmt.Println("go"+"lang")
}