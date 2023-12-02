package main

import "fmt"

func main() {
	type nama string       //inisialisasi tipe data dengan nama baru
	var user nama = "Asep" //membuat variable asep dengan tipe data nama di isi dengan Asep
	fmt.Println("Hello", user)
	user = "Sahrul"
	fmt.Println("Hello", user)
}
