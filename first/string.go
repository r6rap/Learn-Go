package main

import "fmt"

func hitungHuruf(kalimat string) int {
	return len(kalimat)
}

func main(){
	fmt.Println("Jumlah huruf: ",hitungHuruf("rafif"))
}