package main

import "fmt"

var nama string = "xina"
var number1 int = 0
var number2 float64 = 1.299
var run = true

func main() {
	fmt.Println(number2)

	if nama == "xina" {
		fmt.Println("selamat datang kembali", nama)
	} else {
		fmt.Println("get outttttttt")
	}

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	for run == true {
		fmt.Println(number1)
		if number1 < 10 {
			number1++
		} else {
			run = false
		}
	}
	hasil := tambah(10, 20)
	fmt.Println(hasil)
}

func tambah(a int, b int) int {
	return a + b
}
