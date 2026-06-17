package main

import (
	"os"
	"time"
)

/*
   =============================================================================
   Go 標準ライブラリ 定義集: os パッケージ (究極網羅・実戦コード例版)
   ※ 型衝突を避けるため，デモ用の型名には My 接頭辞を付けている．
   ※ 解説は「である調」とし，句読点は「．」「，」を使用している．
   ※ 構造体のフィールド以外は，コードの上部に使い方を記載している．
   =============================================================================
*/

// --- 1. インターフェース (規格) ---

// 使い方: info, err := os.Stat("file.txt"); fmt.Println(info.Size()) と書き，ファイルの詳細情報を取得する．
type MyFileInfo interface {
	Name() string       // 使い方: fmt.Println("Name:", info.Name()) と書き，"file.txt" のようなベース名を得る．
	Size() int64        // 使い方: if info.Size() > 1024 { ... } と書き，ファイルサイズで処理を分岐する．
	Mode() os.FileMode  // 使い方: if info.Mode().Perm() == 0644 { ... } と書き，権限を確認する．
	ModTime() time.Time // 使い方: if info.ModTime().Before(threshold) { ... } と書き，古いファイルを判定する．
	IsDir() bool        // 使い方: if info.IsDir() { return } と書き，ディレクトリに対する誤操作を防ぐ．
	Sys() any           // 使い方: stat := info.Sys().(*syscall.Stat_t) と書き，OS依存の深い情報を得る．
}

// 使い方: entries, _ := os.ReadDir("."); for _, e := range entries { ... } と書き，フォルダ内を走査する．
type MyDirEntry interface {
	Name() string               // 使い方: if strings.HasSuffix(e.Name(), ".go") { ... } と書き，拡張子で絞り込む．
	IsDir() bool                // 使い方: if e.IsDir() { continue } と書き，サブディレクトリをスキップする．
	Type() os.FileMode          // 使い方: if e.Type() & os.ModeSymlink != 0 { ... } と書き，シンボリックリンクを判定する．
	Info() (os.FileInfo, error) // 使い方: info, _ := e.Info(); size := info.Size() と書き，詳細が必要な時だけ重いStat処理を呼ぶ．
}

// 使い方: sig := <-make(chan os.Signal, 1); if sig == os.Interrupt { ... } と書き，Ctrl+C等の割り込みを捉える．
type MySignal interface {
	String() string // 使い方: fmt.Println("Received:", sig.String()) と書き，シグナルの名前（"interrupt"等）をログに出す．
	Signal()        // 使い方: インターフェースの型を区別するためのダミーであるため，直接呼び出すことはない．
}

// --- 2. 構造体 (実体) とそのメソッド ---

// 使い方: f, _ := os.Open("data.txt"); defer f.Close() と書き，ファイル操作の基本実体として扱う．
type MyFile struct{}

// 使い方: buf := make([]byte, 100); n, _ := f.Read(buf) と書き，ファイルからデータを吸い上げる．
func (f *MyFile) Read(b []byte) (n int, err error) { return 0, nil }

// 使い方: f.Write([]byte("Hello")) と書き，ファイルへデータを書き込む．
func (f *MyFile) Write(b []byte) (n int, err error) { return 0, nil }

// 使い方: defer f.Close() と書き，ファイルディスクリプタの枯渇（OSのエラー）を防ぐ．
func (f *MyFile) Close() error { return nil }

// 使い方: info, _ := f.Stat() と書き，開いているファイル自身のメタデータを得る．
func (f *MyFile) Stat() (os.FileInfo, error) { return nil, nil }

// 使い方: f.Sync() と書き，OSのバッファに溜まったデータをディスクへ物理的に強制書き込み(fsync)させる．
func (f *MyFile) Sync() error { return nil }

// 使い方: f.Truncate(0) と書き，開いているファイルの中身を空（サイズ0）に切り詰める．
func (f *MyFile) Truncate(size int64) error { return nil }

// 使い方: proc, _ := os.FindProcess(1234) と書き，操作対象のOSプロセス実体を取得する．
type MyProcess struct {
	Pid int // 使い方: fmt.Println("PID:", proc.Pid) と書き，プロセスIDをログに残す．
}

// 使い方: err := proc.Kill() と書き，暴走したプロセスを強制的に終了させる．
func (p *MyProcess) Kill() error { return nil }

// 使い方: state, _ := proc.Wait() と書き，子プロセスの終了をブロックして待機する．
func (p *MyProcess) Wait() (*os.ProcessState, error) { return nil, nil }

// 使い方: state.Exited() と書き，プロセスが本当に終わったか確認する．
type MyProcessState struct{}

// 使い方: if state.ExitCode() != 0 { ... } と書き，プログラムが異常終了したか判定する．
func (p *MyProcessState) ExitCode() int { return 0 }

// --- 3. パッケージレベルの関数 (操作) ---

// 使い方: f, err := os.Open("config.json") と書き，読み取り専用で安全にファイルを開く．
func MyOpen(name string) (*MyFile, error) { return nil, nil }

// 使い方: f, err := os.Create("output.log") と書き，空の新しいファイルを作る（あれば上書きされる）．
func MyCreate(name string) (*MyFile, error) { return nil, nil }

// 使い方: f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) と書き，追記モード等の詳細指定で開く．
func MyOpenFile(name string, flag int, perm os.FileMode) (*MyFile, error) { return nil, nil }

// 使い方: data, err := os.ReadFile("key.txt") と書き，数MB程度のファイルを一気にスライスへ読み込む．
func MyReadFile(name string) ([]byte, error) { return nil, nil }

// 使い方: err := os.WriteFile("data.json", jsonBytes, 0644) と書き，データ一式を安全に保存する．
func MyWriteFile(name string, data []byte, perm os.FileMode) error { return nil }

// 使い方: err := os.Mkdir("new_dir", 0755) と書き，新しいフォルダを1つ作る．
func MyMkdir(name string, perm os.FileMode) error { return nil }

// 使い方: err := os.MkdirAll("path/to/deep/dir", 0755) と書き，途中のフォルダも含めて階層ごと一気に作る(mkdir -p)．
func MyMkdirAll(path string, perm os.FileMode) error { return nil }

// 使い方: err := os.Remove("temp.txt") と書き，不要になったファイルや空のフォルダを消す．
func MyRemove(name string) error { return nil }

// 使い方: err := os.RemoveAll("temp_dir") と書き，フォルダの中身ごと強制的にすべて消去する(rm -rf)．
func MyRemoveAll(path string) error { return nil }

// 使い方: err := os.Rename("old.txt", "new.txt") と書き，ファイル名の変更や，別フォルダへの移動を行う．
func MyRename(oldpath, newpath string) error { return nil }

// 使い方: port := os.Getenv("PORT"); if port == "" { port = "8080" } と書き，環境変数から設定を読む．
func MyGetenv(key string) string { return "" }

// 使い方: os.Setenv("DEBUG", "1") と書き，実行中のプログラム内だけで有効な環境変数をセットする．
func MySetenv(key, value string) error { return nil }

// 使い方: dir, _ := os.Getwd() と書き，現在プログラムが実行されているカレントディレクトリのパスを得る．
func MyGetwd() (dir string, err error) { return "", nil }

// 使い方: os.Exit(1) と書き，致命的なエラー発生時にプログラムを異常終了（ステータスコード1）させる．
func MyExit(code int) {}

// --- 4. 重要な標準変数 ---

var (
	// 使い方: fmt.Fprintln(os.Stdout, "Hello") と書き，コンソール画面へ標準出力を行う．
	MyStdout = os.Stdout
	// 使い方: fmt.Fprintln(os.Stderr, "Error!") と書き，エラーメッセージを標準エラー出力へ分離して出す．
	MyStderr = os.Stderr
	// 使い方: io.Copy(os.Stdout, os.Stdin) と書き，キーボードからの標準入力を受け取る．
	MyStdin = os.Stdin
	// 使い方: if len(os.Args) > 1 { target = os.Args[1] } と書き，コマンドライン引数を取得する．
	MyArgs = os.Args
)

func main() {
	println("Updated: Full os spec with exhaustive 'How to use' code examples loaded.")
}
