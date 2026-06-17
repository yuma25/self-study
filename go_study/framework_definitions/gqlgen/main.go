package main

/*
   =============================================================================
   Go フレームワーク 定義集: gqlgen (GraphQL)
   スキーマ優先 (Schema First) の GraphQL ライブラリ．型安全なリゾルバを自動生成する．
   =============================================================================
*/

// --- 1. 初期設定 ---
// 使い方: go run github.com/99designs/gqlgen init と書き，設定ファイルとリゾルバの雛形を作成する．

// --- 2. リゾルバと高度な機能 ---
/*
   // 使い方: schema.graphqls に定義した Query や Mutation に対応するメソッドを生成された resolver.go に実装する．
   func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
       return db.FindUser(id), nil
   }

   // 使い方: Dataloader を実装して，GraphQL 特有の N+1 問題（関連データの大量個別クエリ）を回避する．
   // 使い方: Directive (@auth 等) をスキーマに記述し，認可処理をリゾルバの外部で一元管理する．
   // 使い方: gqlgen.yml を編集して，特定の型を Go の既存型や独自 Scalar にマッピングする．
   // 使い方: union SearchResult = User | Post のように定義し，複数の型を返す柔軟なクエリを実装する．
   // 使い方: Complexity 設定を導入して，ネストの深すぎる悪意のあるクエリからサーバーを保護する．
*/

// --- 3. サーバーの起動 ---
// 使い方: srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{}})) と書き，
// http.Handle("/query", srv) で GraphQL エンドポイントを公開する．

func main() {
	println("Loaded: gqlgen definition.")
}
