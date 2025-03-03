package function

import (
	"fmt"
	"strings"
)

/* Void function: function yang tidak mengembalikan nilai apapun*/

func PrintMessage(message string, arr []string) { //void function
	// strings.Join() digunakan untuk menggabungkan element slice menjadi satu string tunggal dengan pembatas spasi
	name := strings.Join(arr, " ") 
	fmt.Println(message, name)
}

func divideNumber(m, n int) {
	if n == 0 {
		fmt.Printf("Invalid divider. %d cannot divide by %d\n", m, n)
		return /*jika tidak menggunakan return maka akan melanjutkan ke perhitungan pembagian yang akan
				 menyebabkan error */
	}

	calc := m / n
	fmt.Printf("%d / %d = %d\n", m ,n , calc)
}

func Divide() {
	divideNumber(10, 2)
	divideNumber(20, 0)
	divideNumber(15, -3)
}