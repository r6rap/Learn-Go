package chapter1

import "fmt"

func VarTipeData()  {

	var firstName string = "Rafif"

	lastName := "Rizal"

	lastName = "Ali"

	one, isTuesday, twoPointTwo, sayIt := 1, true, 2.2, "Hai"

	_ = "hilang"

	nama := new(string)

	var positiveNumber uint8 = 70
	var negativeNumber = -364785782

	var floating = 2.62

	var exists = true

	var message = `"nama saya rafif", ini ditulis dengan backtip yang akan membuat semua karakter di dalamnya menjadi string`

	const goals = "membuat keluarga yang baik secara finansial maupun hubungan"
	//deklarasi variabel const secara bersamaan
	const(
		gender = "laki laki"
		isToday bool = true
		umur uint8 = 18
		random
	)

	fmt.Println("halo", firstName, lastName + "!")
	fmt.Println(one, isTuesday, twoPointTwo, sayIt)
	fmt.Println(nama)
	fmt.Printf("bilangan positif %d\n" , positiveNumber) //%d = format data numeric non desimal
	fmt.Printf("bilangan negatif %d\n", negativeNumber)
	fmt.Printf("bilangan decimal %f\n" , floating) //%f format data numeric decimal menjadi string
	fmt.Printf("custom digit decimal %.4f\n", floating)
	fmt.Printf("exist? %t \n", exists) //%t format data boolean
	fmt.Println(message)
	fmt.Print("my goals: ", goals, "\n")
	fmt.Println(random)
}