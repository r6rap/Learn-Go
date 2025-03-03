package function

import "fmt"

/* di go function bisa dijadikan sebagai parameter untuk function lain. Syarat untuk menjadikan function sebagai
	parameter strukturnya harus sama */


// type declaration digunakan untuk mendefinisikan tipe data baru dengan menggunakan keyword type
type Filter func(string) string //Filter adalah alias dari func(string) string

func nameFilter (name string, filter Filter) string {
	return "Hello " + filter(name) //filter(name) sama seperti filterName(name)
	//sebelum menampilkan return kita filter dulu di func filterName
}

//struktur func filterName sama seperti func filter
func filterName(nama string) string { //func ini akan memfilter parameter name dari func nameFilter
	sensor := []string{"Anjing", "Cok", "Babi"}

	for _, v := range sensor{
		if nama == v {
			return "..."
		}
	}

	return nama
}

func FilteredName() {
	fmt.Println(nameFilter("Rafif", filterName))

	filter := filterName //function as value
	fmt.Println(nameFilter("Anjing", filter))
}