package function

import(
	"fmt"
	"math"
)

/* Predefined return value function: function yang nilai baliknya sudah didefinisikan.*/

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