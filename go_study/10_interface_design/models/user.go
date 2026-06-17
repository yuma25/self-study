package models

import "fmt"

// User 構造体
// 大文字で始まるフィールドは外部からアクセス可能 (Public)
// 小文字で始まるフィールドはパッケージ内からのみアクセス可能 (Private)
type User struct {
	ID   int
	Name string
	secret string // 外部からは見えない
}

// 構造体の初期化を管理する「コンストラクタ関数」
// Goには言語仕様としてのコンストラクタはないため、New... という関数を作るのが慣習
func NewUser(id int, name string, secret string) *User {
	return &User{
		ID:     id,
		Name:   name,
		secret: secret,
	}
}

// User型に対するメソッド
func (u *User) GetDescription() string {
	// secret は同じパッケージ内なのでアクセス可能
	return fmt.Sprintf("User: %s (ID: %d), Secret: %s", u.Name, u.ID, u.secret)
}

// インターフェースを満たすための別のメソッド
func (u *User) String() string {
	return fmt.Sprintf("<User %s>", u.Name)
}
