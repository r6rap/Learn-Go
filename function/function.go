package function

import(
	"fmt"
	"strings"
	"math/rand"
	"time"
)
/*	Ada 2 tipe function
	1. Void function: function yang tidak mengembalikan nilai apapun
	2. Return function: function yang mengembalikan nilai */

func PrintMessage(message string, arr []string) { //void function
	// strings.Join() digunakan untuk menggabungkan element slice menjadi satu string tunggal dengan pembatas spasi
	name := strings.Join(arr, " ") 
	fmt.Println(message, name)
}

func MathRandom() {
	
}