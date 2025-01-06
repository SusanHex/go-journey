package main

import (
	"encoding/json"
	"log"
)

type NoEmpty struct {
	Name    string `json:"username"`
	Surname string `json:"surname"`
	Year    int    `json:"creationyear,omitempty"`
}

type Password struct {
	Name    string `json:"username"`
	Surname string `json:"surname,omitempty"`
	Year    int    `json:"creationyear,omitempty"`
	Pass    string `json:"-"`
}

func main() {
	pass := Password{Name: "Susan", Surname: "Hex", Year: 2024, Pass: "Hello Beautiful World"}
	notEmpty := NoEmpty{}
	temp, _ := json.Marshal(pass)
	log.Println(string(temp))
	temp, _ = json.Marshal(notEmpty)
	log.Println(string(temp))
}
