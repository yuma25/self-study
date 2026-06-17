package main

import "fmt"

type Product struct {
	Name      string
	Price     float64
	Available bool
}

func main() {
	fmt.Println("--- Go: Imperative Data Manipulation ---")

	products := []Product{
		{Name: "Laptop", Price: 1000, Available: true},
		{Name: "Mouse", Price: 20, Available: true},
		{Name: "Keyboard", Price: 50, Available: false},
		{Name: "Monitor", Price: 200, Available: true},
	}

	// シナリオ: 「在庫あり」の商品の価格を「10%引き」にし、その「合計金額」を出す
	
	// 1. 伝統的な Go の書き方 (For Loop)
	// メモリ効率が良く、デバッグが容易。
	var totalPrice float64
	for _, p := range products {
		if !p.Available {
			continue // Filter
		}
		
		discountedPrice := p.Price * 0.9 // Map
		totalPrice += discountedPrice    // Reduce
	}

	fmt.Printf("Total Price (Discounted): %.2f\n", totalPrice)

	// 2. Generics を使った抽象化 (Go 1.18+)
	// 関数を自作して TS の filter のように振る舞わせることも可能だが、
	// 標準ライブラリに filter/map がないのは「ループの方が速くて明確」という思想のため。
}

/**
 * 実行コマンド:
 * cd go_study && go run 11_data_manipulation/main.go
 */
