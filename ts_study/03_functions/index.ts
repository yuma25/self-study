console.log("--- TypeScript Functions ---");

// 1. 基本的な関数
function add(a: number, b: number): number {
    return a + b;
}

// 2. 複数の戻り値を返す (疑似的: タプルやオブジェクトを使用)
// TS/JSには多値返しはないため、配列（タプル）を返して分割代入するのが一般的
function divide(a: number, b: number): [number, Error | null] {
    if (b === 0) {
        return [0, new Error("0で割ることはできません")];
    }
    return [a / b, null];
}

// 3. 例外処理 (TS/JSの一般的なエラーハンドリング)
function divideWithException(a: number, b: number): number {
    if (b === 0) {
        throw new Error("0で割ることはできません (Exception)");
    }
    return a / b;
}

// 4. アロー関数 (TS/JSで多用される)
const multiply = (a: number, b: number): number => { return a * b; };

// --- 実行 ---
console.log("Add(10, 5):", add(10, 5));

// 多値返しのシミュレーション
const [result, err] = divide(10, 2);
if (err) {
    console.error("Error:", err.message);
} else {
    console.log("Divide(10, 2):", result);
}

// 例外処理のパターン
try {
    const res = divideWithException(10, 0);
    console.log("Result:", res);
} catch (e) {
    if (e instanceof Error) {
        console.log("Expected Exception:", e.message);
    }
}

console.log("Multiply(10, 5):", multiply(10, 5));
