package main

import (
	// go.mod で定義したモジュール名 "example.com/study" からの相対パス
	"example.com/study/09_modules/mypackage"
	"fmt"
)

func main() {
	fmt.Println("--- Go Modules ---")
	mypackage.Hello()
}
