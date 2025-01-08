package main

import (
	"bytes"
	"encoding/json"
	"log"
	"strings"
)

type record struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

func Deserialize(e *json.Decoder, slice interface{}) error {
	return e.Decode(slice)
}

func Seriaize(e *json.Encoder, slice interface{}) error {
	return e.Encode(slice)
}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err == nil {
		log.Println(string(b))
	}
	return err
}

func JSONstream(data interface{}) (string, error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "\t")
	err := encoder.Encode(data)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

func main() {
	temp_slice := []record{}

	buf := strings.NewReader(`[{"key":"RESZD","value":63},{"key":"XUEYA","value":13},{"key":"FUNTIME","value":45}]`)
	err := Deserialize(json.NewDecoder(buf), &temp_slice)
	if err != nil {
		log.Fatal(err)
	}
	buf2 := new(bytes.Buffer)

	encoder := json.NewEncoder(buf2)
	err = Seriaize(encoder, temp_slice)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(PrettyPrint(temp_slice[0]))
	log.Println(JSONstream(temp_slice))
	return
}
