package main

import (
	"time"
)

/*
   =============================================================================
   Go 標準ライブラリ 定義集: time パッケージ (究極網羅・実戦コード例版)
   ※ 型衝突を避けるため，デモ用の型名には My 接頭辞を付けている．
   ※ 解説は「である調」とし，句読点は「．」「，」を使用している．
   ※ 構造体のフィールド以外は，コードの上部に使い方を記載している．
   =============================================================================
*/

// --- 1. 構造体 (実体) とそのメソッド ---

// 使い方: now := time.Now() と書き，プログラム内で特定の瞬間（時刻）を保持する実体として扱う．
type MyTime struct { /* 内部は非公開 */ }

// 使い方: y := t.Year(); m := t.Month(); d := t.Day() と書き，日付から特定の要素だけを抽出する．
func (t MyTime) Year() int         { return 0 }
func (t MyTime) Month() time.Month { return 1 }
func (t MyTime) Day() int          { return 1 }
func (t MyTime) Hour() int         { return 0 }
func (t MyTime) Minute() int       { return 0 }
func (t MyTime) Second() int       { return 0 }

// 使い方: timestamp := t.Unix() と書き，1970年1月1日からの経過秒数（Unixエポック）を取得し，DB保存などに利用する．
func (t MyTime) Unix() int64 { return 0 }

// 使い方: future := t.Add(10 * time.Minute) と書き，現在時刻の10分後の時刻を計算する．
func (t MyTime) Add(d time.Duration) MyTime { return MyTime{} }

// 使い方: nextWeek := t.AddDate(0, 0, 7) と書き，うるう年などを考慮しつつ安全に「7日後」を計算する．
func (t MyTime) AddDate(years int, months int, days int) MyTime { return MyTime{} }

// 使い方: elapsed := t1.Sub(t2) と書き，2つの時刻の差（どれだけの期間が経過したか）を MyDuration として取得する．
func (t MyTime) Sub(u MyTime) time.Duration { return 0 }

// 使い方: if t1.Before(t2) { ... } と書き，t1 が t2 より過去であるかを判定する．
func (t MyTime) Before(u MyTime) bool { return true }

// 使い方: if t1.After(t2) { ... } と書き，t1 が t2 より未来であるかを判定する．
func (t MyTime) After(u MyTime) bool { return true }

// 使い方: if t1.Equal(t2) { ... } と書き，タイムゾーンが異なっても「同じ瞬間」を指しているかを厳密に判定する．
func (t MyTime) Equal(u MyTime) bool { return true }

// 使い方: str := t.Format(time.RFC3339) と書き，時刻をAPI等で利用しやすい文字列にフォーマットする．
func (t MyTime) Format(layout string) string { return "" }

// 使い方: if t.IsZero() { ... } と書き，構造体に時刻がセットされていない（初期値のまま）かを判定する．
func (t MyTime) IsZero() bool { return true }

// 使い方: t = t.UTC() と書き，時刻を世界標準時（UTC）に変換する．
func (t MyTime) UTC() MyTime { return MyTime{} }

// 使い方: t = t.Local() と書き，時刻を実行環境のローカルタイムゾーンに変換する．
func (t MyTime) Local() MyTime { return MyTime{} }

// 使い方: d := 5 * time.Second と書き，関数に渡すタイムアウトなどの「期間」を表現する．
type MyDuration int64

// 使い方: fmt.Printf("Elapsed: %.2f sec\n", d.Seconds()) と書き，ナノ秒を読みやすい「秒」の浮動小数点に変換する．
func (d MyDuration) Seconds() float64 { return 0.0 }

// 使い方: fmt.Println(d.String()) と書き，"1h30m" のような直感的な文字列形式を出力する．
func (d MyDuration) String() string { return "" }

// 使い方: timer := time.NewTimer(time.Second); <-timer.C と書き，指定時間後に一度だけ発火するイベントを待つ．
type MyTimer struct {
	C <-chan time.Time // 使い方: select { case <-timer.C: ... } と書き，発火時の時刻を受け取るチャネルである．
}

// 使い方: if !timer.Stop() { <-timer.C } と書き，発火前のタイマーを安全に停止し，不要なゴルーチンのリークを防ぐ．
func (t *MyTimer) Stop() bool { return true }

// 使い方: ticker := time.NewTicker(time.Second); for t := range ticker.C { ... } と書き，一定間隔で処理を繰り返す．
type MyTicker struct {
	C <-chan time.Time // 使い方: 定期的に現在時刻が流れてくるチャネルである．
}

// 使い方: defer ticker.Stop() と書き，用済みになったティッカーを止め，リソースを解放する．
func (t *MyTicker) Stop() {}

// 使い方: loc, _ := time.LoadLocation("Asia/Tokyo"); t = t.In(loc) と書き，特定の地域に基づくタイムゾーン情報を扱う．
type MyLocation struct { /* 非公開 */ }

// --- 2. パッケージレベルの関数 (操作) ---

// 使い方: now := time.Now() と書き，その瞬間のローカル時刻を取得する．
func MyNow() MyTime { return MyTime{} }

// 使い方: time.Sleep(2 * time.Second) と書き，現在のゴルーチンを一時停止させる．
func MySleep(d time.Duration) {}

// 使い方: elapsed := time.Since(start) と書き，処理の開始時刻 (start) から現在までの経過時間を計測する．
func MySince(t MyTime) time.Duration { return 0 }

// 使い方: remain := time.Until(deadline) と書き，現在から期限 (deadline) までの残り時間を計算する．
func MyUntil(t MyTime) time.Duration { return 0 }

// 使い方: t, err := time.Parse(time.RFC3339, "2023-01-01T12:00:00Z") と書き，文字列を Time 実体に変換する．
func MyParse(layout, value string) (MyTime, error) { return MyTime{}, nil }

// 使い方: t, _ := time.ParseInLocation("2006-01-02", "2023-01-01", loc) と書き，タイムゾーンを考慮して文字列を解析する．
func MyParseInLocation(layout, value string, loc *MyLocation) (MyTime, error) {
	return MyTime{}, nil
}

// 使い方: loc, err := time.LoadLocation("Asia/Tokyo") と書き，IANAタイムゾーンデータベースからロケーションを取得する．
func MyLoadLocation(name string) (*MyLocation, error) { return nil, nil }

// 使い方: select { case <-time.After(10 * time.Second): ... } と書き，簡易的なタイムアウト処理を実装する．
func MyAfter(d time.Duration) <-chan time.Time { return nil }

// 使い方: ticker := time.NewTicker(time.Minute) と書き，定期実行のためのティッカーを生成する．
func MyNewTicker(d time.Duration) *MyTicker { return nil }

// --- 3. 便利な定数 ---

const (
	// 使い方: timeout := 100 * time.Millisecond と書き，直感的にミリ秒単位の期間を指定する．
	MyMillisecond time.Duration = 1000 * 1000 // (1000 * time.Microsecond)
	// 使い方: duration := 5 * time.Second と書き，秒単位の期間を計算する．
	MySecond time.Duration = 1000 * MyMillisecond
	// 使い方: wait := 2 * time.Minute と書き，分単位の期間を計算する．
	MyMinute time.Duration = 60 * MySecond
	// 使い方: limit := 24 * time.Hour と書き，時間単位の期間を計算する．
	MyHour time.Duration = 60 * MyMinute

	// 使い方: t.Format(time.RFC3339) と書き，Web API等で多用される標準的なタイムスタンプ形式にする．
	MyRFC3339 = time.RFC3339
	// 使い方: t.Format(time.DateOnly) と書き，"YYYY-MM-DD" の日付形式にする (Go 1.20+)．
	MyDateOnly = time.DateOnly
	// 使い方: t.Format(time.TimeOnly) と書き，"HH:MM:SS" の時刻形式にする (Go 1.20+)．
	MyTimeOnly = time.TimeOnly
)

func main() {
	println("Updated: Full time spec with explicit 'How to use' code examples loaded.")
}
