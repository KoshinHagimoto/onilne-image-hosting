package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/upload", uploadHandler).Methods("POST")
	r.HandleFunc("/image/{mediaType}/{uniqueString}", viewHandler).Methods("GET")
	r.HandleFunc("/delete/{uniqueString}", deleteHandler).Methods("GET")

	http.ListenAndServe(":8080", r)
}
