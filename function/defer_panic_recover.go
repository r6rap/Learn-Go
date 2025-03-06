package function

import "fmt"

/*	Defer function adalah function yang bisa kita jadwalkan untuk dieksekusi
	setelah sebuah function dieksekusi
	Defer function akan selalu dieksekusi walaupun terjadi error di function yang dieksekusi
	
	Panic digunakan untuk menghentikan sebuah program
	Panic biasanya dipanggil ketika terjadi panic saat program kita berjalan
	Pada saat panic dipanggil semua program akan berhenti, namun defer function tidak
	
	Recover digunakan untuk menangkap data panic, dengan recover proses panic akan berhenti
	dan program tetap berjalan */

func logging() {
	fmt.Println("Done finished function")
	message := recover() //menyimpan data recover di variabel message
	fmt.Println("Terjadi panic/err di:", message)
}

//setelah RunApp dieksekusi maka function logging akan dieksekusi walaupun terjadi error di function RunApp
func RunApp() {
	defer logging()
	fmt.Println("Run App")
}

func Err(error bool) {
	defer logging()
	if error {
		panic("Error")
	}

	fmt.Println("Tidak dieksekusi") //jika error == true, print ini tidak akan dieksekusi
}

func myPanic() {
	panic("Panic")
}

//1
func Recover(error bool) {
	defer logging() //recover disimpan di defer agar bisa menangkap pesan dari panic

	if error {
		panic("Error")
	}
}

//2
func Recover2() {
	defer func() {
		if message := recover(); message != nil{
			fmt.Println("terjadi err di:", message)
		}
	}() //menggunakan IIFE karena function ini akan langsung dieksekusi setelah dideklarasikan

	myPanic()
}