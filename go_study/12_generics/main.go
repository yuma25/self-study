package main

import (
	"fmt"
	"strings"
)

// 1. Generics を使った汎用的な Filter 関数 (Go 1.18+)
// [T any] は「どんな型でも受け取れる」という意味
func Filter[T any](slice []T, f func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// 2. Generics を使った汎用的な Map 関数
// T 型のスライスを U 型のスライスに変換する
func Map[T any, U any](slice []T, f func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = f(v)
	}
	return result
}

type User struct {
	Name string
	Role string
}

func main() {
	fmt.Println("--- Go: Generics and Functional-style ---")

	users := []User{
		{Name: "Alice", Role: "Admin"},
		{Name: "Bob", Role: "User"},
		{Name: "Charlie", Role: "Admin"},
	}

	// 1. Filter: Admin だけを抽出
	admins := Filter(users, func(u User) bool {
		return u.Role == "Admin"
	})

	// 2. Map: 名前だけを抽出して大文字に変換
	adminNames := Map(admins, func(u User) string {return strings.ToUpper(u.Name)})

	fmt.Println("Admin Names (Upper):", len(adminNames))
}

/**
 * 補足:
 * Go の標準ライブラリ（slicesパッケージ）にも最近同様の機能が追加されましたが、
 * 自分で Generics を使ってこれらを実装できるようになると、
 * 型安全で再利用性の高いコードが書けるようになります。
 * 
 * 実行コマンド:
 * cd go_study && go run 12_generics/main.go
 */
