package main

import (
	"regexp"
)

/*
   =============================================================================
   Go 標準ライブラリ 定義集: regexp パッケージ (究極網羅・実戦コード例版)
   =============================================================================
*/

// --- 1. 構造体 (実体) とそのメソッド ---

// 使い方: re := regexp.MustCompile(`^[a-z]+$`) と書き，コンパイル済みの正規表現パターンを保持する．
// 実行中に Compile すると遅いため，通常はグローバル変数として初期化時に MustCompile で作っておく．
type MyRegexp struct {
	*regexp.Regexp
}

// 使い方: if re.MatchString("abc") { ... } と書き，文字列がパターンに完全に一致するか（または含まれるか）を判定する．
func (re *MyRegexp) MyMatchString(s string) bool { return re.MatchString(s) }

// 使い方: str := re.FindString("foo 123 bar") と書き，パターンに一致する最初の部分文字列（例では "123"）を抽出する．
func (re *MyRegexp) MyFindString(s string) string { return re.FindString(s) }

// 使い方: strs := re.FindAllString("a1 b2 c3", -1) と書き，パターンに一致するすべての部分文字列を抽出する．-1 は制限なしを意味する．
func (re *MyRegexp) MyFindAllString(s string, n int) []string { return re.FindAllString(s, n) }

// 使い方: parts := re.FindStringSubmatch("name: alice") と書き，() で囲んだキャプチャグループの抽出結果を配列で得る．parts[1] に "alice" が入る．
func (re *MyRegexp) MyFindStringSubmatch(s string) []string { return re.FindStringSubmatch(s) }

// 使い方: indices := re.FindStringSubmatchIndex("name: alice") と書き，キャプチャグループごとの開始・終了位置をペアで得る．
func (re *MyRegexp) MyFindStringSubmatchIndex(s string) []int { return re.FindStringSubmatchIndex(s) }

// 使い方: result := re.ReplaceAllString("a b c", "-") と書き，パターンに一致する部分をすべて別の文字列に置換する．
func (re *MyRegexp) MyReplaceAllString(src, repl string) string { return re.ReplaceAllString(src, repl) }

// 使い方: result := re.ReplaceAllStringFunc("a1 b2", func(s string) string { ... }) と書き，一致した各部分に対して関数で動的に置換内容を決定する．
func (re *MyRegexp) MyReplaceAllStringFunc(src string, repl func(string) string) string {
	return re.ReplaceAllStringFunc(src, repl)
}

// 使い方: parts := re.Split("a, b, c", -1) と書き，正規表現を区切り文字として文字列を分割する．
func (re *MyRegexp) MySplit(s string, n int) []string { return re.Split(s, n) }

// 使い方: names := re.SubexpNames() と書き，(?P<name>...) で定義した名前付きグループの名称一覧を取得する．
func (re *MyRegexp) MySubexpNames() []string { return re.SubexpNames() }

// 使い方: re.Longest() と書き，常に最も長い一致（最長一致）を優先するように設定する（デフォルトは左端一致）．
func (re *MyRegexp) MyLongest() { re.Longest() }

// --- 2. パッケージレベルの関数 (操作) ---

// 使い方: re, err := regexp.Compile(`\d+`) と書き，実行時に動的に正規表現をコンパイルする．文法エラーがあれば err が返る．
func MyCompile(expr string) (*MyRegexp, error) {
	re, err := regexp.Compile(expr)
	if err != nil {
		return nil, err
	}
	return &MyRegexp{re}, nil
}

// 使い方: re := regexp.MustCompile(`\d+`) と書き，初期化時にコンパイルする．文法エラーがあれば panic を起こして即座にプログラムを止める．
func MyMustCompile(str string) *MyRegexp {
	return &MyRegexp{regexp.MustCompile(str)}
}

// 使い方: if matched, _ := regexp.MatchString(`\d+`, "abc123def"); matched { ... } と書き，一度きりの使い捨ての判定を行う．内部で都度コンパイルされるためループ内では非推奨．
func MyMatchString(pattern string, s string) (matched bool, err error) {
	return regexp.MatchString(pattern, s)
}

// 使い方: safeStr := regexp.QuoteMeta(`[foo]`) と書き，正規表現の特殊文字（[, ], *, ? 等）をエスケープして，単なる文字列として検索できるようにする．
func MyQuoteMeta(s string) string { return regexp.QuoteMeta(s) }

func main() {
	println("Loaded: regexp spec.")
}
