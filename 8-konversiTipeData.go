package main

import "fmt"

func main() {
// int punya banyak tipe di Go, misalnya int8, int16, int32, int64
// kadang kita perlu melakukan konversi tipe data agar sesuai dengan kebutuhan kita
// hati hati, kita kalo conversi int ke tipe data yang lebih kecil, bisa menyebabkan overflow (terpotong nilainya)

	var x int64 = 300
	var y int8 = int8(x)

	fmt.Println(y) // ❗ hasilnya 44 karena terjadi overflow pada konversi tipe data dari int64 ke int8 
// karena nilai 300 melebihi batas maksimal int8 yaitu 127


// contoh dalam operator
	var a int = 10
	var b int64 = 20

// fmt.Println(a + b) // ❌ error
	fmt.Println(a + int(b)) // ✅ benar


// contoh konversi dari string ke int
	var name = "Azhar"
	var char = name[0]
	fmt.Println("char:", char) // hasilnya 65, karena yang ditampilkan adalah nilai ASCII dari karakter A

	var charString = string(char) // string(nilai)
	fmt.Println("charString:", charString) // hasilnya A, karena sudah dikonversi ke string
}