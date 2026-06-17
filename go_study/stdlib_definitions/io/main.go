package main

import (
	"io"
)

/*
   =============================================================================
   Go 標準ライブラリ 定義集: io パッケージ (実戦コード例配置調整版)
   ※ 型衝突を避けるため，デモ用の型名には My 接頭辞を付けている．
   ※ 解説は「である調」とし，句読点は「．」「，」を使用している．
   ※ 構造体のフィールド以外は，コードの上部に使い方を記載している．
   =============================================================================
*/

// --- 1. インターフェース (規格) ---

// 使い方: data := make([]byte, 1024); n, err := r.Read(data) と書き，バッファにデータを吸い出す．
type MyReader interface {
	Read(p []byte) (n int, err error)
}

// 使い方: w.Write([]byte("Hello")) と書き，ストリームへデータを流し込む．
type MyWriter interface {
	Write(p []byte) (n int, err error)
}

// 使い方: defer f.Close() と書き，関数の最後で必ずリソースを解放する．
type MyCloser interface {
	Close() error
}

// 使い方: f.Seek(0, io.SeekStart) と書き，ファイルの先頭に読み書き位置を巻き戻す．
type MySeeker interface {
	Seek(offset int64, whence int) (int64, error)
}

// 使い方: 構造体に実装し，io.Copy 等から呼ばれた際に直接バッファを介さずゼロコピー転送を行う．
type MyReaderFrom interface {
	ReadFrom(r MyReader) (n int64, err error)
}

// 使い方: 構造体に実装し，io.Copy 等から呼ばれた際に自身のデータを相手に直接書き込む．
type MyWriterTo interface {
	WriteTo(w MyWriter) (n int64, err error)
}

// 使い方: r.ReadAt(buf, 100) と書き，現在位置を変えずにファイルの100バイト目から並行して読み込む．
type MyReaderAt interface {
	ReadAt(p []byte, off int64) (n int, err error)
}

// 使い方: w.WriteAt(buf, 100) と書き，現在位置を変えずにファイルの100バイト目へ並行して書き込む．
type MyWriterAt interface {
	WriteAt(p []byte, off int64) (n int, err error)
}

// 使い方: b, err := r.ReadByte() と書き，バッファなしで1バイトだけ細かく読み取る．
type MyByteReader interface {
	ReadByte() (byte, error)
}

// 使い方: r.UnreadByte() と書き，構文解析中に「読みすぎた1バイト」をストリームへ戻す．
type MyByteScanner interface {
	MyByteReader
	UnreadByte() error
}

// --- 2. 複合インターフェース ---

// 使い方: HTTPリクエストのボディ等，読み取った後に閉じる必要があるものを受け取る型として指定する．
type MyReadCloser interface {
	MyReader
	MyCloser
}

// 使い方: ネットワークソケット等，双方向通信かつ解放が必要なものを受け取る型として指定する．
type MyReadWriteCloser interface {
	MyReader
	MyWriter
	MyCloser
}

// --- 3. 構造体 (実体としてのヘルパー) ---

// 使い方: io.LimitReader(r, 1024) で作成し，巨大な入力から最大1KBだけを安全に読み取る．
type MyLimitedReader struct {
	R MyReader // 使い方: 制限をかけたい元のリーダーを指定する．
	N int64    // 使い方: 残りの読み取り可能バイト数である．Read毎に減少し，0でEOFになる．
}

func (l *MyLimitedReader) Read(p []byte) (n int, err error) { return 0, nil }

// 使い方: io.NewSectionReader(r, 100, 50) で作成し，ファイルの一部（100バイト目からの50バイト）を独立したファイルのように扱う．
type MySectionReader struct{}

func (s *MySectionReader) Read(p []byte) (n int, err error) { return 0, nil }
func (s *MySectionReader) Seek(offset int64, whence int) (int64, error) { return 0, nil }
func (s *MySectionReader) ReadAt(p []byte, off int64) (n int, err error) { return 0, nil }

// 使い方: io.Pipe() で作成し，ゴルーチン間で「書いた端から読む」という同期的なストリーム通信を行う．
type MyPipeReader struct{}

func (r *MyPipeReader) Read(data []byte) (n int, err error) { return 0, nil }
func (r *MyPipeReader) Close() error { return nil }

type MyPipeWriter struct{}

func (w *MyPipeWriter) Write(data []byte) (n int, err error) { return 0, nil }
func (w *MyPipeWriter) Close() error { return nil }

// --- 4. パッケージレベルの関数 (操作) ---

// 使い方: io.Copy(os.Stdout, file) と書き，ファイルを画面へメモリを節約しつつ垂れ流す．
func MyCopy(dst MyWriter, src MyReader) (written int64, err error) { return 0, nil }

// 使い方: io.CopyN(w, r, 256) と書き，ストリームから正確に256バイトだけを切り出して転送する．
func MyCopyN(dst MyWriter, src MyReader, n int64) (written int64, err error) { return 0, nil }

// 使い方: data, err := io.ReadAll(r) と書き，ストリームの全データを一括でメモリ（スライス）に読み込む．
func MyReadAll(r MyReader) ([]byte, error) { return nil, nil }

// 使い方: io.ReadFull(r, buf) と書き，用意したバッファ buf が完全に埋まるまで意地でも読み込み続ける．
func MyReadFull(r MyReader, buf []byte) (n int, err error) { return 0, nil }

// 使い方: io.MultiReader(file1, file2) と書き，複数のファイルを連結して1つの連続したファイルのように読む．
func MyMultiReader(readers ...MyReader) MyReader { return nil }

// 使い方: io.MultiWriter(os.Stdout, logFile) と書き，画面とファイルの両方に同時に書き込む魔法のペンを作る．
func MyMultiWriter(writers ...MyWriter) MyWriter { return nil }

// 使い方: io.TeeReader(r, w) と書き，読み込んだ内容をそのまま分岐させて別の Writer（ログ等）へも流し込む．
func MyTeeReader(r MyReader, w MyWriter) MyReader { return nil }

// 使い方: io.WriteString(w, "Hello") と書き，[]byte キャストのオーバーヘッドを避けて文字列を直接書き込む．
func MyWriteString(w MyWriter, s string) (n int, err error) { return 0, nil }

// --- 5. 重要な変数・定数 (標準ライブラリの値を参照) ---

var (
	// 使い方: if err == io.EOF { break } と書き，ファイルの終端（読み込み完了）を検知してループを抜ける．
	MyEOF = io.EOF
	// 使い方: if err == io.ErrUnexpectedEOF { ... } と書き，指定バイト数に達する前に通信が切れた等の異常終了を検知する．
	MyErrUnexpectedEOF = io.ErrUnexpectedEOF
	// 使い方: io.Copy(io.Discard, r) と書き，ストリームのデータをどこにも保存せず単に消費（読み捨てる）する．
	MyDiscard = io.Discard
)

func main() {
	println("Updated: io spec with explicit 'How to use' code examples loaded.")
}
