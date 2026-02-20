package user

import (
	"fmt"
)

func Login() (string, string) {
	var nama string
	var password string

	fmt.Println("silahkan masukan nama dan password anda : ")
	fmt.Scanln(&nama)
	fmt.Scanln(&password)
	return nama, password
}
