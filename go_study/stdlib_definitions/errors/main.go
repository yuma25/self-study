package main

import (
	"errors"
)

/*
   =============================================================================
   Go 標準ライブラリ 定義集: errors パッケージ (究極網羅・実戦コード例版)
   =============================================================================
*/

// --- 1. 基本インターフェース ---

// 使い方: type MyError struct{} のように構造体を定義し，Error() string を実装して独自のエラー型を作る．
type MyErrorInterface interface {
	Error() string
}

// --- 2. パッケージレベルの関数と定数 ---

// 使い方: if errors.Is(err, errors.ErrUnsupported) { ... } と書き，機能が未実装（サポートされていない）であることを示す標準的なエラーを検知する（Go 1.21+）．
// 自作関数で「この操作はこの型には対応していない」と伝えたい場合に使う．
var MyErrUnsupported error = errors.ErrUnsupported

// 使い方: err := errors.New("file not found") と書き，単純なテキストメッセージを持つ新しいエラーを生成する．
// パッケージのグローバル変数（ErrNotFound 等）を定義する際によく用いられる．
func MyNew(text string) error { return nil }

// 使い方: if errors.Is(err, os.ErrNotExist) { ... } と書き，エラーチェーン（ラップされたエラー）の中に，特定の基準エラーが含まれているかを判定する．
// 従来の実体比較 (err == os.ErrNotExist) に代わる，Go 1.13以降の標準的なエラー検査方法である．
func MyIs(err, target error) bool { return false }

// 使い方: var pathErr *os.PathError; if errors.As(err, &pathErr) { fmt.Println(pathErr.Path) } と書き，エラーを特定の型に安全にキャスト（抽出）する．
// 単なるエラーメッセージの比較ではなく，エラー構造体が持つ詳細な情報（ステータスコード等）を取り出したい時に必須である．
func MyAs(err error, target any) bool { return false }

// 使い方: originalErr := errors.Unwrap(err) と書き，ラップされた（fmt.Errorf("%w")等で作られた）エラーから，一つ内側の元エラーを取り出す．
// 通常は Is や As が内部で再帰的に呼び出すため，開発者が直接呼ぶ機会は少ない．
func MyUnwrap(err error) error { return nil }

// 使い方: err := errors.Join(err1, err2) と書き，複数の独立したエラーを1つに束ねる（Go 1.20+）．
// 複数のバリデーションエラーをまとめてクライアントに返したい場合などに活用する．
func MyJoin(errs ...error) error { return nil }

// --- 3. 関連テクニックとパターン ---

// 使い方: fmt.Errorf("read failed: %w", err) と書き，元となるエラーを「包んで（ラップして）」新しい情報を付加する．
// %w を使うことで，errors.Is や errors.As が内部のエラーまで再帰的に探索できるようになる．

/*
   [カスタムエラーの実装例]
   type QueryError struct {
       Query string
       Err   error
   }

   func (e *QueryError) Error() string { return e.Query + ": " + e.Err.Error() }

   // 独自のエラーでラップをサポートする場合，Unwrap() を実装する
   func (e *QueryError) Unwrap() error { return e.Err }
*/

/*
   [高度なエラー検査]
   - errors.Join で束ねられたエラーに対して errors.Is/As を呼び出すと，
     束ねられた全てのエラーを再帰的にチェックする．
   - 例: 3つのエラーを Join した場合，そのうちのどれか1つでもターゲットに合致すれば Is は true を返す．
*/

/*
   [注意: Error vs Panic]
   - Go の Error は「予測される異常系」であり，呼び出し側で対処することを期待する．
   - Panic は「回復不能な致命的エラー」や「プログラミングミス」であり，errors パッケージの対象外．
   - errors を使う目的は，panic を避け，プログラムの制御を維持することにある．
*/

func main() {
	println("Loaded: errors spec.")
}
