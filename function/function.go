package function

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

/*	Ada 2 tipe function
	1. Void function: function yang tidak mengembalikan nilai apapun
	2. Return function: function yang mengembalikan nilai
	3. Function multiple return: function yang memiliki banyak nilai balik/return
	4. Predefined return value function: function yang nilai baliknya sudah didefinisikan
	5. Variadic function: function yang bisa menerima banyak parameter seperti slice
	6. Closure function:  */

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

func calculate(d float64) (float64, float64) { //function multiple return
	luas := math.Pi * math.Pow(d / 2, 2) //pow digunakan untuk pangkat, (d/2)2 pangkat 2

	keliling := math.Pi * d //Pi = 22/7

	return luas, keliling
}

func Circle() {
	var diameter float64 = 14
	var luas, keliling = calculate(diameter)

	fmt.Printf("Luas lingkaran: %.2f \n", luas)
	fmt.Printf("Keliling lingkaran: %.2f \n", keliling)
}

func calc(d float64) (luas float64, keliling float64) { //predefined return value function
	luas = math.Pi * math.Pow(d / 2, 2)
	keliling = math.Pi * d

	return //karena nilai balik sudah ditentukan di awal, untuk mengembalikan nilai cukup dengan return
}

func Circle2() {
	var d float64 = 14
	var luas, keliling = calc(d)

	fmt.Printf("Luas lingkaran: %.2f \n", luas)
	fmt.Printf("Keliling lingkaran: %.2f \n", keliling)
}

func average(num... int) float64 { //variadic function
	total := 0

	for _, numbers := range num{ //pengaksesannya bisa menggunakan for range
		total += numbers
	}

	avg := float64(total) / float64(len(num)) //convert ke float menyesuaikan return type

	return avg
}

func Avg() {
	data := average(80, 60, 40, 60, 50, 55, 90, 55, 68)
	/*Sprintf sama seperti Printf hanya saja fungsi ini tidak menampilkan nilai, melainkan mengembalikan nilainya
	dalam bentuk string */
	msg := fmt.Sprintf("Rata rata: %.2f", data)

	fmt.Println(msg)
}

func Avg2() {
	data := []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3}
	hasil := average(data...) //slice bisa digunakan untuk argumen variadic function

	fmt.Printf("Rata rata: %.2f\n", hasil)
}

func hobbies(name string, hobbie...string) {
	hobi := strings.Join(hobbie, ", ")

	fmt.Println("Nama:", name)
	fmt.Println("Hobi:", hobi)
}

func Info() {
	nama := "Rafif Rizal"
	hobi := []string{"Belajar", "Lari", "Bersepeda"}

	hobbies(nama, hobi...)
}