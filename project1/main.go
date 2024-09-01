package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Configuration holds the server configuration settings.
type Configuration struct {
	ServerPort string `json:"serverport"`
}

func main() {
	// Open the configuration file
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close() // Ensure the file is closed after reading

	// Decode the configuration from the JSON file
	conf := new(Configuration)
	err = json.NewDecoder(file).Decode(conf)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Print the server port
	fmt.Println("Run Server .......", conf.ServerPort)

	// Run the server on the specified port
	runWeb(conf.ServerPort)
}

// runWeb starts a simple HTTP server on the specified port
func runWeb(port string) {
	fmt.Printf("Starting server on port %s...\n", port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome, Your IP is : %s\n", r.RemoteAddr)
	})

	// Start the HTTP server
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
