package main

import (
	"context"
	"sync"
	"time"
)

/*
   =============================================================================
   Go 標準ライブラリ 定義集: sync / context パッケージ (究極網羅・実戦コード例版)
   =============================================================================
*/

// --- 1. sync パッケージ (並行処理・同期プリミティブ) ---

// 使い方: var mu sync.Mutex; mu.Lock(); defer mu.Unlock() と書き，複数のゴルーチンから同時にアクセスされないよう保護する（排他ロック）．
type MyMutex struct {
	sync.Mutex
}

func (m *MyMutex) MyLock()   { m.Lock() }
func (m *MyMutex) MyUnlock() { m.Unlock() }

// 使い方: if mu.TryLock() { defer mu.Unlock(); ... } と書き，ロックが即座に取得できる場合のみ処理を行う（非ブロッキング・Go 1.18+）．
func (m *MyMutex) MyTryLock() bool { return m.TryLock() }

// 使い方: var rw sync.RWMutex; rw.RLock(); defer rw.RUnlock() と書き，読み取りは複数許可し，書き込み時のみ排他する（読取/書込ロック）．
type MyRWMutex struct {
	sync.RWMutex
}

func (rw *MyRWMutex) MyRLock()   { rw.RLock() }
func (rw *MyRWMutex) MyRUnlock() { rw.RUnlock() }
func (rw *MyRWMutex) MyLock()    { rw.Lock() }
func (rw *MyRWMutex) MyUnlock()  { rw.Unlock() }

// 使い方: if rw.TryLock() { ... } または if rw.TryRLock() { ... } と書き，待機せずにロック取得を試みる（Go 1.18+）．
func (rw *MyRWMutex) MyTryLock() bool  { return rw.TryLock() }
func (rw *MyRWMutex) MyTryRLock() bool { return rw.TryRLock() }

// 使い方: var wg sync.WaitGroup; wg.Add(1); go func() { defer wg.Done(); ... }(); wg.Wait() と書き，すべてのゴルーチンの完了を待つ．
type MyWaitGroup struct {
	sync.WaitGroup
}

func (wg *MyWaitGroup) MyAdd(delta int) { wg.Add(delta) }
func (wg *MyWaitGroup) MyDone()         { wg.Done() }
func (wg *MyWaitGroup) MyWait()         { wg.Wait() }

// 使い方: var once sync.Once; once.Do(func() { ... }) と書き，複数のゴルーチンから呼ばれても一生に一度だけ実行する．
type MyOnce struct {
	sync.Once
}

func (o *MyOnce) MyDo(f func()) { o.Do(f) }

// 使い方: f := sync.OnceFunc(func() { ... }) と書き，呼び出すと一度だけ実行される関数を生成する（Go 1.21+）．
func MyOnceFunc(f func()) func() { return sync.OnceFunc(f) }

// 使い方: v := sync.OnceValue(func() T { ... }) と書き，結果をキャッシュして返す関数を生成する（Go 1.21+）．
func MyOnceValue[T any](f func() T) func() T { return sync.OnceValue(f) }

// 使い方: var p = sync.Pool{New: func() any { return make([]byte, 1024) }}; obj := p.Get(); p.Put(obj) と書き，オブジェクトを再利用して GC 負荷を減らす．
type MyPool struct {
	sync.Pool
}

func (p *MyPool) MyGet() any  { return p.Get() }
func (p *MyPool) MyPut(x any) { p.Put(x) }

// 使い方: var m sync.Map; m.Store("k", "v"); v, ok := m.Load("k") と書き，並行安全なマップを使用する．
type MyMap struct {
	sync.Map
}

// 使い方: cond := sync.NewCond(&sync.Mutex{}); cond.L.Lock(); cond.Wait(); cond.L.Unlock() と書き，条件が満たされるまで待機する．
type MyCond struct {
	*sync.Cond
}

func (c *MyCond) MyWait()      { c.Wait() }
func (c *MyCond) MySignal()    { c.Signal() }
func (c *MyCond) MyBroadcast() { c.Broadcast() }

// --- 2. context パッケージ (キャンセル信号・タイムアウト・値の伝播) ---

// 使い方: func MyFunc(ctx context.Context, ...) と書き，リクエストの生存期間やキャンセル信号を伝播させる．
type MyContext interface {
	Done() <-chan struct{}
	Err() error
	Deadline() (deadline time.Time, ok bool)
	Value(key any) any
}

// --- context 標準エラーと定数 ---

var (
	MyCanceled         = context.Canceled         // 使い方: 手動キャンセルされたか確認する．
	MyDeadlineExceeded = context.DeadlineExceeded // 使い方: タイムアウトしたか確認する．
)

// --- context パッケージレベルの関数 ---

// 使い方: ctx := context.Background() と書き，根源となる空のコンテキストを作成する．
func MyBackground() context.Context { return context.Background() }

// 使い方: ctx := context.TODO() と書き，Contextが必要だがまだ設計が決まっていない場合に使用する．
func MyTODO() context.Context { return context.TODO() }

// 使い方: ctx, cancel := context.WithCancel(parent) と書き，手動キャンセル可能なコンテキストを作る．
func MyWithCancel(parent context.Context) (context.Context, context.CancelFunc) {
	return context.WithCancel(parent)
}

// 使い方: ctx, cancel := context.WithCancelCause(parent) と書き，エラー理由付きでキャンセル可能にする（Go 1.20+）．
func MyWithCancelCause(parent context.Context) (context.Context, context.CancelCauseFunc) {
	return context.WithCancelCause(parent)
}

// 使い方: err := context.Cause(ctx) と書き，キャンセル理由のエラーを取得する（Go 1.20+）．
func MyCause(ctx context.Context) error { return context.Cause(ctx) }

// 使い方: ctx, cancel := context.WithTimeout(parent, 5*time.Second); defer cancel() と書き，自動タイムアウトを設定する．
func MyWithTimeout(parent context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(parent, timeout)
}

// 使い方: ctx := context.WithValue(parent, "key", "val") と書き，リクエストスコープの値を格納する．
func MyWithValue(parent context.Context, key, val any) context.Context {
	return context.WithValue(parent, key, val)
}

// 使い方: ctx := context.WithoutCancel(parent) と書き，親のキャンセルを継承しないコンテキストを作る（Go 1.21+）．
func MyWithoutCancel(parent context.Context) context.Context {
	return context.WithoutCancel(parent)
}

// 使い方: stop := context.AfterFunc(ctx, func() { ... }) と書き，キャンセル後に実行する関数を登録する（Go 1.21+）．
func MyAfterFunc(ctx context.Context, f func()) (stop func() bool) {
	return context.AfterFunc(ctx, f)
}

func main() {
	println("Loaded: sync and context specs.")
}
