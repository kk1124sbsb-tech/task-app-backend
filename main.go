package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// ResponseData は、返すJSONの構造を定義する
type ResponseData struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

// helloHandler は、リクエストが来た時の処理
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 1. 返したいデータを作る
	data := ResponseData{
		Message: "Hello from Go Backend!",
		Status:  "success",
	}

	// 2. データの形式を「JSON」に指定する
	w.Header().Set("Content-Type", "application/json")

	// 3. データをJSONに変換して返す
	json.NewEncoder(w).Encode(data)
}

func main() {
	// 「/api/hello」というURLにアクセスが来たら、helloHandlerを実行する
	http.HandleFunc("/api/hello", helloHandler)

	// サーバーをポート8080で起動する
	log.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}