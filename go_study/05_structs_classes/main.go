package main

import "fmt"

// 1. 構造体の定義 (公式: Effective Go > Allocation with make)
type User struct {
	ID   int
	Name string
}

// 2. メソッドの定義 (公式: A Tour of Go > Methods)
// (u User) はレシーバと呼ばれる。値型レシーバはコピーを渡す
func (u User) Greet() {
	fmt.Printf("Hi, I'm %s (ID: %d)\n", u.Name, u.ID)
}

// (u *User) はポインタレシーバ。元の構造体を書き換えられる
func (u *User) UpdateName(newName string) {
	u.Name = newName
}

// 3. 構造体の埋め込み (Composition/Embedding)
// Goには継承がないため、他の構造体を埋め込んで機能を再利用する
type Admin struct {
	User  // User構造体のフィールドとメソッドを引き継ぐ
	Level int
}

func main() {
	fmt.Println("--- Go Structs ---")

	// 初期化
	u := User{ID: 1, Name: "Alice"}
	u.Greet()

	// ポインタレシーバでの変更
	u.UpdateName("Bob")
	u.Greet()

	// 埋め込みを使ったAdmin
	a := Admin{
		User:  User{ID: 99, Name: "Superuser"},
		Level: 10,
	}
	// Userのメソッドを直接呼べる
	a.Greet()
	fmt.Println("Admin Level:", a.Level)
}
