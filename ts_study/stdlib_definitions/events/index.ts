import { EventEmitter } from 'events';

/**
 * =============================================================================
 * TypeScript (Node.js) 標準モジュール定義集: events
 * Go のチャネル（channel）を用いた同期や通知に近い役割を果たします。
 * =============================================================================
 */

// --- 1. イベントの発行と受信 ---

// 使い方: const emitter = new EventEmitter(); emitter.on('event', (data) => { ... }); と書き，カスタムイベントを処理する．
export function MyEventEmitter(): EventEmitter {
    const ee = new EventEmitter();
    
    // 受信設定
    ee.on('greet', (name: string) => {
        console.log(`Hello, ${name}!`);
    });
    
    // 発行（通知）
    ee.emit('greet', 'TypeScript');
    
    return ee;
}

// 使い方: emitter.once('init', () => { ... }); と書き，一度きりのイベントを待ち受ける．
// Go の sync.Once に近い用途でも使われる．
export function MyOnceEvent(ee: EventEmitter): void {
    ee.once('init', () => {
        console.log('Initialized!');
    });
}

// 使い方: emitter.removeAllListeners('event'); と書き，メモリリークを防ぐためにリスナーを解除する．
export function MyRemoveListeners(ee: EventEmitter, event: string): void {
    ee.removeAllListeners(event);
}

// 使い方: emitter.prependListener('event', () => { ... }); と書き，他のリスナーよりも先に実行されるように登録する．
export function MyPrependListener(ee: EventEmitter, event: string, listener: (...args: any[]) => void): void {
    ee.prependListener(event, listener);
}

// 使い方: const count = ee.listenerCount('greet'); と書き，登録されているリスナーの数を確認する．
export function MyListenerCount(ee: EventEmitter, event: string): number {
    return ee.listenerCount(event);
}

// 使い方: ee.setMaxListeners(20); と書き，1つのイベントに登録できるリスナーの最大数を変更する（デフォルトは10）．
export function MySetMaxListeners(ee: EventEmitter, n: number): void {
    ee.setMaxListeners(n);
}

// 使い方: emitter.removeListener('event', callback); または emitter.off('event', callback); と書き，特定のリスナーを削除する．
export function MyOff(ee: EventEmitter, event: string, listener: (...args: any[]) => void): void {
    ee.off(event, listener);
}

// 使い方: ee.on('error', (err) => { ... }); と書き，ストリームやイベント発生元で起きたエラーをキャッチする．
// Node.js では error イベントを処理しないとプロセスがクラッシュするため必須．
export function MyHandleError(ee: EventEmitter): void {
    ee.on('error', (err) => {
        console.error('Handled error:', err);
    });
}

// 使い方: const names = ee.eventNames(); と書き，現在登録されているイベント名の配列を取得する．
export function MyEventNames(ee: EventEmitter): Array<string | symbol> {
    return ee.eventNames();
}

console.log('Loaded: events spec.');
