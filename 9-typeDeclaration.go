package main

import "fmt"
func main() {
	// Type Declarations biasanya digunakan untuk membuat 
	// alias terhadap tipe data yang sudah ada, dengan tujuan 
	// agar lebih mudah dimengerti
	type NoKTP string
	type Married bool
	var noKTP NoKTP = "3275023405820001"
	var married Married = false
	fmt.Println("No KTP:", noKTP)
	fmt.Println("Married:", married)
	// kita juga bisa membuat Type Declarations dari tipe data yang sudah kita deklarasikan sebelumnya
	type StudentNoKTP NoKTP
	var studentNoKTP StudentNoKTP = "3275023405820002"
	fmt.Println("Student No KTP:", studentNoKTP)
}