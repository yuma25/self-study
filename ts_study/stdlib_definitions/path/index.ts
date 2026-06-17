import * as path from 'path';

/**
 * =============================================================================
 * TypeScript (Node.js) 標準モジュール定義集: path
 * Go の path/filepath に相当する機能を提供します。
 * =============================================================================
 */

// --- 1. パス結合・解析 ---

// 使い方: const fullPath = path.join('dir', 'subdir', 'file.txt'); と書き，OSごとの区切り文字を考慮して結合する．
export function MyJoin(...elems: string[]): string {
    return path.join(...elems);
}

// 使い方: const absPath = path.resolve('file.txt'); と書き，絶対パスを解決する．
export function MyResolve(...pathSegments: string[]): string {
    return path.resolve(...pathSegments);
}

// 使い方: const base = path.basename('/dir/file.txt'); と書き，ファイル名（"file.txt"）を取得する．
export function MyBasename(p: string): string {
    return path.basename(p);
}

// 使い方: const dir = path.dirname('/dir/file.txt'); と書き，親ディレクトリのパスを取得する．
export function MyDirname(p: string): string {
    return path.dirname(p);
}

// 使い方: const ext = path.extname('index.ts'); と書き，拡張子（".ts"）を取得する．
export function MyExtname(p: string): string {
    return path.extname(p);
}

// 使い方: const clean = path.normalize('/a/b/../c'); と書き，余分なドットを整理して正規化する．
export function MyNormalize(p: string): string {
    return path.normalize(p);
}

// 使い方: if (path.isAbsolute('/foo/bar')) { ... } と書き，パスが絶対パスか判定する．
export function MyIsAbsolute(p: string): boolean {
    return path.isAbsolute(p);
}

// 使い方: const rel = path.relative('/data/orighost/data', '/data/orighost/test/aaa'); と書き，相対パスを計算する．
export function MyRelative(from: string, to: string): string {
    return path.relative(from, to);
}

// 使い方: const nsPath = path.toNamespacedPath(p); と書き，Windows固有の名前空間付きパスに変換する（非Windowsではそのまま返す）．
export function MyToNamespacedPath(p: string): string {
    return path.toNamespacedPath(p);
}

// --- 2. 文字列とオブジェクトの相互変換 ---

// 使い方: const info = path.parse('/home/user/file.txt'); と書き，root, dir, base, ext, name を持つオブジェクトに分解する．
export function MyParse(p: string): path.ParsedPath {
    return path.parse(p);
}

// 使い方: const p = path.format({ dir: '/home/user', base: 'file.txt' }); と書き，オブジェクトからパス文字列を生成する．
export function MyFormat(p: path.FormatInputPathObject): string {
    return path.format(p);
}

// --- 3. 環境依存の定数 ---

// Linux/Mac なら '/', Windows なら '\'
export const MySeparator = path.sep;

// Linux/Mac なら ':', Windows なら ';' (PATH環境変数用)
export const MyDelimiter = path.delimiter;

console.log('Loaded: path spec.');
