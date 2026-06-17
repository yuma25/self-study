/**
 * TSのモジュールルール: 
 * function だけではなく class や interface, type もすべて export 可能。
 */

export function hello() {
    console.log("Hello from mypackage (function)!");
}

/**
 * 公開クラス (Exported Class)
 * 公式: Handbook > Classes
 */
export class Logger {
    constructor(private prefix: string) {}

    log(message: string): void {
        console.log(`[${this.prefix}] ${message}`);
    }
}

/**
 * 公開インターフェース
 */
export interface Config {
    apiUrl: string;
    retryCount: number;
}
