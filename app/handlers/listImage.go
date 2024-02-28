package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"online-image/app/object"
)

func ListImageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var images []object.Image

		rows, err := db.Query("SELECT id, media_type, unique_string, file_path, created_at FROM images")
		if err != nil {
			http.Error(w, "Error retrieving the images", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var img object.Image
			if err := rows.Scan(&img.ID, &img.MediaType, &img.UniqueString, &img.FilePath, &img.Created_at); err != nil {
				http.Error(w, "Error retrieving the images", http.StatusInternalServerError)
				return
			}
			images = append(images, img)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(images)
	}
}
