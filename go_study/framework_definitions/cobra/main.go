package main

/*
   =============================================================================
   Go フレームワーク 定義集: Cobra (CLI Framework)
   kubectl や hugo でも使われている，強力な CLI アプリケーション作成ツール．
   =============================================================================
*/

// --- 1. コマンド定義 ---

// 使い方: var rootCmd = &cobra.Command{ Use: "app", Run: func(cmd *cobra.Command, args []string) { ... } } と書き，コマンドを定義する．
type MyCobraCommand struct {
	Use   string // 使い方: コマンド名と引数の説明．
	Short string // 使い方: 短い説明文．
	Long  string // 使い方: 詳細な説明文．

	// 使い方: 引数の数をチェックする（cobra.ExactArgs(1) など）．
	Args any

	// 使い方: コマンド実行前に呼ばれる（サブコマンドでも継承される）．
	PersistentPreRun func(cmd *MyCobraCommand, args []string)

	// 使い方: 実際の処理内容．
	Run func(cmd *MyCobraCommand, args []string)
}

func (c *MyCobraCommand) Execute() error { return nil } // 使い方: rootCmd.Execute() で CLI を起動する．

// 使い方: rootCmd.AddCommand(childCmd) と書き，サブコマンドを追加してネスト構造を作る．
func (c *MyCobraCommand) AddCommand(cmds ...*MyCobraCommand) {}

// --- 2. フラグ (引数) ---

// 使い方: rootCmd.PersistentFlags().StringVar(&v, ...) と書き，自身と全ての子孫で使える共通フラグを定義する．
// 使い方: rootCmd.Flags().StringVarP(&name, "name", "n", "default", "description") と書き，特定のコマンド限定のフラグを定義する．
type MyFlags struct { /* 非公開 */ }

func (f *MyFlags) StringVarP(p *string, name, shorthand, value, usage string) {}
func (f *MyFlags) BoolVar(p *bool, name string, value bool, usage string)     {}

// 使い方: rootCmd.MarkFlagRequired("name") と書き，特定のフラグを必須にする．
func (f *MyFlags) MarkRequired(name string) error { return nil }

// --- 3. Viper との連携 ---
// 使い方: viper.BindPFlag("name", rootCmd.Flags().Lookup("name")) と書き，設定ファイルや環境変数と CLI 引数を同期させるのが一般的．

// --- 4. 補完とドキュメント ---
// 使い方: rootCmd.GenBashCompletionFile("bash_completion.sh") と書き，シェル補完ファイルを自動生成する．
// 使い方: cobra.CheckErr(err) と書き，エラーが発生した場合にメッセージを表示して終了する．

func main() {
	println("Loaded: Cobra definition.")
}
