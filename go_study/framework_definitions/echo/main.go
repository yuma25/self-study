package main

/*
   =============================================================================
   Go フレームワーク 定義集: Echo (Web Framework)
   シンプルで高い拡張性を持ち，カスタマイズ性が高い人気のフレームワーク．
   =============================================================================
*/

// --- 1. インスタンスとルーティング ---

// 使い方: e := echo.New() と書き，新しい Echo インスタンスを作成する．
type MyEcho struct { /* 非公開 */ }

func (e *MyEcho) Start(address string) error { return nil } // 使い方: e.Start(":8080") でサーバーを起動する．

// 使い方: e.GET("/users/:id", handler) と書き，パスパラメータを含むルートを定義する．
func (e *MyEcho) GET(path string, h any, m ...any) {}

// 使い方: e.Validator = &MyValidator{} と書き，構造体バリデーションの仕組みを登録する．
type MyEchoValidator interface {
	Validate(i any) error
}

// 使い方: e.HTTPErrorHandler = func(err error, c echo.Context) { ... } と書き，エラー発生時の共通処理をカスタマイズする．
type MyEchoErrorHandler func(error, any)

// --- 2. コンテキストとバインド ---

// 使い方: func(c echo.Context) error { ... } と書き，リクエスト処理を記述する．
type MyEchoContext interface {
	Param(name string) string      // 使い方: c.Param("id") でパスパラメータを取得する．
	QueryParam(name string) string // 使い方: c.QueryParam("search") でクエリ文字列を取得する．
	FormValue(name string) string  // 使い方: フォームデータを取得する．

	// 使い方: if err := c.Bind(&user); err != nil { ... } と書き，JSON/Form/Query などを構造体に自動マッピングする．
	Bind(i any) error

	// 使い方: if err := c.Validate(&user); err != nil { ... } と書き，登録した Validator で値を検証する．
	Validate(i any) error

	// 使い方: return c.JSON(200, data) と書き，レスポンスを返す．返り値として error を返すのが Echo の特徴．
	JSON(code int, i any) error
	String(code int, s string) error
	NoContent(code int) error // 使い方: 204 No Content 等，ボディ無しのレスポンスを返す．
	Blob(code int, contentType string, b []byte) error // 使い方: バイナリデータを返す．

	// 使い方: c.Set("key", val); v := c.Get("key") と書き，コンテキスト内でデータを共有する．
	Set(key string, val any)
	Get(key string) any
}

/*
   [カスタムコンテキストのパターン]
   type CustomContext struct {
       echo.Context
   }
   // 使い方: 独自のメソッド（GetUser()等）を追加した構造体で echo.Context をラップし，
   // ミドルウェアで差し替えることで，ハンドラー内での型安全なデータアクセスを可能にする．
*/

// --- 3. 静的ファイル ---

// 使い方: e.Static("/static", "assets") と書き，静的コンテンツを配信する．
func (e *MyEcho) Static(prefix, root string) {}

// --- 4. ミドルウェア ---

// 使い方: e.Use(middleware.Logger(), middleware.Recover()) と書き，標準のミドルウェアを使用する．
// 使い方: e.Pre(middleware.AddTrailingSlash()) と書き，ルーティングの「前」に実行したい処理を登録する．

func main() {
	println("Loaded: Echo definition.")
}
