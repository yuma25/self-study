# Study (Go & TypeScript)

このリポジトリは，Go言語とTypeScriptの学習用サンプルコードをまとめたレポジトリである．
各ディレクトリには，基本文法から標準ライブラリ，応用的なデザインパターンまでのコードが含まれている．

## プロジェクト構成

### 1. [Go Study](./go_study/)
Go言語の学習コンテンツである．

- **01 - 18**: 基本文法，データ構造，インターフェース，並行処理などの基礎から，Web I/O などの応用まで．
- **stdlib_definitions**: `fmt`，`os`，`http`，`context` など，主要な標準パッケージの使用例．
- **framework_definitions**: `gin`，`gorm`，`echo`，`grpc` など，主要なサードパーティ製フレームワークの導入例．

### 2. [TypeScript Study](./ts_study/)
TypeScriptの学習コンテンツである．

- **01 - 18**: 基本型，クラス，インターフェース，ユーティリティ型，DOM操作，Web I/O などの学習．
- **stdlib_definitions**: Node.js環境における `fs`，`path`，`http`，`crypto` などの標準モジュールの使用例．

## 実行方法

### Go
各ディレクトリに移動して `go run` で実行可能である．

```bash
cd go_study/01_basics
go run main.go
```

### TypeScript
`ts-node` 等を使用して実行することを想定している．

```bash
cd ts_study/01_basics
# ts-node がインストールされている場合
ts-node index.ts
```

## 言語
コード内のコメントおよび説明は主に日本語で記述されている．
