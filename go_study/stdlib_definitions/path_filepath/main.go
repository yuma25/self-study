package main

import (
	"io/fs"
	"path"
	"path/filepath"
)

/*
   =============================================================================
   Go 標準ライブラリ 定義集: path/filepath パッケージ (究極網羅・実戦コード例版)
   =============================================================================
*/

// --- 1. path/filepath (OS依存のパス操作) ---

// 使い方: fullPath := filepath.Join("dir", "subdir", "file.txt") と書き，OSごとの正しい区切り文字（/ や \）を使って安全にパスを結合する．
func MyJoin(elem ...string) string { return filepath.Join(elem...) }

// 使い方: ext := filepath.Ext("image.png") と書き，拡張子（例: ".png"）を取得する．無ければ空文字を返す．
func MyExt(p string) string { return filepath.Ext(p) }

// 使い方: base := filepath.Base("/dir/file.txt") と書き，最後の要素（ファイル名 "file.txt"）だけを抽出する．
func MyBase(p string) string { return filepath.Base(p) }

// 使い方: dir := filepath.Dir("/dir/file.txt") と書き，最後の要素を除いた親ディレクトリのパス（"/dir"）を抽出する．
func MyDir(p string) string { return filepath.Dir(p) }

// 使い方: cleanPath := filepath.Clean("/dir/../dir/./file.txt") と書き，冗長なドット（.. や .）を解決して最短のきれいなパス（"/dir/file.txt"）に正規化する．
func MyClean(p string) string { return filepath.Clean(p) }

// 使い方: absPath, _ := filepath.Abs("file.txt") と書き，現在の作業ディレクトリを基準とした絶対パス（"/home/user/file.txt"）に変換する．
func MyAbs(p string) (string, error) { return filepath.Abs(p) }

// 使い方: relPath, _ := filepath.Rel("/home/user", "/home/user/dir/file.txt") と書き，ある基準パスから見た相対パス（"dir/file.txt"）を計算する．
func MyRel(basepath, targpath string) (string, error) { return filepath.Rel(basepath, targpath) }

// 使い方: dir, file := filepath.Split("/a/b/c.txt") と書き，ディレクトリパスとファイル名に分割する．
func MySplit(p string) (dir, file string) { return filepath.Split(p) }

// 使い方: paths := filepath.SplitList("a:b:c") と書き，PATH環境変数のような区切り文字（Linuxは: , Windowsは;）で分割する．
func MySplitList(p string) []string { return filepath.SplitList(p) }

// 使い方: p := filepath.FromSlash("a/b/c") と書き，スラッシュをOS固有の区切り文字に置換する．
func MyFromSlash(p string) string { return filepath.FromSlash(p) }

// 使い方: p := filepath.ToSlash("a\\b\\c") と書き，OS固有の区切り文字をスラッシュに置換する．URL等に使う場合に便利．
func MyToSlash(p string) string { return filepath.ToSlash(p) }

// 使い方: match, _ := filepath.Match("*.txt", "test.txt") と書き，ファイル名がパターンに合致するか判定する．
func MyMatch(pattern, name string) (matched bool, err error) { return filepath.Match(pattern, name) }

// 使い方: if filepath.IsAbs(path) { ... } と書き，そのパスが絶対パスとして記述されているかを判定する．
func MyIsAbs(p string) bool { return filepath.IsAbs(p) }

// 使い方: filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error { ... }) と書き，指定ディレクトリ以下のすべてのファイルとフォルダを再帰的に走査（検索）する．
func MyWalkDir(root string, fn fs.WalkDirFunc) error { return filepath.WalkDir(root, fn) }

// 使い方: matches, _ := filepath.Glob("*.txt") と書き，ワイルドカード（* や ?）に一致するファイル名の一覧を取得する．
func MyGlob(pattern string) (matches []string, err error) { return filepath.Glob(pattern) }

// --- 重要な定数 ---

const (
	// 使い方: fmt.Println(string(filepath.Separator)) と書き，実行環境のパス区切り文字（Linux/Mac なら '/', Windows なら '\'）を取得する．
	MySeparator = filepath.Separator

	// 使い方: fmt.Println(string(filepath.ListSeparator)) と書き，PATH環境変数の区切り文字（Linuxは ':' , Windowsは ';'）を取得する．
	MyListSeparator = filepath.ListSeparator
)

// --- 2. path (スラッシュ固定・URL/ZIP用) ---
// path/filepath と異なり，OSに関わらず常に '/' を区切り文字として扱う．
// クラウドストレージのキー，URLのパス，ZIP内のパス操作にはこちらを使うのが鉄則．

// 使い方: p := path.Join("api", "v1", "users") と書き，OSに関係なく "api/v1/users" を作成する．
func MyPathJoin(elem ...string) string { return path.Join(elem...) }

func main() {
	println("Loaded: path/filepath spec.")
}
