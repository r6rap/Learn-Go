package latihansoal

import (
	"fmt"
	"time"
	"os"
)

func timer(runtime int, ch chan <- bool) {
	time.AfterFunc(time.Duration(runtime)*time.Second, func() {
		ch <- true // ketika tidak ada aktivitas selama 5 detik, maka mengirim true ke channel ch
	})
}

func watch(runtime int, ch <- chan bool) {
	<- ch // menerima data dari channel ch
	fmt.Println("Timeout no answer more than", runtime, "second")
	os.Exit(0)
}

func TestSimple() {
	timeout := 5 // set runtime 5 detik
	ch := make(chan bool)

	go timer(timeout, ch)
	go watch(timeout, ch)

	var input string
	fmt.Println("Berapa hasil 5+6")
	fmt.Scanln(&input)

	if input == "11" {
		fmt.Println("Anda benar")
	} else {
		fmt.Println("Anda salah")
	}
}