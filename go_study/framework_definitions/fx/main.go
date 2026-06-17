package main

/*
   =============================================================================
   Go フレームワーク 定義集: fx (Dependency Injection)
   Uber 製の DI フレームワーク．ライフサイクル管理やモジュール化を容易にする．
   =============================================================================
*/

// --- 1. アプリケーションの構築 ---

// 使い方: fx.New(fx.Provide(NewDB), fx.Invoke(StartServer)).Run() と書き，依存関係を登録して実行する．
type MyFxApp struct { /* 非公開 */ }

func (a *MyFxApp) Run() {} // 使い方: 登録されたすべての Invoke を実行し，アプリを起動する．

// --- 2. 依存関係の提供と組織化 (Provide, Module) ---
// 使い方: fx.Provide(func() *DB { return &DB{} }) と書き，コンポーネントを登録する．
// 使い方: fx.Module("auth", fx.Provide(NewAuthService)) と書き，依存関係を機能単位でグループ化する．

// --- 3. アノテーションと装飾 (Annotate, Decorate) ---

// 使い方: fx.Annotate(NewHandler, fx.ParamTags(`group:"routes"`)) と書き，特定のグループに属する依存関係を一括で受け取る．
// 使い方: fx.Decorate(func(l *slog.Logger) *slog.Logger { ... }) と書き，既存の依存関係をラップして挙動を変更（デコレート）する．

// --- 4. 値グループ (Value Groups) ---
/*
   // 使い方: fx.In / fx.Out を使って，複数の依存関係をリストとしてまとめたり，名前を付けて区別したりする．
   type Params struct {
       fx.In
       Routes []Route `group:"routes"`
   }
*/

// --- 5. デバッグ ---
// 使い方: fx.DotGraph(&graph); fmt.Println(graph) と書き，依存関係グラフを DOT 形式で出力して可視化する．

// --- 6. 実行 (Invoke) ---
// 使い方: fx.Invoke(func(db *DB) { ... }) と書き，アプリ起動時に実行したい初期化処理やサーバー起動処理を記述する．
// 引数に渡した型は自動的に Provide されたものから注入される．

// --- 4. ライフサイクル ---
/*
   // 使い方: fx.Hook を使って，アプリの起動時（OnStart）と終了時（OnStop）の処理を記述する．
   lc.Append(fx.Hook{
       OnStart: func(context.Context) error { ... },
       OnStop:  func(context.Context) error { ... },
   })
*/

func main() {
	println("Loaded: fx definition.")
}
