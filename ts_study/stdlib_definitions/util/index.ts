import * as util from 'util';

/**
 * =============================================================================
 * TypeScript (Node.js) 標準モジュール定義集: util
 * 開発に役立つ各種ユーティリティ関数を提供します。
 * Go の fmt や reflect パッケージの一部の機能に相当します。
 * =============================================================================
 */

// --- 1. 非同期処理の変換 ---

// 使い方: const readFile = util.promisify(fs.readFile); と書き，従来のコールバック形式の関数を Promise を返す形式に変換する．
// 現代の Node.js では fs/promises 等の利用が優先されるが，古いライブラリの Promise 化に必須．
export function MyPromisify<T extends (...args: any[]) => any>(fn: T): any {
    return util.promisify(fn);
}

// 使い方: const callbackFn = util.callbackify(asyncFn); と書き，Promise 形式の関数をコールバック形式に変換する．
export function MyCallbackify<T extends (...args: any[]) => Promise<any>>(fn: T): any {
    return util.callbackify(fn);
}

// --- 2. 書式化・デバッグ出力 ---

// 使い方: const str = util.format('Hello %s, count: %d', 'world', 123); と書き，C言語の printf スタイルで文字列を構築する．
// Go の fmt.Sprintf に相当．
export function MyFormat(format: string, ...args: any[]): string {
    return util.format(format, ...args);
}

// 使い方: const debugStr = util.inspect(obj, { depth: null, colors: true }); と書き，オブジェクトの中身を詳細に文字列化する（デバッグ用）．
export function MyInspect(obj: any, options?: util.InspectOptions): string {
    return util.inspect(obj, options);
}

// --- 3. 比較・型判定 ---

// 使い方: if (util.isDeepStrictEqual(obj1, obj2)) { ... } と書き，2つのオブジェクトが構造的に完全に一致するか（再帰的に）比較する．
export function MyIsDeepStrictEqual(val1: any, val2: any): boolean {
    return util.isDeepStrictEqual(val1, val2);
}

// 使い方: if (util.types.isPromise(obj)) { ... } と書き，対象が Promise かどうかを正確に判定する．
// 内部には isDate, isRegExp, isMap, isSet 等，数多くの型判定関数が含まれる．
export const MyTypes = util.types;

// --- 4. デバッグ・システムエラー ---

// 使い方: const debug = util.debuglog('foo'); debug('message'); と書き，NODE_DEBUG=foo 環境変数が設定されている時のみ出力するロガーを作る．
export function MyDebugLog(section: string): (msg: string, ...args: any[]) => void {
    return util.debuglog(section);
}

// 使い方: const name = util.getSystemErrorName(-2); と書き，数値のエラーコードからエラー名（'ENOENT'等）を取得する．
export function MyGetSystemErrorName(err: number): string {
    return util.getSystemErrorName(err);
}

// --- 5. 継承 (Legacy) ---

// 使い方: util.inherits(Child, Parent); と書き，プロトタイプ継承を設定する（現在は class 構文の利用が推奨）．
export function MyInherits(constructor: any, superConstructor: any): void {
    util.inherits(constructor, superConstructor);
}

// --- 5. テキストのエンコード・デコード ---

// 使い方: const decoder = new util.TextDecoder('utf-8'); const str = decoder.decode(buffer); と書き，バイト配列を文字列に変換する．
// Global にも存在するが，util にも定義されている．
export const MyTextDecoder = util.TextDecoder;

// 使い方: const encoder = new util.TextEncoder(); const buf = encoder.encode('text'); と書き，文字列を Uint8Array に変換する．
export const MyTextEncoder = util.TextEncoder;

console.log('Loaded: util spec.');
