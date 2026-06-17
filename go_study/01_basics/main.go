package main

import (
	"fmt"
)

func main() {
	// 1. 変数宣言 (明示的)
	var message string = "Hello Go!"
	
	// 2. 型推論 (Short Variable Declaration)
	// 関数内でのみ使用可能
	version := 1.22

	// 3. 定数
	const language = "Go"

	// 4. ゼロ値 (初期化しない場合のデフォルト値)
	var count int
	var isActive bool

	fmt.Printf("--- %s Study ---\n", language)
	fmt.Println("Message:", message)
	fmt.Println("Version:", version)
	fmt.Println("Default Count:", count)   // 0 が表示される
	fmt.Println("Default Active:", isActive) // false が表示される
}
