package main

import "fmt"

func main() {
	// // Tugas 01
	// fmt.Print("Halo, dunia!")

	// // Tugas 03
	// var A, B int

	// fmt.Scan(&A, &B)
	// fmt.Print(A + B)

	// // Tugas 04
	// var angka int32

	// fmt.Scan(&angka)
	// if angka > 0 {
	// 	fmt.Print(angka)
	// }

	// // Tugas 06
	// var kandang, jumlahBebek, i int32

	// fmt.Scan(&kandang)
	// for i = 0; i < kandang; i++ {
	// 	var bebek int32
	// 	fmt.Scan(&bebek)
	// 	jumlahBebek += bebek
	// }

	// fmt.Print(jumlahBebek)

	// // Tugas 07
	// var loop, j int8

	// fmt.Scan(&loop)
	// for j = 1; j <= loop; j++ {
	// 	if i%10 == 0 {
	// 		continue
	// 	}
	// 	if i == 42 {
	// 		fmt.Print("ERROR")
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	// Tugas 08
	// var loop, i int16
	// var angka, j int32
	// fmt.Scan(&loop)
	// for i = 0; i < loop; i++ {
	// 	fmt.Scan(&angka)
	// 	var checker bool = true
	// 	for j = 2; j < angka && j < 999; j++ {
	// 		if angka%j == 0 {
	// 			checker = false
	// 			break
	// 		}
	// 	}
	// 	if checker && angka != 1 {
	// 		fmt.Println("YA")
	// 	} else {
	// 		fmt.Println("BUKAN")
	// 	}
	// }

	// Tugas 09
	// var loop, i , maks int32
	// var banyak int16
	// var array = []int16{}
	// banyak = 0
	// maks = 0
	// fmt.Scan(&loop)
	// for i = 0; i < loop; i++ {
	// 	var angka int16
	// 	fmt.Scan(&angka)
	// 	if banyak == 0 {
	// 		banyak = angka
	// 	}
	// 	array = append(array, angka)
	// }
	// for i = 0; i < loop; i++ {
	// 	fmt.Print(array[i])
	// }

	// Tugas 10
	// var A, B, C int32
	// fmt.Scan(&A, &B)
	// C = reverse(reverse(A) + reverse(B))
	// fmt.Print(C)

	// Tugas 11
	// var S, T string
	// fmt.Scan(&S, &T)

	// for strings.Contains(S, T) {
	// 	index := strings.Index(S, T)
	// 	S = S[:index] + S[index+len(T):]
	// }

	// fmt.Println(S)

	// Tugas 12 faktorial !!
	// var N int32
	// var hasil int32
	// fmt.Scan(&N)

	// hasil = faktor(N)
	// fmt.Print(hasil)

	// Tugas 13 gambar gunung
	// var n int8
	// fmt.Scan(&n)
	// gambar(n)

}

func reverse(nilai int32) int32 {
	var temp int32
	var ret int32 = 0

	for temp = nilai; temp > 0; temp = temp / 10 {
		ret = (ret * 10) + (temp % 10)
	}

	return ret

}

func faktor(N int32) int32 {
	var output int32
	var pembagi int8
	if N%2 == 0 {
		pembagi = 2
	} else {
		pembagi = 1
	}
	if N <= 0 {
		output = 1
	} else {
		output = N / int32(pembagi) * faktor(N-1)
	}

	return output
}

func gambar(N int8) {
	var i int8
	if N > 0 {
		gambar(N - 1)
		for i = 1; i <= N; i++ {
			fmt.Print("*")
		}
		fmt.Println()
		gambar(N - 1)
	}
}
