package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
)

const uploadDir = "../../pkg/storage/"

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	// gorilla/muxを使ってパスパラメータを取得
	vars := mux.Vars(r)
	imageId := vars["uniqueString"]

	imagePath := filepath.Join(uploadDir, imageId)
	http.ServeFile(w, r, imagePath)
}
