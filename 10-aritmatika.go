package main

import "fmt"
func main() {
	// list aritmatika di golang
	a := 10
	b := 3
	hasil := a + b
	if hasil%2 == 1 {
		fmt.Println("Hasil penjumlahan adalah bilangan ganjil")
	} else {
		fmt.Println("Hasil penjumlahan adalah bilangan genap")
	}
	// operasi aritmatika
	fmt.Println("Penjumlahan:", a+b)
	fmt.Println("Pengurangan:", a-b)
	fmt.Println("Perkalian:", a*b)
	fmt.Println("Pembagian:", a/b)
	fmt.Println("Sisa bagi (modulus):", a%b)
}