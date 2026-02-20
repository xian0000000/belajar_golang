package main

import "fmt"

type Senjata struct {
	Suara string
}

var pedang = Senjata{"slashhhhhhh"}
var kapak = Senjata{"darrrrrrr"}

func (s Senjata) tebas() string {
	return s.Suara
}

type Serang interface {
	tebas() string
}

func main() {
	var s Serang

	s = pedang
	fmt.Println(s.tebas())
	s = kapak
	fmt.Println(s.tebas())
}
