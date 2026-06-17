package main

import (
	"example.com/study/14_final_project/taskmanager"
	"fmt"
)

func main() {
	fmt.Println("--- Go: Task Manager Final Project ---")

	m := taskmanager.NewManager()

	// 1. 追加
	m.Add("Buy milk")
	m.Add("Learn Go Generics")
	m.Add("Write final project")

	// 2. 完了処理 (エラーハンドリング)
	err := m.Complete(2)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// 存在しないID
	err = m.Complete(99)
	if err != nil {
		fmt.Printf("Expected Error: %v\n", err)
	}

	// 3. 一覧表示
	fmt.Println("\nCurrent Tasks:")
	m.List()
}
