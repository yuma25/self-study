package main

/*
   =============================================================================
   Go フレームワーク 定義集: Fiber (Web Framework)
   Node.js の Express にインスパイアされた，極めて高速な（Fasthttpベース）フレームワーク．
   =============================================================================
*/

// --- 1. アプリケーション設定 ---

// 使い方: app := fiber.New(fiber.Config{ Prefork: true }) と書き，マルチプロセスによる高速化設定などを適用する．
type MyFiberApp struct { /* 非公開 */ }

func (a *MyFiberApp) Listen(addr string) error { return nil } // 使い方: app.Listen(":8080") で起動．

// 使い方: app.Use(logger.New()) と書き，ミドルウェアを適用する．
// 使い方: app.Group("/v1") と書き，ルートのグループ化を行う．

// --- 2. ハンドリング ---

// 使い方: app.Get("/", func(c *fiber.Ctx) error { ... }) と書き，Express ライクにルーティングを記述する．
type MyFiberCtx struct { /* 非公開 */ }

func (c *MyFiberCtx) Params(key string) string { return "" }
func (c *MyFiberCtx) Query(key string) string  { return "" }
func (c *MyFiberCtx) Cookies(key string) string { return "" } // 使い方: クッキーを取得する．

// 使い方: if err := c.BodyParser(&user); err != nil { ... } と書き，リクエストボディを構造体に変換する．
func (c *MyFiberCtx) BodyParser(out any) error { return nil }

// 使い方: return c.Status(200).JSON(data) と書き，メソッドチェーンでレスポンスを構築する．
func (c *MyFiberCtx) Status(status int) *MyFiberCtx { return c }
func (c *MyFiberCtx) JSON(data any) error           { return nil }
func (c *MyFiberCtx) SendString(s string) error     { return nil } // 使い方: 文字列を送信する．

// 使い方: c.Locals("user", userObj); val := c.Locals("user") と書き，リクエスト内で値を保持する．
func (c *MyFiberCtx) Locals(key any, value ...any) any { return nil }

// --- 3. ミドルウェア ---

// 使い方: app.Use(func(c *fiber.Ctx) error { return c.Next() }) と書き，共通処理を定義する．
// c.Next() を呼び出さない限り，後続のハンドラーは実行されない．

// --- 4. 高度な機能 ---

// 使い方: app.Static("/", "./public") と書き，静的ファイルを配信する．
// 使い方: app.Mount("/api", otherApp) と書き，別の Fiber アプリをマウント（サブアプリ化）する．
// 使い方: fiber.NewError(400, "bad request") と書き，エラーレスポンスを明示的に生成する．

func main() {
	println("Loaded: Fiber definition.")
}
