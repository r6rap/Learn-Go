package array

import "fmt"

/*Perbedaan tipe nilai (value types) dan tipe referensi (reference type) 

*/
func Slice() {
	var fruits = []string{"avocado", "mango", "banana"}

	fmt.Println(fruits[1])
}