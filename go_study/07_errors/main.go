package main

import (
	"fmt"
)

// 1. カスタムエラーの定義 (公式: Effective Go > Errors)
// errorインターフェース (Error() string を持つ) を満たす構造体
type MyError struct {
	Code    int
	Message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func doSomething(fail bool) error {
	if fail {
		return &MyError{Code: 404, Message: "Not Found"}
	}
	return nil
}

// 2. defer によるクリーンアップ (公式: Effective Go > Defer)
// 関数の終了時に必ず実行される。リソースの解放（ファイルクローズ等）に使う
func process() {
	fmt.Println("Start processing...")
	defer fmt.Println("Clean up: process finished!") // 関数の最後に実行される

	err := doSomething(true)
	if err != nil {
		fmt.Println("Caught error:", err)
		return
	}
	fmt.Println("Success!")
}

func main() {
	fmt.Println("--- Go Errors and Defer ---")
	process()

	// 3. Panic と Recover (最終手段)
	// 基本的にはエラーを戻り値として扱うべきだが、どうしても継続不可能な場合に使用
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Triggering panic...")
	panic("Something went horribly wrong!")
	fmt.Println("This line will not be executed")
}
