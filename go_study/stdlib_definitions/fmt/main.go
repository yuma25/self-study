package main

import (
	"fmt"
	"io"
)

/*
   =============================================================================
   Go 標準ライブラリ 定義集: fmt パッケージ (究極網羅・実戦コード例版)
   ※ 型衝突を避けるため，デモ用の型名には My 接頭辞を付けている．
   ※ 解説は「である調」とし，句読点は「．」「，」を使用している．
   ※ 構造体のフィールド以外は，コードの上部に使い方を記載している．
   =============================================================================
*/

// --- 1. インターフェース (規格) ---

// 使い方: 自作構造体に String() メソッドを定義し，fmt.Println(obj) で独自の文字列表現を出力させる．
type MyStringer interface {
	String() string
}

// 使い方: 独自のエラー型に Error() メソッドを定義し，fmt.Println(err) でエラー内容を表示させる．
type MyError interface {
	Error() string
}

// 使い方: Goの内部表現（%v）をカスタマイズしたい場合に GoString() を実装する．
type MyGoStringer interface {
	GoString() string
}

// 使い方: より高度なフォーマット（%s, %d などの挙動）を独自に制御したい場合に Format を実装する．
type MyFormatter interface {
	Format(f fmt.State, verb rune)
}

// 使い方: Formatter インターフェースの引数として渡され，書式指定子（フラグや幅）の情報を取得するために用いる．
type MyState interface {
	Write(b []byte) (n int, err error)
	Width() (wid int, ok bool)
	Precision() (prec int, ok bool)
	Flag(c int) bool
}

// 使い方: 自作構造体に Scan() メソッドを定義し，fmt.Scan(&obj) 等でカスタムの読み取り規則を適用させる．
type MyScanner interface {
	Scan(state fmt.ScanState, verb rune) error
}

// 使い方: Scanner インターフェースの引数として渡され，入力ストリームからトークンや文字を読み取るために用いる．
type MyScanState interface {
	ReadRune() (r rune, size int, err error)
	UnreadRune() error
	SkipSpace()
	Token(skipSpace bool, f func(rune) bool) (token []byte, err error)
	Width() (wid int, ok bool)
	Read(b []byte) (n int, err error)
}

// --- 2. パッケージレベルの関数 (操作) ---

// --- 出力 (Print系) ---

// 使い方: fmt.Print("Hello", " ", "World") と書き，改行なしで連結して出力する．
func MyPrint(a ...any) (n int, err error) { return 0, nil }

// 使い方: fmt.Println("Result:", val) と書き，要素間にスペースを挟み，最後に改行を入れて出力する．
func MyPrintln(a ...any) (n int, err error) { return 0, nil }

// 使い方: fmt.Printf("ID: %d, Name: %s\n", id, name) と書き，書式指定子を用いて整形出力する．
func MyPrintf(format string, a ...any) (n int, err error) { return 0, nil }

// --- 文字列生成 (Sprint系) ---

// 使い方: str := fmt.Sprint("A", "B") と書き，出力を画面ではなく文字列として受け取る．
func MySprint(a ...any) string { return "" }

// 使い方: str := fmt.Sprintln("A", "B") と書き，スペースと改行を含む文字列を生成する．
func MySprintln(a ...any) string { return "" }

// 使い方: msg := fmt.Sprintf("Error: %v", err) と書き，整形された文字列を生成して変数に格納する．
func MySprintf(format string, a ...any) string { return "" }

// --- io.Writer への出力 (Fprint系) ---

// 使い方: fmt.Fprint(os.Stderr, "Error occurred") と書き，特定の Writer へ改行なしで書き込む．
func MyFprint(w io.Writer, a ...any) (n int, err error) { return 0, nil }

// 使い方: fmt.Fprintln(file, "Log entry") と書き，ログファイル等へ改行付きで書き込む．
func MyFprintln(w io.Writer, a ...any) (n int, err error) { return 0, nil }

// 使い方: fmt.Fprintf(w, "HTTP/1.1 %d OK\n", 200) と書き，HTTPレスポンス等へ整形して書き込む．
func MyFprintf(w io.Writer, format string, a ...any) (n int, err error) { return 0, nil }

// --- 入力 (Scan系: 標準入力から) ---

// 使い方: var name string; fmt.Scan(&name) と書き，標準入力（キーボード等）から空白区切りで変数に値を読み込む．
func MyScan(a ...any) (n int, err error) { return 0, nil }

// 使い方: var id int; fmt.Scanf("%d", &id) と書き，フォーマットを指定して標準入力から値を読み込む．
func MyScanf(format string, a ...any) (n int, err error) { return 0, nil }

// 使い方: fmt.Scanln(&val1, &val2) と書き，改行で読み込みを終了する（1行単位で読み込む）．
func MyScanln(a ...any) (n int, err error) { return 0, nil }

// --- io.Reader からの入力 (Fscan系) ---

// 使い方: fmt.Fscan(file, &val) と書き，ファイルなどの io.Reader から空白区切りで値を読み込む．
func MyFscan(r io.Reader, a ...any) (n int, err error) { return 0, nil }

// 使い方: fmt.Fscanf(file, "ID:%d", &id) と書き，io.Reader からフォーマットに従って値を読み込む．
func MyFscanf(r io.Reader, format string, a ...any) (n int, err error) { return 0, nil }

// 使い方: fmt.Fscanln(file, &val) と書き，io.Reader から改行までを1単位として読み込む．
func MyFscanln(r io.Reader, a ...any) (n int, err error) { return 0, nil }

// --- 文字列からの入力 (Sscan系) ---

// 使い方: fmt.Sscan("123 apple", &id, &name) と書き，既存の文字列から変数へ値を抽出する．
func MySscan(str string, a ...any) (n int, err error) { return 0, nil }

// 使い方: fmt.Sscanf("ID=123", "ID=%d", &id) と書き，文字列から特定のフォーマットに従って値を抽出する．
func MySscanf(str string, format string, a ...any) (n int, err error) { return 0, nil }

// 使い方: fmt.Sscanln("123\n", &val) と書き，文字列から改行までを読み込む．
func MySscanln(str string, a ...any) (n int, err error) { return 0, nil }

// --- エラー生成 (Errorf) ---

// 使い方: err := fmt.Errorf("failed to open file %s: %w", name, originalErr) と書き，エラーをラップ（包む）して返す．
func MyErrorf(format string, a ...any) error { return nil }

// --- 3. 重要な書式指定子 (Verbs) の解説メモ ---
/*
  実務で頻出する Printf / Sprintf 用の書式指定子である．
  - %v  : デフォルトのフォーマット（どんな型でもとりあえず出力できる）．
  - %+v : 構造体のフィールド名付きで出力する（デバッグに超有用）．
  - %#v : Goの構文表現で出力する（型名もわかる）．
  - %T  : 変数の型名を出力する．
  - %d  : 整数を10進数で出力する．
  - %s  : 文字列または []byte を出力する．
  - %f  : 浮動小数点数を出力する（例: %.2f で小数第2位まで）．
  - %t  : 真偽値（true/false）を出力する．
  - %w  : Errorf 専用．発生したエラーを包み込んで（ラップして）新しいエラーを作る．
*/

func main() {
	println("Loaded: fmt spec including Scan family and Formatter interfaces.")
}
