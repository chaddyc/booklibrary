package main

import (
	"log"
	"net/http"
)

func main() {
	fs:= http.FileServer(http.Dir("./static")
	http.Handle(pattern:"/", fs)

	log.Printlin(v...: "Starting the application, listening on :8080...")
	err := http.ListenAndServe(addr:":8080", handler:nil)
	if err != nil {
		log.Fatal(err)
	}
}