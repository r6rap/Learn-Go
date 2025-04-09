package channel

import "fmt"

/*	Channel adalah media perantara / jalur komunikasi antara goroutine yang di mana goroutine a bisa mengirim data dan goroutine lain
	bisa menerima data dari channel tersebut. secara default channel dijalankan secara synchronous.
	jika tanda <- ditulis di sebelah kiri variabel, berarti sedang berlangsung pengiriman data.
	jika tanda <- ditulis di sebelah kiri channel, berarti sedang berlangsung penerimaan data
	jika hanya "ch chan string" maka channel tersebut bisa mengirim sekaligus menerima data */

func Channel() {
	var messages = make(chan string) // membuat channel dengan tipe data string

	var SayHello = func(name string) {
		var data = fmt.Sprintf("Halo %s", name)

		messages <- data // mengirim variabel data lewat channel messages
	}

	// running goroutine
	go SayHello("Rafif")
	go SayHello("Edbert")
	go SayHello("Lemon")

	// menerima data dari channel messages
	message1 := <- messages
	fmt.Println(message1)

	message2 := <- messages
	fmt.Println(message2)

	message3 := <- messages
	fmt.Println(message3)
}