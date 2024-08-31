package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

func main() {
	jsonStream := []byte(`
	{"name":"Ganesh","Age":30,"Location":"Pune"}
	{"name":"Aniket","Age":32,"Location":"Nashik"}
	{"name":"Varun","Age":33,"Location":"Jadgaon"}
	`)

	reader := strings.NewReader(string(jsonStream))
	writet := os.Stdout

	decoder := json.NewDecoder(reader)
	encoder := json.NewEncoder(writet)

	for {
		var v map[string]interface{}

		if err := decoder.Decode(&v); err != nil {
			log.Println(err)
			return
		}

		for k:= range v {
			if k == "Age"{
				delete(v,k)
			}
		}

		if err := encoder.Encode(&v); err != nil {
			log.Println(err)
		}
	}
}
