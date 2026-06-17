package main

import (
	"errors"
	"fmt"
)

// 1. 基本的な関数 (引数と戻り値)
func add(a int, b int) int {
	return a + b
}

// 2. 複数の戻り値を返す関数 (Goの大きな特徴)
// 公式: Effective Go > Multiple return values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("0で割ることはできません")
	}
	return a / b, nil
}

// 3. 名前付き戻り値 (Named result parameters)
// 戻り値に名前を付けると、関数内で変数として使え、単に return と書くだけで返せる
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // x, y が返される
}

func main() {
	fmt.Println("--- Go Functions ---")

	// 基本
	fmt.Println("Add(10, 5):", add(10, 5))

	// 多値返しとエラーハンドリング
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Divide(10, 2):", result)
	}

	// エラーが発生する場合
	_, err2 := divide(10, 0)
	if err2 != nil {
		fmt.Println("Expected Error:", err2)
	}

	// 名前付き戻り値
	x, y := split(17)
	fmt.Printf("Split(17): x=%d, y=%d\n", x, y)
}
