package mypackage

import "fmt"

// Goのルール: 大文字で始まれば外部からアクセス可能 (Exported)
func Hello() {
	fmt.Println("Hello from mypackage!")
}
