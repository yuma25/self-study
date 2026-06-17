package main

import "fmt"

func main() {
	fmt.Println("--- Go Data Structures ---")

	// 1. 配列 (Arrays) - 固定長、あまり直接は使われない
	// [長さ]型
	arr := [...]string{"Go", "TS", "Python"}
	fmt.Println("Array (fixed):", arr)

	// 2. スライス (Slices) - 可変長、Goで最も多用される
	// 公式: Effective Go > Slices
	// スライスは内部配列への「窓（ビュー）」のようなもの
	slice := []string{"Apple", "Banana"}
	slice = append(slice, "Cherry") // 要素の追加
	fmt.Println("Slice (dynamic):", slice)
	fmt.Println("Length:", len(slice), "Capacity:", cap(slice))

	// 部分取り出し (Slicing)
	fmt.Println("Partial:", slice[1:3]) // index 1から2まで

	// 3. マップ (Maps) - キー・バリュー
	// 公式: Effective Go > Maps
	// make(map[キーの型]値の型)
	scores := make(map[string]int)
	scores["Alice"] = 90
	scores["Bob"] = 85

	fmt.Println("Scores:", scores)

	// 要素の存在確認
	val, ok := scores["Charlie"]
	if ok {
		fmt.Println("Charlie's score:", val)
	} else {
		fmt.Println("Charlie is not found")
	}

	// 削除
	delete(scores, "Alice")
	fmt.Println("After delete:", scores)
}
