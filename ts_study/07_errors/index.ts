console.log("--- TypeScript Errors ---");

/**
 * 1. カスタムエラークラス (公式: Handbook > Classes)
 */
class MyError extends Error {
    constructor(public code: number, message: string) {
        super(message);
        this.name = "MyError";
    }
}

/**
 * 2. try-catch-finally (標準的な例外処理)
 */
function process() {
    console.log("Start processing...");
    try {
        throw new MyError(404, "Not Found");
    } catch (e) {
        if (e instanceof MyError) {
            console.log(`Caught custom error: [${e.code}] ${e.message}`);
        } else {
            console.log("Caught unknown error:", e);
        }
    } finally {
        // Go の defer に相当するクリーンアップ処理
        console.log("Finally: process finished!");
    }
}

/**
 * 3. Result 型パターン (Goに近い明示的なエラー処理)
 * TypeScriptの Union Type を使って「成功または失敗」を表現する
 */
type Result<T, E> = { success: true; data: T } | { success: false; error: E };

function doSomethingSafe(fail: boolean): Result<string, string> {
    if (fail) {
        return { success: false, error: "Something failed" };
    }
    return { success: true, data: "Safe data" };
}

// 実行
process();

const res = doSomethingSafe(true);
if (!res.success) {
    console.log("Handle error explicitly:", res.error);
} else {
    console.log("Handle data:", res.data);
}

/**
 * 実行コマンド:
 * npx tsx ts_study/07_errors/index.ts
 */
