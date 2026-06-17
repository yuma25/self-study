import * as zlib from 'zlib';
import { promisify } from 'util';

/**
 * =============================================================================
 * TypeScript (Node.js) 標準モジュール定義集: zlib
 * データの圧縮・解凍（Gzip, Deflate, Brotli 等）機能を提供します。
 * Go の compress/gzip, compress/flate パッケージに相当します。
 * =============================================================================
 */

// --- 1. 高レベル非同期関数 (Promise) ---

// 使い方: const compressed = await gzip('data'); と書き，Gzip圧縮を行う．
const gzip = promisify(zlib.gzip);
// 使い方: const original = await gunzip(compressed); と書き，Gzip解凍を行う．
const gunzip = promisify(zlib.gunzip);

export { gzip, gunzip };

// --- 2. ストリームベースの圧縮 (推奨) ---

// 使い方: fs.createReadStream('in.txt').pipe(zlib.createGzip()).pipe(fs.createWriteStream('in.txt.gz')); 
// と書き，メモリを節約しつつ巨大なファイルを圧縮する．
export function MyCreateGzip(options?: zlib.ZlibOptions): zlib.Gzip {
    return zlib.createGzip(options);
}

export function MyCreateGunzip(options?: zlib.ZlibOptions): zlib.Gunzip {
    return zlib.createGunzip(options);
}

// 使い方: fs.createReadStream('in.txt').pipe(zlib.createBrotliCompress()).pipe(fs.createWriteStream('in.txt.br'));
// と書き，Gzipより高圧縮なBrotliで圧縮する．モダンなWeb配信の標準手法．
export function MyCreateBrotliCompress(options?: zlib.BrotliOptions): zlib.BrotliCompress {
    return zlib.createBrotliCompress(options);
}

export function MyCreateBrotliDecompress(options?: zlib.BrotliOptions): zlib.BrotliDecompress {
    return zlib.createBrotliDecompress(options);
}

// --- 3. その他のアルゴリズム ---

// 使い方: zlib.deflate(data, (err, buf) => { ... }); と書き，Deflate圧縮を行う．
export function MyDeflate(data: zlib.InputType, callback: zlib.CompressCallback): void {
    zlib.deflate(data, callback);
}

// 使い方: const output = zlib.brotliCompressSync(data); と書き，Brotliアルゴリズムで圧縮する（同期版）．
export function MyBrotliCompressSync(data: zlib.InputType): Buffer {
    return zlib.brotliCompressSync(data);
}

// --- 4. 便利な定数 (圧縮レベル等) ---

export const MyConstants = {
    // 圧縮レベル設定用
    Z_NO_COMPRESSION: zlib.constants.Z_NO_COMPRESSION,
    Z_BEST_SPEED: zlib.constants.Z_BEST_SPEED,
    Z_BEST_COMPRESSION: zlib.constants.Z_BEST_COMPRESSION,
    Z_DEFAULT_COMPRESSION: zlib.constants.Z_DEFAULT_COMPRESSION,
};

console.log('Loaded: zlib spec.');
