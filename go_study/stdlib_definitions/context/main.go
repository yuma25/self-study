package main

import (
	"context"
	"time"
)

/*
   =============================================================================
   Go 標準ライブラリ 定義集: context パッケージ (究極網羅・実戦コード例版)
   =============================================================================
*/

// --- 1. インターフェースと基本型 ---

// 使い方: func MyFunc(ctx context.Context, ...) と書き，API境界を越えてキャンセル信号やタイムアウト，値を伝播させる．
// 全てのブロッキング操作の第一引数に渡すのが Go の慣習である．
type MyContext interface {
	// 使い方: select { case <-ctx.Done(): ... } と書き，キャンセルやタイムアウトを検知する．
	Done() <-chan struct{}

	// 使い方: err := ctx.Err(); if err == context.Canceled { ... } と書き，Done が閉じた理由を確認する．
	Err() error

	// 使い方: if d, ok := ctx.Deadline(); ok { ... } と書き，タイムアウトまでの残り時間を確認する．
	Deadline() (deadline time.Time, ok bool)

	// 使い方: val := ctx.Value("key") と書き，リクエストスコープのデータを取得する．
	Value(key any) any
}

// 使い方: ctx, cancel := context.WithCancel(parent) と書き，手動でキャンセル可能なコンテキストを作成する．
// 終了時には必ず defer cancel() を呼び出してリソースを解放すること．
type MyCancelFunc func()

// 使い方: ctx, cancel := context.WithCancelCause(parent) と書き，エラーを伴うキャンセルを可能にする（Go 1.20+）．
// cancel(err) のように呼び出すと，Cause(ctx) でそのエラーを取得できる．
type MyCancelCauseFunc func(cause error)

// --- 2. 標準エラー ---

// 使い方: if err == context.Canceled { ... } と書き，コンテキストが手動でキャンセルされたか確認する．
var MyCanceled error = context.Canceled

// 使い方: if err == context.DeadlineExceeded { ... } と書き，タイムアウト（期限切れ）で終了したか確認する．
var MyDeadlineExceeded error = context.DeadlineExceeded

// --- 3. コンテキストの生成 (ルート) ---

// 使い方: ctx := context.Background() と書き，メイン関数や初期化時，テストの起点となる空のコンテキストを作成する．
func MyBackground() context.Context { return nil }

// 使い方: ctx := context.TODO() と書き，どのコンテキストを使うべきか未定の場合や，将来的に追加予定の場所で暫定的に使用する．
func MyTODO() context.Context { return nil }

// --- 3. 派生コンテキストの作成 (キャンセル・タイムアウト) ---

// 使い方: ctx, cancel := context.WithCancel(parent) と書き，cancel() を呼ぶことで派生先全てに終了信号を送る．
func MyWithCancel(parent context.Context) (ctx context.Context, cancel context.CancelFunc) {
	return nil, nil
}

// 使い方: ctx, cancel := context.WithCancelCause(parent) と書き，cancel(err) でキャンセル理由を添えて終了させる（Go 1.20+）．
func MyWithCancelCause(parent context.Context) (ctx context.Context, cancel context.CancelCauseFunc) {
	return nil, nil
}

// 使い方: err := context.Cause(ctx) と書き，WithCancelCause 等で設定された具体的なエラー理由を取得する（Go 1.20+）．
func MyCause(ctx context.Context) error { return nil }

// 使い方: ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second)) と書き，指定時刻に自動でキャンセルさせる．
func MyWithDeadline(parent context.Context, d time.Time) (context.Context, context.CancelFunc) {
	return nil, nil
}

// 使い方: ctx, cancel := context.WithTimeout(parent, 5*time.Second) と書き，一定時間経過後に自動でキャンセルさせる．
func MyWithTimeout(parent context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	return nil, nil
}

// --- 4. 値の伝播 ---

// 使い方: ctx := context.WithValue(parent, "userID", 123) と書き，リクエスト固有の値をコンテキストに格納する．
// キーには衝突を避けるため，エクスポートされない独自の型（type key struct{}）を使うのが鉄則．
func MyWithValue(parent context.Context, key, val any) context.Context { return nil }

// 使い方: ctx := context.WithoutCancel(parent) と書き，親のキャンセル信号を継承しないコンテキストを作成する（Go 1.21+）．
// 親がキャンセルされてもこの ctx は生き残り，ログ出力やクリーンアップ処理を継続させたい場合に有効．
func MyWithoutCancel(parent context.Context) context.Context { return nil }

// --- 5. モダンな機能 (Go 1.21+) ---

// 使い方: stop := context.AfterFunc(ctx, func() { ... }) と書き，ctx がキャンセルされた後に実行する関数を登録する（Go 1.21+）．
func MyAfterFunc(ctx context.Context, f func()) (stop func() bool) { return nil }

// 使い方: ctx, cancel := context.WithTimeoutCause(parent, timeout, err) と書き，タイムアウト時に独自のエラーを理由として設定する（Go 1.21+）．
func MyWithTimeoutCause(parent context.Context, timeout time.Duration, cause error) (context.Context, context.CancelFunc) {
	return nil, nil
}

func main() {
	println("Loaded: context spec.")
}
