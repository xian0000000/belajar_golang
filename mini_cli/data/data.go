package data

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Nama     string `json:"nama"`
	Password string `json:"password"`
}

type Bank struct {
	Saldo_Bank float64 `json:"saldo_bank"`
}

var Data_d Data

func Saldo_Bank() float64 {
	fmt.Println("total saldo bank :", Data_d.Bank_d.Saldo_Bank)
	return Data_d.Bank_d.Saldo_Bank
}

func Setor(jumlah_setor float64) {
	Data_d.Bank_d.Saldo_Bank += jumlah_setor
	Save(Data_d)
}

func Tarik(jumlah_setor float64) {
	Data_d.Bank_d.Saldo_Bank -= jumlah_setor
	Save(Data_d)
}

func Save(a any) error {
	file, err := os.Create("bank.json")
	if err != nil {
		return nil
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	return encoder.Encode(a)

}
func Load() error {
	file, err := os.Open("bank.json")
	if err != nil {
		return nil
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(&Data_d)
}

type Data struct {
	User_d User `json:"user"`
	Bank_d Bank `json:"bank"`
}
