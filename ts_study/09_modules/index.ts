/**
 * 名前付きインポート (Named Import)
 */
import { hello, Logger, Config } from "./mypackage/lib";

console.log("--- TypeScript Modules (Functions & Classes) ---");

// 1. 関数の使用
hello();

// 2. クラスの使用 (インスタンス化)
const myLogger = new Logger("APP");
myLogger.log("Application started");

// 3. インターフェースの使用
const config: Config = {
    apiUrl: "https://api.example.com",
    retryCount: 3
};
console.log("Config:", config);

/**
 * 実行コマンド:
 * npx tsx ts_study/09_modules/index.ts
 */
