package main

import "fmt"

func main() {
	type nama string         //inisialisasi tipe data dengan nama baru
	const user nama = "Asep" //membuat variable konstanta dimna variable ini tidak bisa diubah asep dengan tipe data nama di isi dengan Asep
	fmt.Println("Hello", user)
	user = "Sahrul"
	fmt.Println("Hello", user)
}
