package main

import (
	"strings"
)

/*
   =============================================================================
   Go 標準ライブラリ 定義集: strings パッケージ (究極網羅・実戦コード例版)
   =============================================================================
*/

// --- 1. 構造体 (実体) とそのメソッド ---

// 使い方: builder := &strings.Builder{}; builder.WriteString("a"); str := builder.String() と書き，効率的に文字列を結合する．
// 大量の文字列を + 演算子で繋ぐとメモリ効率が悪いため，実務では必ず Builder を使用する．
type MyBuilder struct {
	strings.Builder
}

func (b *MyBuilder) MyWrite(p []byte) (int, error) { return b.Write(p) }
func (b *MyBuilder) MyWriteByte(c byte) error      { return b.WriteByte(c) }
func (b *MyBuilder) MyWriteRune(r rune) (int, error) { return b.WriteRune(r) }
func (b *MyBuilder) MyWriteString(s string) (int, error) { return b.WriteString(s) }
func (b *MyBuilder) MyReset()                      { b.Reset() }
func (b *MyBuilder) MyLen() int                    { return b.Len() }
func (b *MyBuilder) MyCap() int                    { return b.Cap() }
func (b *MyBuilder) MyGrow(n int)                  { b.Grow(n) }

// 使い方: reader := strings.NewReader("data"); io.Copy(os.Stdout, reader) と書き，文字列を io.Reader として扱う．
// メモリ上の文字列を，ファイルやネットワーク通信と同じようにストリーム処理したい場合に用いる．
type MyReader struct {
	*strings.Reader
}

func (r *MyReader) MyRead(b []byte) (n int, err error) { return r.Read(b) }
func (r *MyReader) MyReadAt(b []byte, off int64) (n int, err error) { return r.ReadAt(b, off) }
func (r *MyReader) MySeek(offset int64, whence int) (int64, error) { return r.Seek(offset, whence) }
func (r *MyReader) MyLen() int                         { return r.Len() }
func (r *MyReader) MySize() int64                      { return r.Size() }

// --- 2. パッケージレベルの関数 (文字列操作) ---

// --- 検索・判定系 ---

// 使い方: if strings.Contains(url, "https") { ... } と書き，部分文字列が含まれるか判定する．
func MyContains(s, substr string) bool { return strings.Contains(s, substr) }

// 使い方: if strings.EqualFold("Go", "go") { ... } と書き，大文字小文字を無視して一致するか判定する．
func MyEqualFold(s, t string) bool { return strings.EqualFold(s, t) }

// 使い方: if strings.HasPrefix(filename, "test_") { ... } と書き，先頭の文字列が一致するか判定する．
func MyHasPrefix(s, prefix string) bool { return strings.HasPrefix(s, prefix) }

// 使い方: if strings.HasSuffix(filename, ".go") { ... } と書き，末尾の文字列が一致するか（拡張子判定等）を確認する．
func MyHasSuffix(s, suffix string) bool { return strings.HasSuffix(s, suffix) }

// 使い方: idx := strings.Index(path, "/") と書き，部分文字列が最初に現れる位置を取得する．無ければ -1 を返す．
func MyIndex(s, substr string) int { return strings.Index(s, substr) }

// 使い方: idx := strings.LastIndex(path, "/") と書き，最後に出現する位置を取得する．
func MyLastIndex(s, substr string) int { return strings.LastIndex(s, substr) }

// 使い方: count := strings.Count(text, "error") と書き，特定の文字列が出現する回数を数える．
func MyCount(s, substr string) int { return strings.Count(s, substr) }

// --- 変換・置換系 ---

// 使い方: upper := strings.ToUpper("hello") と書き，すべて大文字に変換する．
func MyToUpper(s string) string { return strings.ToUpper(s) }

// 使い方: lower := strings.ToLower("HELLO") と書き，すべて小文字に変換する．
func MyToLower(s string) string { return strings.ToLower(s) }

// 使い方: clean := strings.TrimSpace("  data  \n") と書き，前後の空白や改行を安全に除去する．
func MyTrimSpace(s string) string { return strings.TrimSpace(s) }

// 使い方: clean := strings.Trim(s, "[]") と書き，指定した文字セットを前後から除去する．
func MyTrim(s, cutset string) string { return strings.Trim(s, cutset) }

// 使い方: result := strings.ReplaceAll(text, "old", "new") と書き，すべて一括で置換する．
func MyReplaceAll(s, old, new string) string { return strings.ReplaceAll(s, old, new) }

// 使い方: r := strings.NewReplacer("a", "1", "b", "2"); res := r.Replace("abc") と書き，複数の置換ルールを一気に適用する．
func MyNewReplacer(oldnew ...string) *strings.Replacer { return strings.NewReplacer(oldnew...) }

// 使い方: s := strings.Map(func(r rune) rune { ... }, str) と書き，一文字ずつ関数を通して変換（または削除）する．
func MyMap(mapping func(rune) rune, s string) string { return strings.Map(mapping, s) }

// --- 分割・結合系 ---

// 使い方: parts := strings.Split("a,b,c", ",") と書き，区切り文字で文字列を分割する．
func MySplit(s, sep string) []string { return strings.Split(s, sep) }

// 使い方: words := strings.Fields("  foo bar  baz ") と書き，連続する空白文字を無視して単語に分割する．
func MyFields(s string) []string { return strings.Fields(s) }

// 使い方: str := strings.Join([]string{"a", "b"}, "-") と書き，スライスを指定文字で繋いで1つの文字列にする．
func MyJoin(elems []string, sep string) string { return strings.Join(elems, sep) }

// 使い方: before, after, found := strings.Cut("name=alice", "=") と書き，最初の境界で前後を分割する．
func MyCut(s, sep string) (before, after string, found bool) { return strings.Cut(s, sep) }

// 使い方: repeated := strings.Repeat("A", 10) と書き，指定回数繰り返した文字列を作成する．
func MyRepeat(s string, count int) string { return strings.Repeat(s, count) }

// --- モダンな機能 (Go 1.20+) ---

// 使い方: copied := strings.Clone(s) と書き，巨大な文字列の一部を切り取った際に，元の巨大なメモリを保持しないよう新しいコピーを作成する．
func MyClone(s string) string { return strings.Clone(s) }

func main() {
	println("Loaded: strings spec.")
}
