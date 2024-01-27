package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Tugas 01
	fmt.Print("Halo, dunia!")

	// Tugas 03
	var A, B int

	fmt.Scan(&A, &B)
	fmt.Print(A + B)

	// Tugas 04
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	kalimat := scanner.Text()
	fmt.Print(kalimat)

	// Tugas 05
	var bebek, temanPakDengklek int32

	fmt.Scan(&bebek, &temanPakDengklek)
	fmt.Println("masing-masing ", bebek/temanPakDengklek)
	fmt.Print("bersisa ", bebek%temanPakDengklek)

}
