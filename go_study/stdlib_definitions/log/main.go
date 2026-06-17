package main

import (
	"context"
	"io"
	"log/slog"
)

/*
   =============================================================================
   Go 標準ライブラリ 定義集: log パッケージ群 (究極網羅・実戦コード例版)
   =============================================================================
*/

// --- 1. 従来の log パッケージ (単純出力用) ---

// 使い方: logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime) と書き，独自のプレフィックスやフォーマットを持つロガーを作成する．
type MyLogger struct { /* 非公開 */ }

// 使い方: logger.Println("Server started") と書き，設定されたプレフィックスと日時を自動で付与してログを出力する．
func (l *MyLogger) Println(v ...any) {}

// 使い方: logger.Printf("User %s logged in", name) と書き，書式指定を伴ってログを出力する．
func (l *MyLogger) Printf(format string, v ...any) {}

// 使い方: logger.Fatal("Failed to bind port") と書き，ログを出力した直後に os.Exit(1) を呼んでプログラムを強制終了させる．
func (l *MyLogger) Fatal(v ...any) {}

// 使い方: logger.Fatalf("DB error: %v", err) と書き，書式指定でログを出力した直後に強制終了させる．
func (l *MyLogger) Fatalf(format string, v ...any) {}

// 使い方: logger.Panic("Unexpected state") と書き，ログを出力した直後に panic() を発生させ，スタックトレースを記録する．
func (l *MyLogger) Panic(v ...any) {}

// 使い方: logger.SetOutput(file) と書き，ログの出力先を標準エラー出力から任意の io.Writer（ファイル等）に変更する．
func (l *MyLogger) SetOutput(w io.Writer) {}

// 使い方: logger.SetPrefix("DEBUG: ") と書き，既存のロガーのプレフィックスを動的に変更する．
func (l *MyLogger) SetPrefix(prefix string) {}

// 使い方: logger.SetFlags(log.LstdFlags | log.Lshortfile) と書き，ログに付与するメタ情報（日付，時刻，ファイル名，行番号等）を設定する．
func (l *MyLogger) SetFlags(flag int) {}

// --- 2. パッケージレベルの関数 (操作: デフォルトロガー用) ---

// 使い方: log.Println("Simple log") と書き，組み込みの標準ロガー（出力先: os.Stderr）を使って手軽にログを出す．
func MyPrintln(v ...any) {}

// 使い方: log.Printf("Count: %d", cnt) と書き，標準ロガーで整形出力を行う．
func MyPrintf(format string, v ...any) {}

// 使い方: log.Fatal("Cannot continue") と書き，標準ロガーでエラー出力後にプログラムを即死させる．（初期化失敗時等に使用）
func MyFatal(v ...any) {}

// 使い方: log.SetOutput(file) と書き，標準ロガーの出力先をプログラム全体で一括変更する．
func MySetOutput(w io.Writer) {}

// --- 3. 重要な定数 (フラグ) ---

const (
	// 使い方: log.Ldate | log.Ltime と書き，"2009/01/23 01:23:23" のような標準的な日時をログに付与する．
	MyLdate = 1 << iota
	MyLtime

	// 使い方: log.Lmicroseconds と書き，ミリ秒より細かいマイクロ秒単位の時刻を記録する．
	MyLmicroseconds

	// 使い方: log.Llongfile と書き，"/a/b/c/d.go:23" のような絶対パスと行番号をログに付与してデバッグを容易にする．
	MyLlongfile

	// 使い方: log.Lshortfile と書き，"d.go:23" のようなファイル名と行番号だけを簡潔に付与する．
	MyLshortfile

	// 使い方: log.LstdFlags と書き，初期値である Ldate | Ltime を一発で指定する．
	MyLstdFlags = MyLdate | MyLtime
)

// --- 2. 現代的な log/slog パッケージ (構造化ログ・Go 1.21+) ---

// 使い方: logger := slog.New(slog.NewJSONHandler(os.Stdout, nil)) と書き，キーと値のペアで構成される構造化ログを出力する．
// ログ解析ツール（ELK, CloudWatch Logs等）との親和性が極めて高い．
type MySlogLogger struct { /* 非公開 */ }

// 使い方: logger.Info("User login", "id", 123, "ip", "1.2.3.4") と書き，メッセージの後にキーと値を交互に渡す．
func (l *MySlogLogger) Info(msg string, args ...any) {}

// 使い方: logger.Error("DB connection failed", "err", err, "retries", 3) と書き，エラーの詳細を構造化して出力する．
func (l *MySlogLogger) Error(msg string, args ...any) {}

// 使い方: logger.Warn("Disk near full", slog.Int("usage", 90)) と書き，slog.Attr 型を使って型安全な値を渡す．
func (l *MySlogLogger) Warn(msg string, args ...any) {}

// 使い方: logger.DebugContext(ctx, "Fetch data", "url", url) と書き，context.Context に紐づいたトレースID等を含めたログを出す．
func (l *MySlogLogger) DebugContext(ctx context.Context, msg string, args ...any) {}

// 使い方: loggerWithUser := logger.With("user", "admin") と書き，共通のフィールドをあらかじめ付与した新しいロガーを作る．
func (l *MySlogLogger) With(args ...any) *slog.Logger { return nil }

// --- 3. slog のハンドラーと属性 ---

// 使い方: handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}) と書き，出力形式と最低ログレベルを設定する．
type MySlogJSONHandler struct{}

// 使い方: attr := slog.String("key", "value") と書き，型が明示された属性を作成する．
func MySlogString(key, value string) slog.Attr { return slog.Attr{} }

// 使い方: attr := slog.Group("request", "method", "GET", "path", "/api") と書き，複数のキーを階層構造にする．
func MySlogGroup(key string, args ...any) slog.Attr { return slog.Attr{} }

// --- 4. slog のレベル管理 ---

const (
	MyLevelDebug slog.Level = slog.LevelDebug // -4: 開発時の詳細情報
	MyLevelInfo  slog.Level = slog.LevelInfo  //  0: 通常の運用情報
	MyLevelWarn  slog.Level = slog.LevelWarn  //  4: 注意喚起が必要な事象
	MyLevelError slog.Level = slog.LevelError //  8: 致命的なエラー
)

func main() {
	println("Loaded: log spec.")
}
