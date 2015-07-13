package main

import (
	"log"
	"net/http"

	nyx "shpsec.com/nyx"
)

func main() {
	router := nyx.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
