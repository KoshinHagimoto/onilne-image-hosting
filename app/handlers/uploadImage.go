package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// This is the path where the uploaded images will be stored.
const uploadPath = "../../pkg/storage/"

// This is a simple init function that checks if the uploadPath exists. If it doesn't, it creates it.
func init() {
	if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
		os.MkdirAll(uploadPath, os.ModePerm)
	}
}

// generateFileNameはランダムなファイル名を生成します。
func generateFileName() (string, error) {
	b := make([]byte, 16) //16バイトのランダム値を生成
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// uploadHandlerは画像を受け取り、保存します。
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// リクエストがPOSTであることを確認
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

	//ファイルを解析
	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// ファイル名を生成
	fileName, err := generateFileName()
	if err != nil {
		http.Error(w, "Error generating file name", http.StatusInternalServerError)
		return
	}
	fileName += filepath.Ext(handler.Filename) //ファイル名に拡張子を追加

	filePath := filepath.Join(uploadPath, fileName) //ファイルのパスを生成

	// ファイルを保存
	dst, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Error creating the file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// ファイルをコピー
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "Error saving the file", http.StatusInternalServerError)
		return
	}

	// ファイル名を返す
	fmt.Fprintf(w, "File Uploaded Successfully: %s", filePath)
}
