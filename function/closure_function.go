package function

import "fmt"


/* 	Closure adalah kemampuan sebuah function untuk berinteraksi dengan data data di sekitarnya
di dalam scope yang sama */


func Closure() {
	counter := 0

	increment := func() {
		fmt.Println("Increment")
		counter++ // increment berinteraksi dengan data dari variabel counter
	}

	//counter menjadi 2 karena increment dipanggil 2 kali
	increment()
	increment()
	fmt.Println(counter)
}

func GetMinMax() {
	var min, max int

	minMax := func(n... int) (int, int) {
		for i, v := range n{
			switch {
			case i == 0:
				min, max = v, v
			case v < min:
				min = v
			case v > max:
				max = v
			}
		}
		return min, max
	}

	numbers := []int{2, 3, 4, 3, 4, 2, 3}
	min, max = minMax(numbers...) //insialisasi min, max dengan return dari closure minMax()
	fmt.Printf("data: %v\n min: %v\n max: %v\n", numbers, min, max) //%v digunakan untuk menampilkan data tanpa melihat tipe datanya
}

/*	Closure sebagai return, closure bisa dijadikan sebagai nilai kembalian atau return */

func newSequenceGenerator(start, step int) func() int {
	hasil := start - step

	//karna returnnya adalah closure function
	return func() int{
		hasil += step
		return hasil
	}
}

func New() {
	//maka hasil juga termasuk function sebagai value
	hasil := newSequenceGenerator(5, 3)
	fmt.Printf("hasil %d\n", hasil())
	fmt.Printf("hasil %d\n", hasil())
	fmt.Printf("hasil %d\n", hasil())
}