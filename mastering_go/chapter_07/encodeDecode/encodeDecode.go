package main

import (
	"encoding/json"
	"log"
)

type UseAll struct {
	Name    string `json:"username"`
	Surname string `json:"surname"`
	Year    int    `json:"created"`
}

func main() {
	useall := UseAll{Name: "Susan", Surname: "Hex", Year: 2024}
	t, err := json.Marshal(&useall)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Encoded JSON: %s\n", t)
	}
	// decode json
	json_string := `{"username": "SusanH", "surname":"hello", "created": 2024}`
	err = json.Unmarshal([]byte(json_string), interface{}(&useall))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Decoded JSON:", useall)
	error_json_string := `{"username": "sue", "surname":"hello", "Create": 2023},`
	log.Println("Attempting to unmarshal malformed json:")
	err = json.Unmarshal([]byte(error_json_string), interface{}(&useall))
	if err != nil {
		log.Println(err)
	}
	value, _ := json.Marshal(useall)
	log.Println(string(value))
}
