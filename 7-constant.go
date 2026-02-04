package main

import "fmt"

func main() {
	// deklarasi constant dengan kata kunci const
	const firstName string = "Muhamad"
	const lastName = "Azhar Luthfiadi"
	const age = 20
	fmt.Println("nama lengkap:", firstName, lastName)
	fmt.Println("umur:", age)

	// mencoba mengubah nilai constant akan menyebabkan error
	// firstName = "Luthfiadi" // ini akan error

	// sama seperti variabel, constant juga bisa dideklarasikan secara bersamaan
	const (
		city = "Bandung"
		country = "Indonesia"
	)

	fmt.Println("kota:", city, "negara:", country)
}