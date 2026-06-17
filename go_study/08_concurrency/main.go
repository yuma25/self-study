package main

import (
	"fmt"
	"time"
)

// 重い処理をシミュレートする関数
func task(name string, ch chan string) {
	fmt.Printf("Task %s started...\n", name)
	time.Sleep(2 * time.Second) // 2秒待機
	ch <- fmt.Sprintf("Task %s finished", name) // チャネルに結果を送信
}

func main() {
	fmt.Println("--- Go Concurrency (Goroutines & Channels) ---")
	start := time.Now()
	fmt.Println(start)

	// 1. チャネルの作成
	// 公式: Effective Go > Channels
	ch := make(chan string)

	// 2. ゴルーチンの起動 (go キーワード)
	// 公式: Effective Go > Goroutines
	go task("A", ch)
	go task("B", ch)

	fmt.Println("Waiting for tasks to finish...")

	// 3. チャネルからの受信 (結果が来るまでブロックされる)
	res1 := <-ch
	res2 := <-ch

	fmt.Println("Result 1:", res1)
	fmt.Println("Result 2:", res2)

	fmt.Printf("Total time: %v\n", time.Since(start))
}
