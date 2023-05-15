package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	hellHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, World!\n")
	}

	http.HandleFunc("/", hellHandler)

	log.Println("server start at port 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))

}
