package goroutine

import (
	"fmt"
	"time"
)

/*	goroutine adalah thread ringan yang dikelola oleh go runtime, dengan goroutine kita bisa menjalankan
	fungsi secara bersamaam (concurrently). goroutine dijalankan secara asynchronous */

func print(limit int, message string) {
	for i := 0; i < limit; i++{
		fmt.Println((i+1), message)
	}
}

func Goroutines() {
	go print(5,"Halo")
	print(5, "Hai")

	time.Sleep(2 * time.Second)
}