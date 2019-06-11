package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	log.SetOutput(os.Stdout)
	log.Print("Starting http server")
	http.HandleFunc("/", HomePage)
	log.Fatal(http.ListenAndServe(":80", nil))
}
