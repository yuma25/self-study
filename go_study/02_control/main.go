package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("--- Go Control Structures ---")

	// 1. If with init statement (公式: Effective Go > If)
	// 短い宣言で変数を作り、そのif/else内だけで使う
	if num := 10; num%2 == 0 {
		fmt.Printf("%d は偶数です\n", num)
	}

	// 2. For as While (公式: A Tour of Go > For)
	// Goにwhileはなく、forの条件式だけで表現する
	count := 0
	for count < 3 {
		fmt.Println("Count (while-style):", count)
		count++
	}

	// 3. For Range (公式: Effective Go > For)
	nums := []int{10, 20, 30}
	for index, value := range nums {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// 4. Switch without condition (公式: Effective Go > Switch)
	// 条件を書かないswitchは if-else if よりも読みやすい
	hour := time.Now().Hour()
	fmt.Printf("Current hour: %d\n", hour)
	switch {
	case hour >= 5 && hour < 12:
		fmt.Println("おはよう")
	case hour < 18:
		fmt.Println("こんにちは")
	default:
		fmt.Println("こんばんは")
	}
}
