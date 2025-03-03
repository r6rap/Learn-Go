package function

import(
	"fmt"
	"math/rand"
	"time"
)

/* Return function: function yang mengembalikan nilai.
	cara penulisannya dengan cara menuliskan tipe data return setelah kurung parameter */

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

