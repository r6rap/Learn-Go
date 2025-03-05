package function

import "fmt"

/* Immediately-Invoked Function Expression (IIFE)
	Closure jenis IIFE ini eksekusinya adalah langsung saat dideklarasikan,
	dengan menggunakan IIFE function akan langsung dipanggil, sehingga kode di dalamnya
	langsung dieksekusi dan variabel lokal di dalamnya tidak mencemari scope luar */

func IIFE() {
	data := []int{5, 3, 8, 1, 9, 2}

	min, max := func (num... int) (int, int) {
		min, max := num[0], num[0]
		for _, v := range num {
			if v < min {
				min = v
			}
			if v > max {
				max = v
			}
		}
		return min, max
	}(data...) //ciri IIFE

	fmt.Println("Nilai min:", min)
	fmt.Println("Nilai max:", max)
}