package main

import (
	"net/http"
	"online-image/app/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/upload", handlers.UploadHandler).Methods("POST")
	r.HandleFunc("/image/{mediaType}/{uniqueString}", handlers.ViewHandler).Methods("GET")
	r.HandleFunc("/delete/{uniqueString}", handlers.DeleteHandler).Methods("GET")

	http.ListenAndServe(":8080", r)
}
