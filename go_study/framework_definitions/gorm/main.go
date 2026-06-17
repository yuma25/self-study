package main

/*
   =============================================================================
   Go フレームワーク 定義集: GORM (ORM)
   Go のための最も人気のある ORM ライブラリ．
   =============================================================================
*/

// --- 1. 接続とマイグレーション ---

// 使い方: db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{}) と書き，DB に接続する．
type MyGORMDB struct { /* 非公開 */ }

// 使い方: db.AutoMigrate(&User{}) と書き，構造体からテーブルを自動作成・更新する．
func (db *MyGORMDB) AutoMigrate(dst ...any) error { return nil }

// --- 2. クエリ (CRUD) ---

// 使い方: db.Create(&user) と書き，新しいレコードを挿入する．
func (db *MyGORMDB) Create(value any) *MyGORMDB { return db }

// 使い方: db.First(&user, 1) と書き，主キーで1件取得する．
func (db *MyGORMDB) First(dest any, conds ...any) *MyGORMDB { return db }

// 使い方: db.Where("name = ?", "jinzhu").Find(&users) と書き，条件を指定して検索する．
func (db *MyGORMDB) Where(query any, args ...any) *MyGORMDB { return db }

// 使い方: db.Model(&user).Update("Name", "new_name") と書き，特定のフィールドを更新する．
func (db *MyGORMDB) Update(column string, value any) *MyGORMDB { return db }

// 使い方: db.Delete(&user) と書き，レコードを削除する．
func (db *MyGORMDB) Delete(value any, conds ...any) *MyGORMDB { return db }

// --- 3. 関連と高度なクエリ ---

// 使い方: db.Preload("Orders", "state = ?", "paid").Find(&users) と書き，条件付きで関連データをロードする．
func (db *MyGORMDB) Preload(query string, args ...any) *MyGORMDB { return db }

// 使い方: db.Joins("Company").Find(&users) と書き，JOIN を使って関連データを一度に取得する．
func (db *MyGORMDB) Joins(query string, args ...any) *MyGORMDB { return db }

// 使い方: db.Scopes(AmountGreaterThan1000).Find(&orders) と書き，共通のクエリ条件を関数として再利用する．
func (db *MyGORMDB) Scopes(funcs ...func(*MyGORMDB) *MyGORMDB) *MyGORMDB { return db }

// --- 4. トランザクション ---

// 使い方: db.Transaction(func(tx *gorm.DB) error { ... }) と書き，一連の操作をトランザクション内で実行する．エラーを返すと自動ロールバックされる．
func (db *MyGORMDB) Transaction(fc func(tx *MyGORMDB) error) error { return nil }

// --- 5. フック (Callbacks) ---
/*
   // 使い方: 構造体に BeforeCreate, AfterUpdate 等のメソッドを定義し，DB操作の前後に自動処理を挟む．
   func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
       u.UUID = uuid.New()
       return
   }
*/

// --- 6. 運用上の重要な機能 ---

// 使い方: 構造体に gorm.DeletedAt を含めることで，物理削除の代わりに削除フラグを立てる「論理削除（Soft Delete）」を有効にする．
type MyDeletedAt struct { /* 非公開 */ }

// 使い方: db.Raw("SELECT name FROM users WHERE id = ?", 1).Scan(&name) と書き，生のSQLを実行する．
func (db *MyGORMDB) Raw(sql string, values ...any) *MyGORMDB { return db }

// 使い方: db.Exec("DROP TABLE users") と書き，返り値（レコード）のないSQLを実行する．
func (db *MyGORMDB) Exec(sql string, values ...any) *MyGORMDB { return db }

// 使い方: newLogger := logger.New(...) と書き，スロークエリのログ出力やログレベルをカスタマイズする．

func main() {
	println("Loaded: GORM definition.")
}
