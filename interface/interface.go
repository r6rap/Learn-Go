package interfaces

import (
	"fmt"
	"math"
)

/*	Interface adalah suatu tipe data abstrak yang mendefinisikan sekumpulan method tanpa implementasi.
	Interface menentukan "kontrak" yang harus dipenuhi oleh tipe data yang mengimplementasikan method method
	tersebut.
	Setiap tipe yang memiliki semua method yang didefinisikan di dalam interface akan dianggap secara
	otomatis mengimplementasikan interface tersebut.
*/

/*	Interface hitung mendefinisikan "kontrak" bahwa setiap tipe data yang mengimplementasikannya
	harus memiliki method luas() dan keliling() */
type hitung interface {
	luas() float64
	keliling() float64
}

type persegi struct {
	sisi float64
}

type lingkaran struct {
	jari_jari float64
}

// implementasi method luas() untuk struct persegi
func (p persegi) luas() float64 {
	return p.sisi * p.sisi
}

// implementasi method keliling() untuk struct persegi
func (p persegi) keliling() float64 {
	return 4 * p.sisi
}

// implementasi method luas() untuk struct lingkaran
func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jari_jari, 2)
}

// implementasi method keliling() untuk struct lingkaran
func (l lingkaran) keliling() float64 {
	return 2 * math.Pi * l.jari_jari
}

func (l lingkaran) diameter() float64 {
	return 2 * l.jari_jari
}

func Interface() {
 /* deklarasi variabel bangunDatar bertipe interface hitung */
	var bangunDatar hitung

 /* bisa mengisi bangunDatar dengan nilai dari struct persegi
 	karena persegi mengimplementasikan hitung */
	bangunDatar = persegi{5}
	fmt.Printf("Luas persegi dengan sisi 5: %.2f\n", bangunDatar.luas())
	fmt.Printf("Keliling persegi dengan sisi 5: %.2f\n", bangunDatar.keliling())

	fmt.Println("======================================")

 /* bisa mengisi bangunDatar dengan nilai dari struct lingkaran
 	karena lingkaran mengimplementasikan hitung */
	bangunDatar = lingkaran{14}
	fmt.Printf("Luas lingkaran dengan jari-jari 14: %.2f\n", bangunDatar.luas())
	fmt.Printf("Keliling lingkaran dengan jari-jari 14: %.2f\n", bangunDatar.keliling())

 // method diameter() tidak ada dalam interface hitung, jadi harus dikonversi ke tipe lingkaran
	fmt.Printf("Diameter lingkaran dengan jari-jari 14: %.2f\n", bangunDatar.(lingkaran).diameter())
}