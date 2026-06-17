package main

import (
	"fmt"
	"math"
)

// 1. インターフェースの定義 (公式: Effective Go > Interfaces)
// 必要なメソッドのシグネチャだけを定義する
type Shape interface {
	Area() float64
}

// 2. 実装: Goでは「implements」キーワードは不要 (Duck Typing)
// 構造体がインターフェースと同じメソッドを持っていれば、自動的にそのインターフェースを満たしているとみなされる
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

// 3. インターフェースを引数に取る関数
func PrintArea(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
}

func main() {
	fmt.Println("--- Go Interfaces ---")

	c := Circle{Radius: 5}
	s := Square{Side: 10}

	// Shapeインターフェースとして扱える
	PrintArea(c)
	PrintArea(s)

	// 空のインターフェース interface{} (any)
	// どんな型でも受け取れる (公式: Effective Go > The blank interface)
	var any interface{} = "Hello"
	fmt.Println("Any value:", any)
}
