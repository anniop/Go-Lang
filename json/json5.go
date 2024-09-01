package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"net/http"
)

type Configuration struct {
	ServerPort string `json:"serverport"`
}

func main() {

	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	conf := new(Configuration)

	err = json.NewDecoder(file).Decode(conf)

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Run Server .......", conf.ServerPort)

	

	func RunWeb(address string) {

		http.HandleFunc("/", rootHandler)
	
		http.ListenAndServe(address, nil)
	
	}
	func rootHandler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome, Your IP is : %s\n", r.RemoteAddr)
	}

	RunWeb(conf.ServerPort)
}

