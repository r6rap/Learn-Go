package channel

import (
	"fmt"
)

/*	select digunakan untuk menunggu channel yang siap mengirim atau menerima data, lalu mengeksekusi blok kode yang
	sesuai dengan channel tersebut. */

// func getAverageNumber & getmaxNumber akan dijalankan sebagai goroutine
func getAverageNumber(data []int, ch chan float64) {
	sum := 0
	for _, v := range data {
		sum += v
	}

	ch <- float64(sum) / float64(len(data)) // hasil dari kalkulasi akan dikirim lewat channel
}

func getMaxNumber(data []int, ch chan int) {
	max := data[0]

	for _, v := range data {
		if max < v {
			max = v
		}
	}

	ch <- max // hasil dari pencarian angka terbesar akan dikirim lewat channel
}

func ChSelect() {
	numbers := []int{3, 4, 3, 5, 6, 3, 2, 2, 6, 3, 4, 6, 3}
	fmt.Println("Numbers:", numbers)

	ch1 := make(chan float64) // ch1 sebagai channel untuk menerima data dari getAverageNumber
	go getAverageNumber(numbers, ch1)

	ch2 := make(chan int) // ch2 sebagai channel untuk menerima data dari getMaxNumber
	go getMaxNumber(numbers, ch2)

	ch3 := make(chan string, 1)

	// perulangan 2 kali karena ada 2 channel yang akan menerima data
	for i := 0; i < 3; i++ {
		// go akan mengecek semua case channel yang tersedia
		select {
		case avg := <- ch1: // var avg menerima data dari channel ch1
			fmt.Printf("Average: %.2f\n", avg)
		case max := <- ch2: // var max menerima data dari channel ch2
			fmt.Printf("Max: %d\n",max)
		case ch3 <- "Hello World": //
			fmt.Println("Mengirim data ke ch3")
		}
	}

	msg := <- ch3
	fmt.Println(msg)
}