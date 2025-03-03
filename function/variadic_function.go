package function

import (
	"fmt"
	"strings"
)

/* Variadic function: function yang bisa menerima banyak parameter seperti slice.
	pengaksesan variadic function juga sama seperti slice yaitu menggunakan ... (ellipsis) */

func sum(nums... int) {
	fmt.Print(nums, " ")
	total := 0

	for _, v := range nums{
		total += v
	}

	fmt.Println(total)
}

func Sum() {
	sum(1,2,3,4)
	
	nums := []int{1,2,3,4}
	sum(nums...)
}

func dataDiri(nama string, hobi...string) string {
	data := strings.Join(hobi, ", ")

	return data
}

func DataDiri() {
	hobi := []string{"A", "B", "d"}

	fmt.Println(dataDiri("Pip", hobi...))
}

func average(nilai... int) float64 {
	sum := 0

	for _, v := range nilai{
		sum += v
	}

	avg := float64(sum) / float64(len(nilai))

	return avg
}

func Avg() {
	nilai := []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3}

	fmt.Printf("Rata Rata: %.2f\n",average(nilai...))
}