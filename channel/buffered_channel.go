package channel

import ( 
	"fmt"
	"time"
)
/*	Buffered channel adalah channel yang memiliki ruang (buffer) untuk menyimpan beberapa data sebelum
	data itu diambil oleh goroutine penerima */

func Buffered() {
	messages := make(chan int, 3) // membuat buffered channel dengan kapasitas 3

	go func() {
		for {
			i := <- messages
			fmt.Println("Menerima data ", i)
			time.Sleep(2 * time.Second)
		}
	}()

	for i := 0; i < 6; i++ {
		fmt.Println("Mengirim data", i)
		messages <- i
	}

	time.Sleep(7 * time.Second)
}

/*	Output:
data 0,1,2 masuk ke buffer, buffer[0,1,2]
Mengirim data 0
Mengirim data 1
Mengirim data 2

data 3,4 numpuk di pengirim karena buffer penuh
Mengirim data 3
Mengirim data 4

data 0,1 telah diterima, buffer[2,3,4]
Menerima data  0
Menerima data  1

data 5 numpuk di pengirim
Mengirim data 5

saat data 2 diterima, data 5 langsung masuk ke buffer[3,4,5]
Menerima data  2
Menerima data  3
Menerima data  4
Menerima data  5	
*/