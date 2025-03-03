package function

import(
	"fmt"
	"math"
)

/* Function multiple return: function yang memiliki banyak nilai balik/return */

func luasLingkaran(d float64) (float64, float64) {
	luas := math.Pi * math.Pow( d/2, 2)
	keliling := math.Pi * d

	return luas, keliling
}

func Lingkaran() {
	var diameter float64 = 10
	luas, keliling := luasLingkaran(diameter)

	fmt.Printf("Luas: %.2f\n", luas)
	fmt.Printf("Keliling: %.2f\n", keliling)
}