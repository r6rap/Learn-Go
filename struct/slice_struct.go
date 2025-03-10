package structs

import(
	"fmt"
)

/*	Slice dan struct bisa dikombinasikan seperti pada slice dan map,
	caranya cukup menambahkan [] sebelum tipe data pada saat deklarasi */

type player struct{
	name string
	age int
}

func SliceStruct() {
	allPlayers := []player{
		{name: "Rap", age: 17},
		{name: "Pop", age: 19},
	}

	for i,v := range allPlayers{
		fmt.Println("Index:", i)
		fmt.Println("Name:", v.name+",", "age:", v.age)
	}
}