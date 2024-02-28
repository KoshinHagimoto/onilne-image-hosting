package object

import "time"

type Image struct {
	ID           int       `json:"id"`            // データベースのID
	MediaType    string    `json:"media_type"`    // 画像のメディアタイプ
	UniqueString string    `json:"unique_string"` // 画像の一意の文字列
	FilePath     string    `json:"file_path"`     // 画像のファイルパス
	Created_at   time.Time `json:"created_at"`    // 画像の作成日時
}
