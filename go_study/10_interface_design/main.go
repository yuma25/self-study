package main

import (
	"example.com/study/10_interface_design/models"
	"fmt"
)

/**
 * 1. インターフェースの定義
 * Goの設計思想: インターフェースは「使う側（Consumer）」で定義するのが理想的。
 * これにより、modelsパッケージがこのインターフェースを知らなくても疎結合になれる。
 */
type Describer interface {
	GetDescription() string
}

func PrintSomething(d Describer) {
	fmt.Println("PrintSomething:", d.GetDescription())
}

func main() {
	fmt.Println("--- Go Advanced: Package, Struct, Method, Interface ---")

	// 1. コンストラクタ関数を通じて初期化
	u := models.NewUser(1, "Alice", "Top-Secret-Password")

	// 2. フィールドへのアクセス
	fmt.Println("User Name:", u.Name) // Name は大文字なのでアクセス可能
	// fmt.Println(u.secret) // secret は小文字なのでコンパイルエラーになる

	// 3. メソッドの呼び出し
	fmt.Println("Method Call:", u.GetDescription())

	// 4. インターフェースによる抽象化
	// models.User は GetDescription() を持っているので、自動的に Describer を満たす
	PrintSomething(u)

	// 5. 標準ライブラリのインターフェース
	// fmt.Println は Stringer インターフェース (String() string) を持っていればそれを使う
	fmt.Println("Stringer usage:", u)
}
/**
 * 実行方法 (go.modがある go_study ディレクトリで実行):
 * go run 10_interface_design/main.go
 */
