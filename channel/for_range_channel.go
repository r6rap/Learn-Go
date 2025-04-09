package channel

import "fmt"

/*	for range channel, membuat loop sebanyak data yang diterima oleh channel */

// func ini khusus hanya untuk mengirim data ke channel
func sendMessage(ch chan <- string) {
	for i := 0; i < 20; i++ {
		ch <- fmt.Sprintf("Data: %d",i)
	}
	close(ch) // channel akan diclose ketika semua data sudah terkirim
}

// func ini khusus untuk menerima data dari channel
func printMessage(ch <- chan string) {
	for message := range ch {
		fmt.Println(message)
	}
}

func CH() {
	ch := make(chan string)
	go sendMessage(ch)
	printMessage(ch)
}