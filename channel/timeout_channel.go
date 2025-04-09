package channel

import (
	"fmt"
	"time"
	"math/rand"
)

func sendData(ch chan <- int) {
	randomizer := rand.New(rand.NewSource(time.Now().Unix())) // menghasilkan seed acak

	for i := 0; true; i++ {
		ch <- i
		time.Sleep(time.Duration(randomizer.Int()%10+1) * time.Second) // menghasilkan delay 1-10 detik
	}
}

func receiverData(ch <- chan int) {
	loop:
	for {
		select {
		case msg := <- ch:
			fmt.Print(`Receive data "`, msg, `"`, "\n")
		case <- time.After(time.Second * 5): // case jika tidak ada data yang diterima oleh channel selama 5 detik
			fmt.Println("Timeout, no activities under 5 seconds") // print ini
		break loop // menghentikan perulangan
		}
	}
}

func TimeoutCH() {
	ch := make(chan int)
	go sendData(ch)
	receiverData(ch)
}