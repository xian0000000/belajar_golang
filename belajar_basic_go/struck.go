package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	Nama      string `json:"nama"`
	Umur      int    `json:"umur"`
	Kemampuan string `json:"kemampuan"`
}

var data_json = `{
	"nama":"rian",
	"umur":100,
	"kemampuan":"menghilang"
}`

// var u Data = Data{
// 	"xina", 20, "mampu mengendalikan waktu",
// }

var angka = []any{10, "xina", 30}

var data = map[string]any{
	"nama":       "xina",
	"role":       "admin judol",
	"pendapatan": 99999999999,
	"target":     true,
}

func ubahnama(u *Data) {
	u.Nama = "iki"
}

func (u Data) salam() {
	fmt.Println("assalamualaikum", u.Nama)
}

func main() {
	u := Data{
		"xina", 20, "mampu mengendalikan waktu",
	}

	ubahnama(&u)

	fmt.Println(u)

	u.salam()

	angka = append(angka, 40)
	fmt.Println(angka)
	fmt.Println(angka[1])

	for i, v := range angka {
		fmt.Println(i+1, v)
	}

	angka = append(angka[:1], angka[2:]...)
	fmt.Println(angka)

	fmt.Println(data)
	for k, v := range data {
		fmt.Println(k, "=", v)
	}

	nama := data["nama"]

	fmt.Println(nama)

	var d Data
	json.Unmarshal([]byte(data_json), &d)
	fmt.Println(d)

	jsondata, _ := json.Marshal(u)

	fmt.Println(string(jsondata))

}
