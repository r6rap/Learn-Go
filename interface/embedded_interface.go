package interfaces

import "fmt"

/*	Embedded interface adalah mekanisme untuk menanamkan interface ke interface lain tanpa harus
	mendefinisikan returnnya */

type penggerak interface {
	jalan() string
}

type pengereman interface {
	berhenti() string
}

type kendaraan interface {
	penggerak
	pengereman
}

type mobil struct {
	merk string
}

func (m mobil) jalan() string {
	return "mobil jalan"
}

func (m mobil) berhenti() string {
	return "mobil berhenti"
}

func EmbeddedInterface() {
	var k kendaraan = mobil{"BMW"}

	fmt.Println(k.berhenti())
	fmt.Println(k.jalan())
}