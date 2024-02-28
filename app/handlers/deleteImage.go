package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	// 画像IDをURLから取得
	vars := mux.Vars(r)
	imageId := vars["uniqueString"]

	imagePath := filepath.Join(uploadDir, imageId)

	// ファイルが存在するか確認
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		http.Error(w, "Image not found", http.StatusNotFound)
		return
	}

	//ファイルを削除
	err := os.Remove(imagePath)
	if err != nil {
		http.Error(w, "Error deleting the file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Image %s has been deleted", imageId)
}
