package main

import (
	"strconv"
)

/*
   =============================================================================
   Go 標準ライブラリ 定義集: strconv パッケージ (究極網羅・実戦コード例版)
   =============================================================================
*/

// --- 1. 文字列 から 数値 へ (Parsing) ---

// 使い方: val, err := strconv.Atoi("123"); if err != nil { ... } と書き，文字列を int に変換する．最も多用される関数である．
func MyAtoi(s string) (int, error) { return strconv.Atoi(s) }

// 使い方: val, err := strconv.ParseInt("-123", 10, 64) と書き，基数（10進数等）とビット長（64）を指定して厳密にパースする．
func MyParseInt(s string, base int, bitSize int) (i int64, err error) {
	return strconv.ParseInt(s, base, bitSize)
}

// 使い方: val, err := strconv.ParseUint("123", 10, 64) と書き，符号なし整数としてパースする．
func MyParseUint(s string, base int, bitSize int) (uint64, error) {
	return strconv.ParseUint(s, base, bitSize)
}

// 使い方: val, err := strconv.ParseFloat("3.14", 64) と書き，文字列を float64 型に変換する．
func MyParseFloat(s string, bitSize int) (float64, error) {
	return strconv.ParseFloat(s, bitSize)
}

// 使い方: val, _ := strconv.ParseBool("true") と書き，"1", "t", "T", "true", "TRUE", "True" 等を真偽値に変換する．
func MyParseBool(str string) (bool, error) { return strconv.ParseBool(str) }

// --- 2. 数値 から 文字列 へ (Formatting) ---

// 使い方: str := strconv.Itoa(123) と書き，int 型を文字列へ変換する．Atoi の逆操作である．
func MyItoa(i int) string { return strconv.Itoa(i) }

// 使い方: str := strconv.FormatInt(int64(123), 10) と書き，int64 を任意の進数（2〜36進数）の文字列へ変換する．
func MyFormatInt(i int64, base int) string { return strconv.FormatInt(i, base) }

// 使い方: str := strconv.FormatUint(uint64(123), 10) と書き，符号なし整数を文字列へ変換する．
func MyFormatUint(i uint64, base int) string { return strconv.FormatUint(i, base) }

// 使い方: str := strconv.FormatFloat(3.14, 'f', -1, 64) と書き，浮動小数点数を文字列に整形する．
// 第2引数は形式（'f':通常, 'e':科学記数法, 'g':最適な表記等），第3引数は精度（-1は最短）を指定する．
func MyFormatFloat(f float64, fmt byte, prec, bitSize int) string {
	return strconv.FormatFloat(f, fmt, prec, bitSize)
}

// 使い方: str := strconv.FormatBool(true) と書き，真偽値を "true" または "false" という文字列に変換する．
func MyFormatBool(b bool) string { return strconv.FormatBool(b) }

// --- 3. 高性能な書き込み (Appending) ---
// メモリ割り当てを減らすため，既存の []byte に直接変換結果を追記する．バッファ処理やロガーの実装に多用される．

// 使い方: dst = strconv.AppendInt(dst, 123, 10) と書き，整数を文字列化したものをスライスの末尾に追加する．
func MyAppendInt(dst []byte, i int64, base int) []byte { return strconv.AppendInt(dst, i, base) }

// 使い方: dst = strconv.AppendFloat(dst, 3.14, 'f', -1, 64) と書き，浮動小数点数をスライスに追加する．
func MyAppendFloat(dst []byte, f float64, fmt byte, prec, bitSize int) []byte {
	return strconv.AppendFloat(dst, f, fmt, prec, bitSize)
}

// --- 4. クォート操作 (Quoting) ---

// 使い方: str := strconv.Quote(`"hello"`) と書き，安全な二重引用符付きの Go 文字列リテラル（例: `"\"hello\""`）を生成する．
func MyQuote(s string) string { return strconv.Quote(s) }

// 使い方: unquoted, _ := strconv.Unquote(`"hello"`) と書き，引用符を取り除き，エスケープ（\n等）を解釈して元の文字列に戻す．
func MyUnquote(s string) (string, error) { return strconv.Unquote(s) }

// 使い方: if strconv.CanBackquote(s) { ... } と書き，文字列が改行や制御文字を含まず，バッククォート（`）で囲めるか判定する．
func MyCanBackquote(s string) bool { return strconv.CanBackquote(s) }

// --- 5. エラー型 ---

// 使い方: if nErr, ok := err.(*strconv.NumError); ok { fmt.Println(nErr.Func, nErr.Num) } と書き，どの関数でどの値が失敗したかを取得する．
type MyNumError struct {
	Func string // 失敗した関数名 (Atoi, ParseInt 等)
	Num  string // 失敗した入力値
	Err  error  // 具体的なエラー (ErrRange, ErrSyntax 等)
}

var (
	// 使い方: if errors.Is(err, strconv.ErrRange) { ... } と書き，オーバーフローを検知する．
	MyErrRange = strconv.ErrRange
	// 使い方: if errors.Is(err, strconv.ErrSyntax) { ... } と書き，不適切な書式の文字列（"abc" 等）を検知する．
	MyErrSyntax = strconv.ErrSyntax
)

func main() {
	println("Loaded: strconv spec.")
}
