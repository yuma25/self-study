package main

import (
	"fmt"
	"io"
	"strings"
)

// 1. fmt.Stringer インターフェースの実装
// これを実装すると、fmt.Println 等で独自の表示形式が使われる
type Book struct {
	Title  string
	Author string
}

func (b Book) String() string {
	return fmt.Sprintf("『%s』(%s 著)", b.Title, b.Author)
}

// 2. io.Reader インターフェースの実装
// Go の中で最も重要なインターフェースの一つ
// Read(p []byte) (n int, err error) を実装する
type UpperReader struct {
	src string
	pos int
}

func (r *UpperReader) Read(p []byte) (n int, err error) {
	if r.pos >= len(r.src) {
		return 0, io.EOF // 読み終わり
	}

	// 読み込める分だけ読み込む
	n = copy(p, strings.ToUpper(r.src[r.pos:]))
	r.pos += n
	return n, nil
}

func main() {
	fmt.Println("--- Go: Standard Interfaces ---")

	// Stringer の動作確認
	book := Book{Title: "Go言語の設計思想", Author: "Rob Pike"}
	fmt.Println("Book info:", book)

	// io.Reader の動作確認
	// どんな入力ソースも io.Reader という「共通の口」で扱える
	reader := &UpperReader{src: "hello world",pos : 4}
	
	// io.ReadAll は io.Reader インターフェースを受け取る関数
	// 自作の UpperReader が標準関数でそのまま使える！
	content, _ := io.ReadAll(reader)
	fmt.Printf("Read all: %s\n", content)
}

/**
 * 補足:
 * Go の凄さは、標準ライブラリの多くが interface で定義されている点です。
 * 自分で io.Reader や io.Writer を実装すれば、ファイル、ネットワーク、
 * 暗号化、圧縮などの標準機能を自分のデータ型に対して即座に適用できます。
 * 
 * 実行コマンド:
 * cd go_study && go run 13_standard_interfaces/main.go
 */
