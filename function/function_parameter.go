package function

import "fmt"
import "strings"

/* di go function bisa dijadikan sebagai parameter untuk function lain. Syarat untuk menjadikan function sebagai
	parameter strukturnya harus sama */


// type declaration digunakan untuk mendefinisikan tipe data baru dengan menggunakan keyword type
type Filter func(string) string //Filter adalah alias dari func(string) string

		//atau bisa langsung fliter Filter
func nameFilter (name string, filter func(string) string) string  {
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


func filter(data []string, callback func(string) bool) []string {
	var result []string
		for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}
   
func Min() {
	var data = []string{"wick", "jason", "ethan"}

	var dataContainsO = filter(data, func(yo string) bool { //yo berisi string dari slice data
		//proses filter yang akan menghasilkan true ataupun false
		return strings.Contains(yo, "o") //Contains untuk mengecek apakah substring(param2) bagian dari string(param1)
	})

	var dataLenght5 = filter(data, func(yo string) bool {
		return len(yo) == 5
	})

	fmt.Println("data asli \t\t:", data)
	// data asli : [wick jason ethan]

	fmt.Println("filter ada huruf \"o\"\t:", dataContainsO)
	// filter ada huruf "o" : [jason]

	fmt.Println("filter jumlah huruf \"5\"\t:", dataLenght5)
	// filter jumlah huruf "5" : [jason ethan]
   }