package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// 1. ログファイルの設定 (os.OpenFile)
	// 読み書き、作成、追記モードで開く
	logFile, err := os.OpenFile("server.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Log file error: %v\n", err)
		return
	}
	defer logFile.Close()

	// 2. インターフェースの結合 (io.MultiWriter)
	// 「標準出力(画面)」と「ログファイル」の両方に同時に書き込む魔法のWriterを作る
	mw := io.MultiWriter(os.Stdout, logFile)

	// ハンドラーの定義
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// すべてのリクエストを画面とファイルに記録
		fmt.Fprintf(mw, "Request received: %s %s\n", r.Method, r.URL.Path)
		
		fmt.Fprintln(w, "Welcome! This is an Interface-driven Web Server.")
	})

	// 3. ファイルストリーミング (io.Copy)
	// ファイルを読み込みながら、直接HTTPレスポンス(io.Writer)に流し込む
	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open("server.log")
		if err != nil {
			http.Error(w, "File not found", 404)
			return
		}
		defer file.Close()

		w.Header().Set("Content-Type", "text/plain")
		// file (io.Reader) から w (io.Writer) へ全データを流す！
		// メモリを節約しながら巨大なファイルも高速に送信できる
		io.Copy(w, file)
	})

	fmt.Println("Server starting at :8080...")
	fmt.Println("Endpoints:")
	fmt.Println("  - http://localhost:8080/     (Basic request)")
	fmt.Println("  - http://localhost:8080/log  (Stream server.log to browser)")
	
	http.ListenAndServe(":8080", nil)
}
