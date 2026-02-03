package main

import "fmt"

func main(){
	// fmt.Println("Nama saya azhar");
	// fmt.Println("Umur saya 21 tahun");
	// fmt.Println("saya mahasiswa");

	fmt.Println(len("Nama saya azhar")); //--> len() untuk menghitung jumlah karakter
	fmt.Println("Umur saya 21 tahun"[0]) //--> mengakses karakter pada index ke 3 (bentuknya byte, nanti harus di rubah ke string);
	fmt.Println("saya mahasiswa");


}