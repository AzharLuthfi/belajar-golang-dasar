package main

import "fmt"

func main() {
	// 1. deklarasi variabel dengan tipe data string
	// var name string; 

	// 2. bisa juga tanpa tipe data tapi langsung diisi nilai
	// var name = "Muhamad Azhar Luthfiadi"

	// 3.cara lain deklarasi variabel tanpa var dan tipe data dengan menggunakan :=
	name := "Muhamad Azhar Luthfiadi"      

	name = "Muhamd Azhar Luthfiadi";
	fmt.Println("ini isi variabel pertama:", name);

	name = "Azhar Luthfiadi";
	fmt.Println("ini isi variabel setelah diganti:", name);

	// multiple variabel
	var (firstName = "Muhamad"
		lastName  = "Azhar Luthfiadi"
		age       = 20)

	fmt.Println("nama lengkap:", firstName, lastName)
	fmt.Println("umur:", age)

}