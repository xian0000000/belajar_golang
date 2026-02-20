package fungsi

import (
	"belajar_golang/mini_cli/data"
	"belajar_golang/mini_cli/user"
	"fmt"
	"os"
	"os/exec"
)

func Setor_Uang() {
	var jumlah_setor float64

	fmt.Println("masukan jumlah yang ingin di setor : ")
	fmt.Scanln(&jumlah_setor)
	data.Setor(jumlah_setor)
}

func Tarik_Uang() {
	var jumlah_tarik float64

	fmt.Println("masukan jumlah yang ingin di Tarik : ")
	fmt.Scanln(&jumlah_tarik)
	data.Tarik(jumlah_tarik)
}

func lobby() int {
	var opsi int

	fmt.Println("pilih opsi")
	fmt.Println("1.Setor Uang")
	fmt.Println("2.Tarik Uang")
	fmt.Println("masukan opsi :")
	fmt.Scanln(&opsi)
	return opsi
}

func clearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Main_Loop() {
	clearTerminal()
	data.Load()
	nama, password := user.Login()
	if nama != data.Data_d.User_d.Nama || password != data.Data_d.User_d.Password {
		fmt.Println("maaf nama dan password salah")
		return
	}
	fmt.Println("login berhasil")
	run := true
	for run {
		clearTerminal()
		fmt.Println(data.Data_d.User_d.Nama)
		fmt.Println("selamat datang di bank digital")
		data.Saldo_Bank()
		opsi := lobby()

		if opsi == 1 {
			Setor_Uang()
		} else if opsi == 2 {
			Tarik_Uang()
		}

	}

}
