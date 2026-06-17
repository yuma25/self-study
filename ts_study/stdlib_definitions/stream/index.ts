import * as stream from 'stream';
import { pipeline, finished } from 'stream/promises';

/**
 * =============================================================================
 * TypeScript (Node.js) 標準モジュール定義集: stream
 * データの逐次処理（ストリーミング）のための抽象インターフェースを提供します。
 * Go の io.Reader, io.Writer, io.Pipe 等に相当します。
 * =============================================================================
 */

// --- 1. 基本となる 4 つのストリーム型 ---

// 使い方: class MyReadable extends stream.Readable { ... } と書き，独自のデータ読み込み元を作る．
// Go の io.Reader に相当．
export type MyReadable = stream.Readable;

// 使い方: class MyWritable extends stream.Writable { ... } と書き，独自のデータ書き込み先を作る．
// Go の io.Writer に相当．
export type MyWritable = stream.Writable;

// 使い方: 読み書き両方の機能を持つストリーム（TCPソケット等）に使用する．
export type MyDuplex = stream.Duplex;

// 使い方: 入力を受け取り，加工して出力するストリーム（圧縮，暗号化等）に使用する．
export type MyTransform = stream.Transform;

// --- 2. 高レベルなストリーム操作 (Promiseベース) ---

// 使い方: await pipeline(rs, transform, ws); と書き，複数のストリームを安全に連結する．
// Go の io.Copy を多段にしたような役割．エラー発生時にすべてのストリームを自動で閉じ，メモリリークを防ぐ．
export async function MyPipeline(streams: Array<stream.Readable | stream.Writable | stream.Duplex>): Promise<void> {
    await pipeline(streams as any);
}

// 使い方: await finished(rs); と書き，ストリームの読み書きが完全に終了（またはエラー）するのを待機する．
export async function MyFinished(s: stream.Readable | stream.Writable | stream.Duplex): Promise<void> {
    await finished(s);
}

// 使い方: const combined = stream.compose(rs, transform, ws); と書き，複数のストリームを1つのストリームに統合する（Node.js 16+）．
export function MyCompose(...streams: Array<stream.Stream | Iterable<any> | AsyncIterable<any>>): stream.Duplex {
    return (stream as any).compose(...streams);
}

// --- 3. ユーティリティ ---

// 使い方: stream.Readable.from(['a', 'b', 'c']) と書き，配列や反復可能オブジェクトからストリームを生成する．
export function MyStreamFrom(iterable: Iterable<any> | AsyncIterable<any>): stream.Readable {
    return stream.Readable.from(iterable);
}

// 使い方: rs.pipe(ws) と書き，読み込みストリームと書き込みストリームを連結する（古典的な手法）．
// 現代のコードではエラー処理が自動化されている pipeline の利用が推奨される．

console.log('Loaded: stream spec.');
