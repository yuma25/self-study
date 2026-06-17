package main

import (
	"sync"
)

/*
   =============================================================================
   Go 標準ライブラリ 定義集: sync パッケージ (究極網羅・実戦コード例版)
   =============================================================================
*/

// --- 1. 排他制御 (Mutex) ---

// 使い方: var mu sync.Mutex; mu.Lock(); defer mu.Unlock() と書き，複数のゴルーチンから同時にアクセスされないよう保護する（排他ロック）．
type MyMutex struct {
	sync.Mutex
}

// 使い方: if mu.TryLock() { defer mu.Unlock(); ... } と書き，ロックが即座に取得できる場合のみ処理を行う（非ブロッキング・Go 1.18+）．
func (m *MyMutex) MyTryLock() bool { return m.TryLock() }

// 使い方: var rw sync.RWMutex; rw.RLock(); defer rw.RUnlock() と書き，読み取りは複数許可し，書き込み時のみ排他する（読取/書込ロック）．
type MyRWMutex struct {
	sync.RWMutex
}

// 使い方: if rw.TryLock() { ... } または if rw.TryRLock() { ... } と書き，待機せずにロック取得を試みる（Go 1.18+）．
func (rw *MyRWMutex) MyTryLock() bool  { return rw.TryLock() }
func (rw *MyRWMutex) MyTryRLock() bool { return rw.TryRLock() }

// --- 2. 同期・待機 (WaitGroup, Cond) ---

// 使い方: var wg sync.WaitGroup; wg.Add(1); go func() { defer wg.Done(); ... }(); wg.Wait() と書き，すべてのゴルーチンの完了を待つ．
type MyWaitGroup struct {
	sync.WaitGroup
}

// 使い方: cond := sync.NewCond(&sync.Mutex{}); cond.L.Lock(); cond.Wait(); cond.L.Unlock() と書き，特定の条件が満たされるまでゴルーチンを待機させる．
type MyCond struct {
	*sync.Cond
}

// 使い方: cond.Signal() と書き，待機しているゴルーチンのうち「1つ」を再開させる．
func (c *MyCond) MySignal() { c.Signal() }

// 使い方: cond.Broadcast() と書き，待機している「すべての」ゴルーチンを再開させる．
func (c *MyCond) MyBroadcast() { c.Broadcast() }

// --- 3. 一回限りの実行 (Once) ---

// 使い方: var once sync.Once; once.Do(func() { ... }) と書き，複数のゴルーチンから呼ばれても，指定した関数を一生に一度だけ実行する（初期化処理等）．
type MyOnce struct {
	sync.Once
}

// 使い方: f := sync.OnceFunc(func() { ... }) と書き，呼び出すと一度だけ実行される関数を生成する（Go 1.21+）．
func MyOnceFunc(f func()) func() { return sync.OnceFunc(f) }

// 使い方: v := sync.OnceValue(func() T { ... }) と書き，最初の呼び出し結果をキャッシュして返す関数を生成する（Go 1.21+）．
func MyOnceValue[T any](f func() T) func() T { return sync.OnceValue(f) }

// 使い方: v1, v2 := sync.OnceValues(func() (T1, T2) { ... }) と書き，2つの返り値をキャッシュして返す関数を生成する（Go 1.21+）．
func MyOnceValues[T1, T2 any](f func() (T1, T2)) func() (T1, T2) { return sync.OnceValues(f) }

// --- 4. 並行マップとプール (Map, Pool) ---

// 使い方: var m sync.Map; m.Store("key", "val"); val, ok := m.Load("key") と書き，複数のゴルーチンから安全にアクセスできるマップを使用する．
// 頻繁な書き込みには向かないが，一度書いたものを大量に読み出すケース（キャッシュ等）で標準の map + Mutex より高速．
type MyMap struct {
	sync.Map
}

// 使い方: var p = sync.Pool{New: func() any { return new(T) }}; obj := p.Get().(*T); defer p.Put(obj) と書き，オブジェクトを再利用して GC の負荷を軽減する．
type MyPool struct {
	sync.Pool
}

// --- 5. インターフェース ---

// 使い方: func DoSomething(l sync.Locker) { l.Lock(); defer l.Unlock(); ... } と書き，Mutex と RWMutex の両方を受け入れられるようにする．
type MyLocker interface {
	Lock()
	Unlock()
}

func main() {
	println("Loaded: sync spec.")
}
