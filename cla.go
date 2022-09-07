package main

import (
	"fmt"
	"os"
)

type student struct {
	id        int
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {

	var nama = []string{"Hari", "Thomas", "Brian", "Fitri"}
	var alamat = []string{"Jalan Tentara Pelajar No 1", "Jalan Tentara Pelajar No 2", "Jalan Tentara Pelajar No 3"}
	var pekerjaan = []string{"Programmer", "Programmer", "Programmer"}
	var alasan = []string{"Make money", "Better life future", "Better life future"}

	args := os.Args[1]

	var std student
	var found int = 0

	for i := 0; i < len(nama); i++ {

		if nama[i] == args {
			std.id = i
			std.nama = nama[i]
			std.alamat = alamat[i]
			std.pekerjaan = pekerjaan[i]
			std.alasan = alasan[i]

			found = 1
		}

	}

	if found == 1 {
		fmt.Println("Data ditemukan!")
		fmt.Println("ID : ", std.id)
		fmt.Println("Nama : ", std.nama)
		fmt.Println("Alamat : ", std.alamat)
		fmt.Println("Pekerjaan : ", std.pekerjaan)
		fmt.Println("Alasan : Alasan ", args, "menjadi", std.pekerjaan, "adalah untuk ", std.alasan)
	} else {
		fmt.Println("Data tidak ditemukan!")
	}

}
