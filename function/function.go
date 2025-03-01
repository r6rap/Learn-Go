package function

import(
	"fmt"
	"strings"
	"math/rand"
	"time"
)
/*	Ada 2 tipe function
	1. Void function: function yang tidak mengembalikan nilai apapun
	2. Return function: function yang mengembalikan nilai */

func PrintMessage(message string, arr []string) { //void function
	// strings.Join() digunakan untuk menggabungkan element slice menjadi satu string tunggal dengan pembatas spasi
	name := strings.Join(arr, " ") 
	fmt.Println(message, name)
}

//rand.New digunakan untuk membuat objek randomizer
var randomizer = rand.New(rand.NewSource((time.Now().Unix())))

func MathRandom() {
	var randomValue int

	randomValue = randomWithRange(2, 10)
	fmt.Println("Random number:", randomValue)

	randomValue = randomWithRange(2, 10)
	fmt.Println("Random number", randomValue)

	randomValue = randomWithRange(2, 10)
	fmt.Println("Random number", randomValue)
}

func randomWithRange(min, max int) int { //return function, menuliskan tipe data return setelah kurung parameter
	var value = randomizer.Int()%(max-min+1) + min
	return value
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