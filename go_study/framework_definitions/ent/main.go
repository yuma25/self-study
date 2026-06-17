package main

/*
   =============================================================================
   Go フレームワーク 定義集: ent (Entity Framework)
   Facebook (Meta) 製のエンティティベースの ORM．コード生成により型安全なクエリが可能．
   =============================================================================
*/

// --- 1. スキーマ定義 ---
// 使い方: go run -mod=mod entgo.io/ent/cmd/ent new User と書き，スキーマの雛形を作成する．
// ent/schema/user.go にフィールドやエッジ（関連）を記述する．

// --- 2. クライアントとクエリ ---

// 使い方: client, _ := ent.Open("mysql", dsn) と書き，クライアントを作成する．
type MyEntClient struct { /* 非公開 */ }

// 使い方: u, _ := client.User.Create().SetName("a").Save(ctx) と書き，型安全にレコードを作成する．
// 使い方: users, _ := client.User.Query().Where(user.AgeGT(18)).All(ctx) と書き，流れるようなインターフェースで検索する．
type MyUserQuery struct { /* 非公開 */ }

func (q *MyUserQuery) Where(ps ...any) *MyUserQuery { return q }
func (q *MyUserQuery) Limit(n int) *MyUserQuery     { return q }
func (q *MyUserQuery) All(ctx any) ([]any, error)   { return nil, nil }

// --- 3. フックとプライバシー ---
// 使い方: client.User.Use(func(next ent.Mutator) ent.Mutator { ... }) と書き，ミドルウェアのように処理を挟む．
// 使い方: Privacy レイヤーを設定して，行レベルのアクセス制御を自動化する．

// --- 4. インターセプター (Interceptors) ---
// 使い方: client.Intercept(func(next ent.Querier) ent.Querier { ... }) と書き，読み取りクエリ全体に共通処理（ソフトデリートのフィルタリング等）を適用する．

// --- 5. スキーマ定義のヒント ---
/*
   func (User) Fields() []ent.Field {
       return []ent.Field{
           field.String("name").Unique(),
           field.Int("age").Positive(),
       }
   }
   func (User) Indexes() []ent.Index {
       return []ent.Index{
           index.Fields("name", "age").Unique(),
       }
   }
*/

// --- 6. マイグレーション ---
// 使い方: client.Schema.Create(ctx, migrate.WithForeignKeys(false)) と書き，マイグレーション時の挙動を細かく制御する．

func main() {
	println("Loaded: ent definition.")
}
