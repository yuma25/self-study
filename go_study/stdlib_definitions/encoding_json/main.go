package main

import (
	"bytes"
	"encoding/json"
)

/*
   =============================================================================
   Go 標準ライブラリ 定義集: encoding/json パッケージ (究極網羅・実戦コード例版)
   =============================================================================
*/

// --- 1. 構造体 (実体) とそのメソッド ---

// 使い方: encoder := json.NewEncoder(os.Stdout); encoder.Encode(data) と書き，メモリを節約しつつストリームへ直接JSONを書き出す．
// 巨大なデータの送信や，ファイル保存時など，io.Writer を対象とする場合に必須の部品である．
type MyEncoder struct { /* 非公開 */ }

func (enc *MyEncoder) Encode(v any) error { return nil }

// 使い方: enc.SetIndent("", "  ") と書き，出力されるJSONを人間が読みやすいように整形（プリティプリント）する．
func (enc *MyEncoder) SetIndent(prefix, indent string) {}

// 使い方: enc.SetEscapeHTML(false) と書き，"<" などを "\u003c" にエスケープする標準の安全機構を無効化する．
func (enc *MyEncoder) SetEscapeHTML(on bool) {}

// 使い方: decoder := json.NewDecoder(r.Body); decoder.Decode(&user) と書き，HTTPリクエスト等から逐次JSONを読み込む．
// メモリ上に一度全データをロードする必要がないため，サーバーの負荷を劇的に下げる．
type MyDecoder struct { /* 非公開 */ }

func (dec *MyDecoder) Decode(v any) error { return nil }

// 使い方: dec.DisallowUnknownFields() と書き，構造体に定義されていない未知のフィールドがJSONに含まれていた場合にエラーにする．安全なAPI設計に有効である．
func (dec *MyDecoder) DisallowUnknownFields() {}

// 使い方: dec.UseNumber() と書き，JSONの数値をデフォルトの float64 ではなく json.Number 型として読み込ませ，精度落ちを防ぐ．
func (dec *MyDecoder) UseNumber() {}

// 使い方: if dec.More() { ... } と書き，ストリーム（配列等）にまだ読み込める要素があるか確認する．
func (dec *MyDecoder) More() bool { return false }

// 使い方: token, err := dec.Token() と書き，JSONのデリミタ（[, ], {, }）や基本値を1つずつトークンとして読み込む．巨大な配列のストリーム処理に必須．
func (dec *MyDecoder) Token() (json.Token, error) { return nil, nil }

// --- 2. パッケージレベルの関数 (操作) ---

// 使い方: jsonData, err := json.Marshal(user) と書き，Goの構造体やマップを []byte 形式のJSONデータに変換する．
// メモリ上にデータを一気に確保するため，小〜中規模のデータ変換に適している．
func MyMarshal(v any) ([]byte, error) { return nil, nil }

// 使い方: jsonData, _ := json.MarshalIndent(user, "", "  ") と書き，最初から整形された（インデント付きの）JSONを生成する．
func MyMarshalIndent(v any, prefix, indent string) ([]byte, error) { return nil, nil }

// 使い方: err := json.Unmarshal(jsonData, &user) と書き，[]byte のJSONデータを，あらかじめ用意した構造体のポインタに流し込む．
func MyUnmarshal(data []byte, v any) error { return nil }

// 使い方: if json.Valid(data) { ... } と書き，与えられた []byte が構文的に正しいJSONかどうかを，完全なパースを伴わずに高速に判定する．
func MyValid(data []byte) bool { return false }

// 使い方: err := json.Compact(&buf, data) と書き，JSONデータから不要な空白や改行をすべて取り除き，最小サイズに圧縮する．
func MyCompact(dst *bytes.Buffer, src []byte) error { return nil }

// 使い方: err := json.Indent(&buf, data, "", "  ") と書き，既存の []byte JSONデータにインデントを追加して整形する．
func MyIndent(dst *bytes.Buffer, src []byte, prefix, indent string) error { return nil }

// 使い方: json.HTMLEscape(&buf, data) と書き，JSON内の "<", ">", "&" などを安全にエスケープする．
func MyHTMLEscape(dst *bytes.Buffer, src []byte) {}

// --- 3. 構造体タグ (Struct Tags) ---
// Goのフィールド定義の末尾に `json:"..."` の形式で記述する．

/*
   - `json:"name"`          : JSON上のキー名を "name" に指定する．
   - `json:"name,omitempty"` : 値がゼロ値（0, "", nil等）の場合，JSON出力からその項目を除外する．
   - `json:"-"`             : そのフィールドを常にJSON出力・入力から除外する（秘密情報など）．
   - `json:"-,"`            : キー名を "-" という文字列として扱う（特殊ケース）．
   - `json:",string"`       : 数値や真偽値を，JSON上では文字列（"100", "true"）として扱う．JavaScriptの精度問題対策に有効．
*/

// --- 4. 特殊な型 ---

// 使い方: json.RawMessage(`{"key": "value"}`) と書き，JSONのパース処理を後回し（遅延評価）にしたいフィールドに使用する．
// APIレスポンスの中で，中身の構造が変わる動的な部分（プラグインの設定等）をとりあえずそのまま保持したい時に用いる．
type MyRawMessage []byte

func (m MyRawMessage) MarshalJSON() ([]byte, error) { return m, nil }
func (m *MyRawMessage) UnmarshalJSON(data []byte) error { return nil }

// 使い方: var n json.Number; n.Int64() と書き，JSONから送られてきた数値を，丸め誤差なく安全に文字列として保持し，後から変換する．
type MyNumber string

func (n MyNumber) String() string { return string(n) }
func (n MyNumber) Float64() (float64, error) { return 0.0, nil }
func (n MyNumber) Int64() (int64, error) { return 0, nil }

// --- 4. 重要なインターフェース ---

// 使い方: 独自の構造体に MarshalJSON を実装し，デフォルトの変換ルールを無視して独自の JSON 形式を出力させる．
type MyMarshaler interface {
	MarshalJSON() ([]byte, error)
}

// 使い方: 独自の構造体に UnmarshalJSON を実装し，受け取った JSON 文字列から自作の複雑なパース処理を行って値をセットする．
type MyUnmarshaler interface {
	UnmarshalJSON([]byte) error
}

func main() {
	println("Loaded: encoding/json spec.")
}
