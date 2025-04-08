package reflects

import (
	"fmt"
	"reflect"
)

/*	Reflect adalah package bawaan go yang bisa digunakan untuk menginpeksi variabel, mengambil informasi dari
	suatu variabel untuk dilihat meta datanya atau untuk keperluan manipulasi */

type person struct {
	name string
	age int
}

func Reflect() {
	var unknown any = "Hello World"

	// reflect.ValueOf() digunakan untuk mengambil value dari variabel
	// reflect.TypeOf() digunakan untuk mengambil tipe data dari variabel
	fmt.Println(reflect.ValueOf(unknown)) // output Hello World
	fmt.Println(reflect.TypeOf(unknown)) // output string
}

func Change() {
	/*	Untuk mengubah nilai, kamu harus mendapatkan pointer dan menggunakan .Elem() agar bisa mengakses nilai dari pointer.*/
	var x int = 20
	fmt.Println("Nilai x sebelum diubah:", x)

	// gunakan pointer dan reflect.Value agar bisa diubah
	val := reflect.ValueOf(&x).Elem()
	val.SetInt(40)
	fmt.Println("Nilai x setelah diubah:", x)
}

func (p *person) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(p) // reflect value dari variabel p diambil

	// .Kind() digunakan untuk mendapatkan jenis data
	if reflectValue.Kind() == reflect.Pointer { // mengecek apakah reflectValue adalah pointer
		reflectValue = reflectValue.Elem() // jika true maka akan mengambil nilai asli dari pointer itu
	}

	var reflectType = reflect.TypeOf(p) // reflect type dari variabel p diambil

	if reflectType.Kind() == reflect.Pointer {
		reflectType = reflectType.Elem()
	}

	// melakukan iterasi sebanyak jumlah field dalam struct person
	for i := 0; i < reflectType.NumField(); i++ {
		fmt.Println("Nama field:", reflectType.Field(i).Name)
		fmt.Println("Tipe data:", reflectType.Field(i).Type)
		fmt.Println("Nilai:", reflectValue.Field(i))
	}
}

func Reflects() {
	p1 := &person{name: "Rafif", age: 19}
	p1.getPropertyInfo()
}